package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

const (
	defaultLiquidationLimit = 25
	maxLiquidationLimit     = 200
)

var liquidationRelations = []string{
	"Borrower",
	"Liquidator",
	"Market",
	"Market.CollateralAsset",
	"Market.DebtAsset",
}

type liquidationRepository struct {
	db *gorm.DB
}

func NewLiquidationRepository(db *gorm.DB) domain.LiquidationRepository {
	return &liquidationRepository{db: db}
}

func (r *liquidationRepository) List(
	ctx context.Context,
	query domain.LiquidationQuery,
) ([]domain.Liquidation, error) {
	if query.MarketID != nil {
		if err := requirePositiveID(*query.MarketID, "market id"); err != nil {
			return nil, err
		}
	}

	liquidations := make([]domain.Liquidation, 0)

	statement := withRelations(r.db.WithContext(ctx), liquidationRelations)

	if query.MarketID != nil {
		statement = statement.Where("market_id = ?", *query.MarketID)
	}

	statement = applyKeysetCursor(statement, query.After)

	result := statement.
		Order("block_time DESC, id DESC").
		Limit(boundedLimit(query.Limit, defaultLiquidationLimit, maxLiquidationLimit)).
		Find(&liquidations)

	if result.Error != nil {
		return nil, translate(result.Error, "liquidation list")
	}

	return liquidations, nil
}

func (r *liquidationRepository) ByID(ctx context.Context, id int64) (domain.Liquidation, error) {
	if err := requirePositiveID(id, "liquidation id"); err != nil {
		return domain.Liquidation{}, err
	}

	var liquidation domain.Liquidation

	statement := withRelations(r.db.WithContext(ctx), liquidationRelations).
		First(&liquidation, "id = ?", id)

	if statement.Error != nil {
		return domain.Liquidation{}, translate(statement.Error, "liquidation by id")
	}

	return liquidation, nil
}

func (r *liquidationRepository) Insert(ctx context.Context, liquidation *domain.Liquidation) error {
	if liquidation == nil {
		return domain.ErrInvalidInput
	}

	if err := requirePositiveID(liquidation.BorrowerUserID, "borrower user id"); err != nil {
		return err
	}

	if err := requirePositiveID(liquidation.LiquidatorUserID, "liquidator user id"); err != nil {
		return err
	}

	return translate(
		r.db.WithContext(ctx).
			Omit("Event", "Market", "Borrower", "Liquidator").
			Create(liquidation).Error,
		"liquidation insert",
	)
}
