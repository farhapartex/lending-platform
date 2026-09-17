package indexer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/farhapartex/lending-platform/core-service/internal/chain"
	"github.com/farhapartex/lending-platform/core-service/internal/config"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

type BootstrapParams struct {
	Chain      config.ChainConfig
	Eth        *ethclient.Client
	MarketView domain.MarketViewReader
	Assets     domain.AssetRepository
	Markets    domain.MarketRepository
}

func EnsureMarket(ctx context.Context, params BootstrapParams) (domain.Market, error) {
	existing, err := params.Markets.ByPoolAddress(ctx, params.Chain.ChainID, strings.ToLower(params.Chain.Contracts.Pool))
	if err == nil {
		return existing, nil
	}

	if !errors.Is(err, domain.ErrNotFound) {
		return domain.Market{}, err
	}

	collateral, err := ensureAsset(ctx, params, params.Chain.Contracts.CollateralToken, true, false)
	if err != nil {
		return domain.Market{}, err
	}

	debt, err := ensureAsset(ctx, params, params.Chain.Contracts.DebtToken, false, true)
	if err != nil {
		return domain.Market{}, err
	}

	view, err := params.MarketView.ReadMarketView(ctx)
	if err != nil {
		return domain.Market{}, err
	}

	market := domain.Market{
		ChainID:                   params.Chain.ChainID,
		CollateralAssetID:         collateral.ID,
		DebtAssetID:               debt.ID,
		PoolAddress:               strings.ToLower(params.Chain.Contracts.Pool),
		CollateralVaultAddress:    strings.ToLower(params.Chain.Contracts.Vault),
		ControllerAddress:         strings.ToLower(params.Chain.Contracts.Controller),
		LiquidationManagerAddress: strings.ToLower(params.Chain.Contracts.LiquidationManager),
		InterestRateModelAddress:  strings.ToLower(params.Chain.Contracts.RateModel),
		OracleAdapterAddress:      strings.ToLower(params.Chain.Contracts.Oracle),
		MaxLTVBps:                 view.MaxLtvBps,
		LiquidationThresholdBps:   view.LiquidationThresholdBps,
		LiquidationBonusBps:       view.LiquidationBonusBps,
		ReserveFactorBps:          view.ReserveFactorBps,
		RecommendedLTVBps:         recommendedFrom(view.MaxLtvBps),
		KinkUtilizationBps:        view.KinkUtilizationBps,
		MinDeposit:                bigmath.FromBig(orZero(view.MinDeposit)),
		Status:                    domain.MarketStatusActive,
		DeployedAtBlock:           int64(params.Chain.StartBlock),
	}

	if err := params.Markets.Upsert(ctx, &market); err != nil {
		return domain.Market{}, err
	}

	return market, nil
}

func ensureAsset(
	ctx context.Context,
	params BootstrapParams,
	address string,
	isCollateral bool,
	isBorrowable bool,
) (domain.Asset, error) {
	normalized := strings.ToLower(address)

	existing, err := params.Assets.ByAddress(ctx, params.Chain.ChainID, normalized)
	if err == nil {
		return existing, nil
	}

	if !errors.Is(err, domain.ErrNotFound) {
		return domain.Asset{}, err
	}

	info, err := chain.ReadToken(ctx, params.Eth, address, params.Chain.RequestTimeout)
	if err != nil {
		return domain.Asset{}, fmt.Errorf("reading token %s failed: %w", address, err)
	}

	asset := domain.Asset{
		ChainID:         params.Chain.ChainID,
		Address:         normalized,
		AddressChecksum: address,
		Symbol:          info.Symbol,
		Name:            info.Name,
		Decimals:        info.Decimals,
		IsCollateral:    isCollateral,
		IsBorrowable:    isBorrowable,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}

	if err := params.Assets.Upsert(ctx, &asset); err != nil {
		return domain.Asset{}, err
	}

	return asset, nil
}

func recommendedFrom(maxLtvBps int32) int32 {
	recommended := maxLtvBps * 55 / 75

	if recommended < 1 {
		return maxLtvBps
	}

	return recommended
}

func orZero(value *big.Int) *big.Int {
	if value == nil {
		return big.NewInt(0)
	}

	return value
}
