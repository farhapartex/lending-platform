package repository

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

type marketRepository struct {
	db *gorm.DB
}

func NewMarketRepository(db *gorm.DB) domain.MarketRepository {
	return &marketRepository{db: db}
}

func (r *marketRepository) ByID(ctx context.Context, id int64) (domain.Market, error) {
	if err := requirePositiveID(id, "market id"); err != nil {
		return domain.Market{}, err
	}

	var market domain.Market

	err := r.withAssets(ctx).First(&market, "markets.id = ?", id).Error
	if err != nil {
		return domain.Market{}, translate(err, "market by id")
	}

	return market, nil
}

func (r *marketRepository) ByPoolAddress(
	ctx context.Context,
	chainID int64,
	poolAddress string,
) (domain.Market, error) {
	normalized, err := normalizeAddress(poolAddress, "pool address")
	if err != nil {
		return domain.Market{}, err
	}

	var market domain.Market

	err = r.withAssets(ctx).
		First(&market, "markets.chain_id = ? AND markets.pool_address = ?", chainID, normalized).Error
	if err != nil {
		return domain.Market{}, translate(err, "market by pool address")
	}

	return market, nil
}

func (r *marketRepository) List(ctx context.Context, chainID int64) ([]domain.Market, error) {
	markets := make([]domain.Market, 0)

	err := r.withAssets(ctx).
		Where("markets.chain_id = ?", chainID).
		Order("markets.id ASC").
		Find(&markets).Error
	if err != nil {
		return nil, translate(err, "market list")
	}

	return markets, nil
}

func (r *marketRepository) Upsert(ctx context.Context, market *domain.Market) error {
	if market == nil {
		return domain.ErrInvalidInput
	}

	addresses := []struct {
		field  string
		target *string
	}{
		{field: "pool address", target: &market.PoolAddress},
		{field: "collateral vault address", target: &market.CollateralVaultAddress},
		{field: "controller address", target: &market.ControllerAddress},
		{field: "liquidation manager address", target: &market.LiquidationManagerAddress},
		{field: "interest rate model address", target: &market.InterestRateModelAddress},
		{field: "oracle adapter address", target: &market.OracleAdapterAddress},
	}

	for _, address := range addresses {
		normalized, err := normalizeAddress(*address.target, address.field)
		if err != nil {
			return err
		}

		*address.target = normalized
	}

	return translate(r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "chain_id"}, {Name: "pool_address"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"collateral_asset_id",
				"debt_asset_id",
				"collateral_vault_address",
				"controller_address",
				"liquidation_manager_address",
				"interest_rate_model_address",
				"oracle_adapter_address",
				"max_ltv_bps",
				"liquidation_threshold_bps",
				"liquidation_bonus_bps",
				"reserve_factor_bps",
				"recommended_ltv_bps",
				"kink_utilization_bps",
				"min_deposit",
				"status",
				"updated_at",
			}),
		}).
		Omit("CollateralAsset", "DebtAsset").
		Create(market).Error, "market upsert")
}

func (r *marketRepository) withAssets(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).
		Preload("CollateralAsset").
		Preload("DebtAsset")
}
