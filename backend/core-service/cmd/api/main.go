package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/farhapartex/lending-platform/core-service/internal/chain"
	"github.com/farhapartex/lending-platform/core-service/internal/config"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/indexer"
	"github.com/farhapartex/lending-platform/core-service/internal/platform/database"
	"github.com/farhapartex/lending-platform/core-service/internal/platform/logger"
	"github.com/farhapartex/lending-platform/core-service/internal/repository"
	"github.com/farhapartex/lending-platform/core-service/internal/service"
	transporthttp "github.com/farhapartex/lending-platform/core-service/internal/transport/http"
	"github.com/farhapartex/lending-platform/core-service/pkg/idmask"
)

var (
	version = "dev"
	commit  = "unknown"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "core-service failed to start: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	startedAt := time.Now()

	cfg, err := config.Load(version)
	if err != nil {
		return err
	}

	log := logger.New(cfg.LogLevel, cfg.AppEnv, cfg.ServiceName, cfg.ServiceVersion)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	healthService := service.NewHealthService(service.HealthServiceParams{
		ServiceName: cfg.ServiceName,
		Version:     cfg.ServiceVersion,
		Environment: cfg.AppEnv,
		StartedAt:   startedAt,
	})

	masker, err := idmask.New(cfg.EffectiveIDMaskSecret())
	if err != nil {
		return fmt.Errorf("id masking could not be set up: %w", err)
	}

	stores, closeDatabase, err := buildStores(cfg, log)
	if err != nil {
		return err
	}

	if closeDatabase != nil {
		defer closeDatabase()
	}

	if stores.db != nil && cfg.Chain.IndexerEnabled {
		stopIndexer, err := startIndexer(ctx, cfg, stores, log)
		if err != nil {
			log.Error("the indexer could not start", slog.String("error", err.Error()))
		} else if stopIndexer != nil {
			defer stopIndexer()
		}
	}

	assets := loadAssetMetadata(ctx, stores, cfg)

	router := transporthttp.NewRouter(transporthttp.RouterParams{
		Config:             cfg,
		Logger:             log,
		HealthService:      healthService,
		TransactionService: stores.transactions,
		LiquidationService: stores.liquidations,
		Masker:             masker,
		CollateralDecimals: assets.collateralDecimals,
		CollateralSymbol:   assets.collateralSymbol,
		DebtDecimals:       assets.debtDecimals,
		DebtSymbol:         assets.debtSymbol,
	})

	server := &http.Server{
		Addr:              net.JoinHostPort("", strconv.Itoa(cfg.HTTPPort)),
		Handler:           router,
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Info(
			"http server listening",
			slog.Int("port", cfg.HTTPPort),
			slog.String("commit", commit),
			slog.String("health_endpoint", transporthttp.APIBasePath+"/health"),
		)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("http server stopped unexpectedly: %w", err)
	case <-ctx.Done():
		log.Info("shutdown signal received", slog.Duration("grace_period", cfg.ShutdownGrace))
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownGrace)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown failed: %w", err)
	}

	log.Info("shutdown complete")

	return nil
}

type stores struct {
	transactions domain.TransactionService
	liquidations domain.LiquidationService
	db           *gorm.DB
}

func buildStores(cfg config.Config, log *slog.Logger) (stores, func(), error) {
	if cfg.DatabaseURL == "" {
		log.Warn("DATABASE_URL is not set, database backed endpoints will not be served")

		return stores{}, nil, nil
	}

	db, err := database.Open(database.Options{
		DSN:             cfg.DatabaseURL,
		MaxOpenConns:    cfg.DatabaseMaxOpenConns,
		MaxIdleConns:    cfg.DatabaseMaxIdleConns,
		ConnMaxLifetime: cfg.DatabaseConnMaxLifetime,
		ConnMaxIdleTime: cfg.DatabaseConnMaxIdleTime,
		LogQueries:      cfg.DatabaseLogQueries,
	})
	if err != nil {
		return stores{}, nil, fmt.Errorf("database connection failed: %w", err)
	}

	closeDatabase := func() {
		pool, poolErr := db.DB()
		if poolErr != nil {
			return
		}

		if closeErr := pool.Close(); closeErr != nil {
			log.Error("closing the database failed", slog.String("error", closeErr.Error()))
		}
	}

	checkpoints := repository.NewCheckpointRepository(db)

	built := stores{
		db: db,
		transactions: service.NewTransactionService(service.TransactionServiceParams{
			Users:        repository.NewUserRepository(db),
			Transactions: repository.NewTransactionRepository(db),
			Checkpoints:  checkpoints,
		}),
		liquidations: service.NewLiquidationService(service.LiquidationServiceParams{
			Liquidations: repository.NewLiquidationRepository(db),
			Positions:    repository.NewPositionRepository(db),
			Users:        repository.NewUserRepository(db),
			Checkpoints:  checkpoints,
		}),
	}

	return built, closeDatabase, nil
}

