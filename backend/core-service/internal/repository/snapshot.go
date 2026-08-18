package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

const (
	defaultSnapshotLimit = 200
	maxSnapshotLimit     = 2000
)

type marketSnapshotRepository struct {
	db *gorm.DB
}

func NewMarketSnapshotRepository(db *gorm.DB) domain.MarketSnapshotRepository {
	return &marketSnapshotRepository{db: db}
}

func (r *marketSnapshotRepository) Insert(ctx context.Context, snapshot *domain.MarketSnapshot) error {
	if snapshot == nil {
		return domain.ErrInvalidInput
	}

	if err := requirePositiveID(snapshot.MarketID, "market id"); err != nil {
		return err
	}

	return translate(r.db.WithContext(ctx).Omit("Market").Create(snapshot).Error, "market snapshot insert")
}

func (r *marketSnapshotRepository) Latest(ctx context.Context, marketID int64) (domain.MarketSnapshot, error) {
	if err := requirePositiveID(marketID, "market id"); err != nil {
		return domain.MarketSnapshot{}, err
	}

	var snapshot domain.MarketSnapshot

	err := r.db.WithContext(ctx).
		Where("market_id = ?", marketID).
		Order("captured_at DESC, id DESC").
		First(&snapshot).Error
	if err != nil {
		return domain.MarketSnapshot{}, translate(err, "latest market snapshot")
	}

	return snapshot, nil
}

func (r *marketSnapshotRepository) Since(
	ctx context.Context,
	marketID int64,
	since time.Time,
	limit int,
) ([]domain.MarketSnapshot, error) {
	if err := requirePositiveID(marketID, "market id"); err != nil {
		return nil, err
	}

	snapshots := make([]domain.MarketSnapshot, 0)

	err := r.db.WithContext(ctx).
		Where("market_id = ? AND captured_at >= ?", marketID, since).
		Order("captured_at ASC, id ASC").
		Limit(boundedLimit(limit, defaultSnapshotLimit, maxSnapshotLimit)).
		Find(&snapshots).Error
	if err != nil {
		return nil, translate(err, "market snapshots since")
	}

	return snapshots, nil
}
