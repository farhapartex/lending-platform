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
	"syscall"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/config"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
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

	router := transporthttp.NewRouter(transporthttp.RouterParams{
		Config:             cfg,
		Logger:             log,
		HealthService:      healthService,
		TransactionService: stores.transactions,
		LiquidationService: stores.liquidations,
		Masker:             masker,
	})

	server := &http.Server{
		Addr:              net.JoinHostPort("", strconv.Itoa(cfg.HTTPPort)),
		Handler:           router,
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

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
		transactions: service.NewTransactionService(service.TransactionServiceParams{
			Users:        repository.NewUserRepository(db),
			Transactions: repository.NewTransactionRepository(db),
			Checkpoints:  checkpoints,
		}),
		liquidations: service.NewLiquidationService(service.LiquidationServiceParams{
			Liquidations: repository.NewLiquidationRepository(db),
			Checkpoints:  checkpoints,
		}),
	}

	return built, closeDatabase, nil
}
