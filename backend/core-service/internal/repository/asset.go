package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/ethaddr"
)

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) domain.AssetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) ByID(ctx context.Context, id int64) (domain.Asset, error) {
	if err := requirePositiveID(id, "asset id"); err != nil {
		return domain.Asset{}, err
	}

	var asset domain.Asset

	if err := r.db.WithContext(ctx).First(&asset, "id = ?", id).Error; err != nil {
		return domain.Asset{}, translate(err, "asset by id")
	}

	return asset, nil
}

func (r *assetRepository) ByAddress(ctx context.Context, chainID int64, address string) (domain.Asset, error) {
	normalized, err := normalizeAddress(address, "asset address")
	if err != nil {
		return domain.Asset{}, err
	}

	var asset domain.Asset

	query := r.db.WithContext(ctx).First(&asset, "chain_id = ? AND address = ?", chainID, normalized)
	if query.Error != nil {
		return domain.Asset{}, translate(query.Error, "asset by address")
	}

	return asset, nil
}

func (r *assetRepository) List(ctx context.Context, chainID int64) ([]domain.Asset, error) {
	assets := make([]domain.Asset, 0)

	query := r.db.WithContext(ctx).
		Where("chain_id = ?", chainID).
		Order("symbol ASC").
		Find(&assets)

	if query.Error != nil {
		return nil, translate(query.Error, "asset list")
	}

	return assets, nil
}

func (r *assetRepository) Upsert(ctx context.Context, asset *domain.Asset) error {
	if asset == nil {
		return domain.ErrInvalidInput
	}

	normalized, checksum, err := ethaddr.NormalizeWithChecksum(asset.Address)
	if err != nil {
		return fmt.Errorf("%w: asset address %s", domain.ErrInvalidInput, err)
	}

	asset.Address = normalized
	asset.AddressChecksum = checksum

	return translate(r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "chain_id"}, {Name: "address"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"address_checksum",
				"symbol",
				"name",
				"decimals",
				"is_collateral",
				"is_borrowable",
				"price_feed_address",
				"price_feed_decimals",
				"updated_at",
			}),
		}).
		Create(asset).Error, "asset upsert")
}
