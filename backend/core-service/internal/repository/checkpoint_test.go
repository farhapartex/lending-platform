package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/repository"
)

func addCheckpoint(t *testing.T, tx *gorm.DB, stream string, block int64) domain.IndexerCheckpoint {
	t.Helper()

	checkpoint := domain.IndexerCheckpoint{
		StreamName:            stream,
		ChainID:               testChainID,
		LastProcessedBlock:    block,
		LastProcessedLogIndex: 3,
		UpdatedAt:             time.Now().UTC(),
	}

	if err := tx.Create(&checkpoint).Error; err != nil {
		t.Fatalf("could not insert the checkpoint: %v", err)
	}

	return checkpoint
}

func TestCheckpointByStreamReturnsTheStoredProgress(t *testing.T) {
	tx := newTx(t)
	stored := addCheckpoint(t, tx, domain.IndexerStreamProtocolEvents, 4_218)

	found, err := repository.NewCheckpointRepository(tx).
		ByStream(context.Background(), domain.IndexerStreamProtocolEvents)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.ID != stored.ID {
		t.Fatalf("expected checkpoint %d, got %d", stored.ID, found.ID)
	}

	if found.LastProcessedBlock != 4_218 {
		t.Fatalf("expected block 4218, got %d", found.LastProcessedBlock)
	}

	if found.LastProcessedLogIndex != 3 {
		t.Fatalf("expected log index 3, got %d", found.LastProcessedLogIndex)
	}
}

func TestCheckpointByStreamIsScopedToTheStream(t *testing.T) {
	tx := newTx(t)
	addCheckpoint(t, tx, domain.IndexerStreamProtocolEvents, 4_218)
	addCheckpoint(t, tx, "snapshots", 99)

	repo := repository.NewCheckpointRepository(tx)

	events, err := repo.ByStream(context.Background(), domain.IndexerStreamProtocolEvents)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if events.LastProcessedBlock != 4_218 {
		t.Fatalf("expected the events stream, got block %d", events.LastProcessedBlock)
	}

	snapshots, err := repo.ByStream(context.Background(), "snapshots")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if snapshots.LastProcessedBlock != 99 {
		t.Fatalf("expected the snapshots stream, got block %d", snapshots.LastProcessedBlock)
	}
}

func TestCheckpointByStreamReportsAnUnknownStreamAsNotFound(t *testing.T) {
	tx := newTx(t)

	_, err := repository.NewCheckpointRepository(tx).ByStream(context.Background(), "nothing-writes-this")

	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestCheckpointByStreamTrimsTheStreamName(t *testing.T) {
	tx := newTx(t)
	addCheckpoint(t, tx, domain.IndexerStreamProtocolEvents, 4_218)

	found, err := repository.NewCheckpointRepository(tx).
		ByStream(context.Background(), "  "+domain.IndexerStreamProtocolEvents+"  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.LastProcessedBlock != 4_218 {
		t.Fatalf("expected surrounding spaces to be ignored, got block %d", found.LastProcessedBlock)
	}
}

func TestCheckpointByStreamRejectsAnEmptyStream(t *testing.T) {
	tx := newTx(t)
	repo := repository.NewCheckpointRepository(tx)

	for _, stream := range []string{"", "   ", "\t"} {
		_, err := repo.ByStream(context.Background(), stream)

		if !errors.Is(err, domain.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput for %q, got %v", stream, err)
		}
	}
}

func TestCheckpointByStreamSurfacesADatabaseFailure(t *testing.T) {
	tx := newTx(t)

	if err := tx.Exec("DROP TABLE indexer_checkpoints").Error; err != nil {
		t.Fatalf("could not drop the table: %v", err)
	}

	_, err := repository.NewCheckpointRepository(tx).
		ByStream(context.Background(), domain.IndexerStreamProtocolEvents)

	if err == nil {
		t.Fatal("expected a database failure to surface")
	}

	if errors.Is(err, domain.ErrNotFound) || errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected a raw failure rather than a domain error, got %v", err)
	}
}
