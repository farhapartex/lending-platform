package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

type protocolEventRepository struct {
	db *gorm.DB
}

func NewProtocolEventRepository(db *gorm.DB) domain.ProtocolEventRepository {
	return &protocolEventRepository{db: db}
}

func (r *protocolEventRepository) Insert(ctx context.Context, event *domain.ProtocolEvent) error {
	if event == nil {
		return domain.ErrInvalidInput
	}

	event.CreatedAt = time.Now().UTC()

	statement := r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chain_id"}, {Name: "tx_hash"}, {Name: "log_index"}},
		DoNothing: true,
	}).Create(event)

	if statement.Error != nil {
		return translate(statement.Error, "protocol event")
	}

	if statement.RowsAffected == 0 {
		return r.loadExisting(ctx, event)
	}

	return nil
}

func (r *protocolEventRepository) loadExisting(ctx context.Context, event *domain.ProtocolEvent) error {
	var existing domain.ProtocolEvent

	statement := r.db.WithContext(ctx).
		Where("chain_id = ? AND tx_hash = ? AND log_index = ?", event.ChainID, event.TxHash, event.LogIndex).
		First(&existing)

	if statement.Error != nil {
		return translate(statement.Error, "protocol event")
	}

	event.ID = existing.ID

	return nil
}

type indexedBlockRepository struct {
	db *gorm.DB
}

func NewIndexedBlockRepository(db *gorm.DB) domain.IndexedBlockRepository {
	return &indexedBlockRepository{db: db}
}

func (r *indexedBlockRepository) Insert(ctx context.Context, block *domain.IndexedBlock) error {
	if block == nil {
		return domain.ErrInvalidInput
	}

	block.CreatedAt = time.Now().UTC()

	statement := r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chain_id"}, {Name: "block_number"}},
		DoUpdates: clause.AssignmentColumns([]string{"block_hash", "parent_hash", "block_time"}),
	}).Create(block)

	if statement.Error != nil {
		return translate(statement.Error, "indexed block")
	}

	return nil
}

func (r *indexedBlockRepository) ByNumber(ctx context.Context, chainID int64, number int64) (domain.IndexedBlock, error) {
	var block domain.IndexedBlock

	statement := r.db.WithContext(ctx).
		Where("chain_id = ? AND block_number = ?", chainID, number).
		First(&block)

	if statement.Error != nil {
		return domain.IndexedBlock{}, translate(statement.Error, "indexed block")
	}

	return block, nil
}

func (r *indexedBlockRepository) DeleteFrom(ctx context.Context, chainID int64, number int64) (int64, error) {
	statement := r.db.WithContext(ctx).
		Where("chain_id = ? AND block_number >= ?", chainID, number).
		Delete(&domain.IndexedBlock{})

	if statement.Error != nil {
		return 0, translate(statement.Error, "indexed block")
	}

	return statement.RowsAffected, nil
}

func (r *indexedBlockRepository) RecordReorg(ctx context.Context, event *domain.ReorgEvent) error {
	if event == nil {
		return domain.ErrInvalidInput
	}

	now := time.Now().UTC()
	event.DetectedAt = now
	event.CreatedAt = now

	if statement := r.db.WithContext(ctx).Create(event); statement.Error != nil {
		return translate(statement.Error, "reorg event")
	}

	return nil
}

type positionRepository struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) domain.PositionRepository {
	return &positionRepository{db: db}
}

func (r *positionRepository) Upsert(ctx context.Context, position *domain.Position) error {
	if position == nil {
		return domain.ErrInvalidInput
	}

	now := time.Now().UTC()
	position.UpdatedAt = now

	if position.CreatedAt.IsZero() {
		position.CreatedAt = now
	}

	statement := r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}, {Name: "market_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"supply_shares",
			"collateral_amount",
			"debt_scaled",
			"health_factor_bps",
			"collateral_value",
			"debt_value",
			"is_liquidatable",
			"liquidatable_since",
			"valued_at_price",
			"valued_at_block",
			"valued_at",
			"last_event_block",
			"updated_at",
		}),
	}).Create(position)

	if statement.Error != nil {
		return translate(statement.Error, "position")
	}

	return nil
}

func (r *positionRepository) Liquidatable(ctx context.Context, marketID int64, limit int) ([]domain.Position, error) {
	positions := make([]domain.Position, 0, boundedLimit(limit, 25, 100))

	statement := r.db.WithContext(ctx).
		Where("market_id = ? AND is_liquidatable = TRUE AND debt_scaled > 0", marketID).
		Order("health_factor_bps ASC, id ASC").
		Limit(boundedLimit(limit, 25, 100)).
		Preload("User").
		Find(&positions)

	if statement.Error != nil {
		return nil, translate(statement.Error, "liquidatable positions")
	}

	return positions, nil
}
