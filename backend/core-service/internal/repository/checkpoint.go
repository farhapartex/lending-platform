package repository

import (
	"context"
	"strings"

	"gorm.io/gorm"

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