func startIndexer(ctx context.Context, cfg config.Config, built stores, log *slog.Logger) (func(), error) {
	client, err := chain.Dial(ctx, chain.ClientParams{
		RPCURL:          cfg.Chain.RPCURL,
		ExpectedChainID: cfg.Chain.ChainID,
		RequestTimeout:  cfg.Chain.RequestTimeout,
	})
	if err != nil {
		return nil, err
	}

	marketReader, err := chain.NewMarketReader(client.Eth(), chain.MarketReaderParams{
		LensAddress:    cfg.Chain.Contracts.Lens,
		RequestTimeout: cfg.Chain.RequestTimeout,
	})
	if err != nil {
		client.Close()

		return nil, err
	}

	accountReader, err := chain.NewAccountReader(client.Eth(), chain.MarketReaderParams{
		LensAddress:    cfg.Chain.Contracts.Lens,
		RequestTimeout: cfg.Chain.RequestTimeout,
	})
	if err != nil {
		client.Close()

		return nil, err
	}

	market, err := indexer.EnsureMarket(ctx, indexer.BootstrapParams{
		Chain:      cfg.Chain,
		Eth:        client.Eth(),
		MarketView: marketReader,
		Assets:     repository.NewAssetRepository(built.db),
		Markets:    repository.NewMarketRepository(built.db),
	})
	if err != nil {
		client.Close()

		return nil, err
	}

	decoder, err := indexer.NewDecoder(indexer.ContractSet{
		Pool:               common.HexToAddress(cfg.Chain.Contracts.Pool),
		Vault:              common.HexToAddress(cfg.Chain.Contracts.Vault),
		Controller:         common.HexToAddress(cfg.Chain.Contracts.Controller),
		LiquidationManager: common.HexToAddress(cfg.Chain.Contracts.LiquidationManager),
	})
	if err != nil {
		client.Close()

		return nil, err
	}

	tracker := indexer.NewPositionTracker(indexer.PositionTrackerParams{
		ChainID:   cfg.Chain.ChainID,
		Market:    market,
		Accounts:  accountReader,
		Users:     repository.NewUserRepository(built.db),
		Positions: repository.NewPositionRepository(built.db),
	})

	runner := indexer.NewRunner(indexer.RunnerParams{
		Chain:   cfg.Chain,
		Client:  client,
		Decoder: decoder,
		Ingestor: indexer.NewIngestor(indexer.IngestorParams{
			ChainID:      cfg.Chain.ChainID,
			Market:       market,
			Users:        repository.NewUserRepository(built.db),
			Events:       repository.NewProtocolEventRepository(built.db),
			Transactions: repository.NewTransactionRepository(built.db),
			Liquidations: repository.NewLiquidationRepository(built.db),
		}),
		Events:      repository.NewProtocolEventRepository(built.db),
		Positions:   tracker,
		Blocks:      repository.NewIndexedBlockRepository(built.db),
		Checkpoints: repository.NewCheckpointRepository(built.db),
		Logger:      log,
	})

	head, err := client.HeadBlock(ctx)
	if err == nil {
		if refreshed, backfillErr := tracker.Backfill(ctx, int64(head)); backfillErr != nil {
			log.Warn("backfilling positions stopped early", slog.String("error", backfillErr.Error()))
		} else if refreshed > 0 {
			log.Info("valued the positions already on record", slog.Int("positions", refreshed))
		}
	}

	go runner.Run(ctx)
	go tracker.Revalue(ctx, client.HeadBlock, cfg.Chain.SnapshotInterval, log)

	log.Info(
		"indexer attached to the market",
		slog.Int64("market_id", market.ID),
		slog.String("pool", market.PoolAddress),
	)

	return client.Close, nil
}

type assetMetadata struct {
	collateralDecimals int16
	collateralSymbol   string
	debtDecimals       int16
	debtSymbol         string
}

func loadAssetMetadata(ctx context.Context, built stores, cfg config.Config) assetMetadata {
	fallback := assetMetadata{
		collateralDecimals: 18,
		collateralSymbol:   "WETH",
		debtDecimals:       6,
		debtSymbol:         "USDC",
	}

	if built.db == nil {
		return fallback
	}

	assets := repository.NewAssetRepository(built.db)

	collateral, err := assets.ByAddress(ctx, cfg.Chain.ChainID, strings.ToLower(cfg.Chain.Contracts.CollateralToken))
	if err != nil {
		return fallback
	}

	debt, err := assets.ByAddress(ctx, cfg.Chain.ChainID, strings.ToLower(cfg.Chain.Contracts.DebtToken))
	if err != nil {
		return fallback
	}

	return assetMetadata{
		collateralDecimals: collateral.Decimals,
		collateralSymbol:   collateral.Symbol,
		debtDecimals:       debt.Decimals,
		debtSymbol:         debt.Symbol,
	}
}
