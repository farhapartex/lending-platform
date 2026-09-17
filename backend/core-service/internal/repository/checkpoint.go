package repository

import (
	"context"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

type checkpointRepository struct {
	db *gorm.DB
}

func NewCheckpointRepository(db *gorm.DB) domain.CheckpointRepository {
	return &checkpointRepository{db: db}
}

func (r *checkpointRepository) ByStream(ctx context.Context, stream string) (domain.IndexerCheckpoint, error) {
	trimmed := strings.TrimSpace(stream)
	if trimmed == "" {
		return domain.IndexerCheckpoint{}, domain.ErrInvalidInput
	}

	var checkpoint domain.IndexerCheckpoint

	statement := r.db.WithContext(ctx).First(&checkpoint, "stream_name = ?", trimmed)

	if statement.Error != nil {
		return domain.IndexerCheckpoint{}, translate(statement.Error, "indexer checkpoint")
	}

	return checkpoint, nil
}

func (r *checkpointRepository) Save(ctx context.Context, checkpoint *domain.IndexerCheckpoint) error {
	if checkpoint == nil || strings.TrimSpace(checkpoint.StreamName) == "" {
		return domain.ErrInvalidInput
	}

	checkpoint.UpdatedAt = time.Now().UTC()

	statement := r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "stream_name"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"chain_id",
			"last_processed_block",
			"last_processed_log_index",
			"last_processed_block_hash",
			"updated_at",
		}),
	}).Create(checkpoint)

	if statement.Error != nil {
		return translate(statement.Error, "indexer checkpoint")
	}

	return nil
}
