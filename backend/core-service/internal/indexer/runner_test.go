package indexer_test

import (
	"context"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/farhapartex/lending-platform/core-service/internal/config"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/indexer"
)

const testChainID int64 = 31337

type fakeChain struct {
	head   uint64
	blocks map[uint64]domain.BlockRef
	logs   []types.Log
}

func (f *fakeChain) HeadBlock(_ context.Context) (uint64, error) {
	return f.head, nil
}

func (f *fakeChain) SafeHeadBlock(_ context.Context, confirmations uint64) (uint64, error) {
	if f.head < confirmations {
		return 0, nil
	}

	return f.head - confirmations, nil
}

func (f *fakeChain) BlockByNumber(_ context.Context, number uint64) (domain.BlockRef, error) {
	block, ok := f.blocks[number]
	if !ok {
		return domain.BlockRef{}, domain.ErrBlockNotFound
	}

	return block, nil
}

func (f *fakeChain) Logs(_ context.Context, _ uint64, _ uint64, _ []common.Address) ([]types.Log, error) {
	return f.logs, nil
}

type fakeEvents struct {
	deletedFrom []int64
}

func (f *fakeEvents) Insert(_ context.Context, _ *domain.ProtocolEvent) error {
	return nil
}

func (f *fakeEvents) DeleteFrom(_ context.Context, _ int64, blockNumber int64) (int64, error) {
	f.deletedFrom = append(f.deletedFrom, blockNumber)

	return 0, nil
}

type fakeBlocks struct {
	stored      map[int64]domain.IndexedBlock
	deletedFrom []int64
	reorgs      []domain.ReorgEvent
}

func (f *fakeBlocks) Insert(_ context.Context, block *domain.IndexedBlock) error {
	f.stored[block.BlockNumber] = *block

	return nil
}

func (f *fakeBlocks) ByNumber(_ context.Context, _ int64, number int64) (domain.IndexedBlock, error) {
	block, ok := f.stored[number]
	if !ok {
		return domain.IndexedBlock{}, domain.ErrNotFound
	}

	return block, nil
}

func (f *fakeBlocks) DeleteFrom(_ context.Context, _ int64, number int64) (int64, error) {
	f.deletedFrom = append(f.deletedFrom, number)

	return 0, nil
}

func (f *fakeBlocks) RecordReorg(_ context.Context, event *domain.ReorgEvent) error {
	f.reorgs = append(f.reorgs, *event)

	return nil
}

type fakeCheckpoints struct {
	checkpoint domain.IndexerCheckpoint
	found      bool
}

func (f *fakeCheckpoints) ByStream(_ context.Context, _ string) (domain.IndexerCheckpoint, error) {
	if !f.found {
		return domain.IndexerCheckpoint{}, domain.ErrNotFound
	}

	return f.checkpoint, nil
}

func (f *fakeCheckpoints) Save(_ context.Context, checkpoint *domain.IndexerCheckpoint) error {
	f.checkpoint = *checkpoint
	f.found = true

	return nil
}

func newRunner(chain *fakeChain, events *fakeEvents, blocks *fakeBlocks, checkpoints *fakeCheckpoints) *indexer.Runner {
	decoder, err := indexer.NewDecoder(indexer.ContractSet{
		Pool:               poolAddress,
		Vault:              vaultAddress,
		Controller:         controllerAddress,
		LiquidationManager: managerAddress,
	})
	if err != nil {
		panic(err)
	}

	return indexer.NewRunner(indexer.RunnerParams{
		Chain: config.ChainConfig{
			ChainID:       testChainID,
			Confirmations: 1,
			BatchSize:     100,
			PollInterval:  time.Second,
		},
		Client:      chain,
		Decoder:     decoder,
		Ingestor:    indexer.NewIngestor(indexer.IngestorParams{ChainID: testChainID}),
		Events:      events,
		Blocks:      blocks,
		Checkpoints: checkpoints,
		Logger:      slog.New(slog.NewTextHandler(io.Discard, nil)),
	})
}

func TestTickRewindsWhenTheChainIsShorterThanTheRecordedProgress(t *testing.T) {
	chain := &fakeChain{head: 45, blocks: map[uint64]domain.BlockRef{}}
	events := &fakeEvents{}
	blocks := &fakeBlocks{stored: map[int64]domain.IndexedBlock{}}
	checkpoints := &fakeCheckpoints{
		found:      true,
		checkpoint: domain.IndexerCheckpoint{StreamName: domain.IndexerStreamProtocolEvents, ChainID: testChainID, LastProcessedBlock: 56},
	}

	if err := newRunner(chain, events, blocks, checkpoints).Tick(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(events.deletedFrom) != 1 || events.deletedFrom[0] != 45 {
		t.Fatalf("expected events from the head onward to be removed, got %v", events.deletedFrom)
	}

	if checkpoints.checkpoint.LastProcessedBlock != 44 {
		t.Fatalf("expected the checkpoint to fall behind the head, got %d", checkpoints.checkpoint.LastProcessedBlock)
	}

	if len(blocks.reorgs) != 1 {
		t.Fatalf("expected the rewind to be recorded once, got %d", len(blocks.reorgs))
	}
}

func TestTickRewindsWhenTheStoredBlockHashNoLongerMatches(t *testing.T) {
	chain := &fakeChain{
		head: 100,
		blocks: map[uint64]domain.BlockRef{
			60: {Number: 60, Hash: "0xnew", Time: time.Unix(1, 0)},
			99: {Number: 99, Hash: "0x99", Time: time.Unix(2, 0)},
		},
	}
	events := &fakeEvents{}
	blocks := &fakeBlocks{
		stored: map[int64]domain.IndexedBlock{
			60: {ChainID: testChainID, BlockNumber: 60, BlockHash: "0xold"},
		},
	}
	checkpoints := &fakeCheckpoints{
		found:      true,
		checkpoint: domain.IndexerCheckpoint{StreamName: domain.IndexerStreamProtocolEvents, ChainID: testChainID, LastProcessedBlock: 60},
	}

	if err := newRunner(chain, events, blocks, checkpoints).Tick(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(events.deletedFrom) != 1 || events.deletedFrom[0] != 28 {
		t.Fatalf("expected a rewind of the lookback window, got %v", events.deletedFrom)
	}
}

func TestTickLeavesAMatchingCheckpointAlone(t *testing.T) {
	chain := &fakeChain{
		head: 100,
		blocks: map[uint64]domain.BlockRef{
			60: {Number: 60, Hash: "0xsame", Time: time.Unix(1, 0)},
			99: {Number: 99, Hash: "0x99", Time: time.Unix(2, 0)},
		},
	}
	events := &fakeEvents{}
	blocks := &fakeBlocks{
		stored: map[int64]domain.IndexedBlock{
			60: {ChainID: testChainID, BlockNumber: 60, BlockHash: "0xsame"},
		},
	}
	checkpoints := &fakeCheckpoints{
		found:      true,
		checkpoint: domain.IndexerCheckpoint{StreamName: domain.IndexerStreamProtocolEvents, ChainID: testChainID, LastProcessedBlock: 60},
	}

	if err := newRunner(chain, events, blocks, checkpoints).Tick(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(events.deletedFrom) != 0 {
		t.Fatalf("expected no rewind when the hash still matches, got %v", events.deletedFrom)
	}

	if checkpoints.checkpoint.LastProcessedBlock != 99 {
		t.Fatalf("expected progress up to the safe head, got %d", checkpoints.checkpoint.LastProcessedBlock)
	}
}
