package service

import (
	"context"
	"errors"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
)

func indexedAt(
	ctx context.Context,
	checkpoints domain.CheckpointRepository,
	now func() time.Time,
) (domain.IndexedAt, error) {
	moment := now().UTC()

	if checkpoints == nil {
		return domain.IndexedAt{Time: moment}, nil
	}

	checkpoint, err := checkpoints.ByStream(ctx, domain.IndexerStreamProtocolEvents)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return domain.IndexedAt{Time: moment}, nil
		}

		return domain.IndexedAt{}, err
	}

	block := checkpoint.LastProcessedBlock

	return domain.IndexedAt{Block: &block, Time: checkpoint.UpdatedAt.UTC()}, nil
}

func boundedSize(limit int, fallback int, ceiling int) int {
	if limit < 1 {
		return fallback
	}

	if limit > ceiling {
		return ceiling
	}

	return limit
}

func trimToPage[T any](found []T, pageSize int, keyOf func(T) cursor.Key) ([]T, cursor.Key) {
	if len(found) <= pageSize {
		return found, cursor.Key{}
	}

	items := found[:pageSize]

	return items, keyOf(items[len(items)-1])
}

func clockOr(now func() time.Time) func() time.Time {
	if now == nil {
		return time.Now
	}

	return now
}
