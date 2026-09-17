package indexer

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sort"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/chain"
	"github.com/farhapartex/lending-platform/core-service/internal/config"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

const reorgLookback = 32

type RunnerParams struct {
	Chain       config.ChainConfig
	Client      *chain.Client
	Decoder     *Decoder
	Ingestor    *Ingestor
	Events      domain.ProtocolEventRepository
	Blocks      domain.IndexedBlockRepository
	Checkpoints domain.CheckpointRepository
	Logger      *slog.Logger
}

type Runner struct {
	params RunnerParams
}

func NewRunner(params RunnerParams) *Runner {
	return &Runner{params: params}
}

func (r *Runner) Run(ctx context.Context) {
	ticker := time.NewTicker(r.params.Chain.PollInterval)
	defer ticker.Stop()

	r.params.Logger.Info(
		"indexer started",
		slog.Int64("chain_id", r.params.Chain.ChainID),
		slog.Uint64("start_block", r.params.Chain.StartBlock),
		slog.Duration("poll_interval", r.params.Chain.PollInterval),
	)

	for {
		if err := r.Tick(ctx); err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}

			r.params.Logger.Error("indexer pass failed", slog.String("error", err.Error()))
		}

		select {
		case <-ctx.Done():
			r.params.Logger.Info("indexer stopped")

			return
		case <-ticker.C:
		}
	}
}

func (r *Runner) Tick(ctx context.Context) error {
	checkpoint, err := r.loadCheckpoint(ctx)
	if err != nil {
		return err
	}

	if err := r.reconcile(ctx, &checkpoint); err != nil {
		return err
	}

	from := uint64(checkpoint.LastProcessedBlock) + 1
	if checkpoint.LastProcessedBlock == 0 && r.params.Chain.StartBlock > 0 {
		from = r.params.Chain.StartBlock
	}

	safeHead, err := r.params.Client.SafeHeadBlock(ctx, r.params.Chain.Confirmations)
	if err != nil {
		return err
	}

	if from > safeHead {
		return nil
	}

	to := from + r.params.Chain.BatchSize - 1
	if to > safeHead {
		to = safeHead
	}

	logs, err := r.params.Client.Logs(ctx, from, to, r.params.Decoder.Contracts().Addresses())
	if err != nil {
		return err
	}

	sort.SliceStable(logs, func(left, right int) bool {
		if logs[left].BlockNumber != logs[right].BlockNumber {
			return logs[left].BlockNumber < logs[right].BlockNumber
		}

		return logs[left].Index < logs[right].Index
	})

	blockTimes := map[uint64]time.Time{}
	applied := 0

	for _, log := range logs {
		event, ok := r.params.Decoder.Decode(log)
		if !ok {
			continue
		}

		blockTime, err := r.blockTime(ctx, blockTimes, log.BlockNumber)
		if err != nil {
			return err
		}

		if err := r.params.Ingestor.Apply(ctx, event, blockTime); err != nil {
			return err
		}

		applied++
	}

	if err := r.recordProgress(ctx, &checkpoint, to); err != nil {
		return err
	}

	if applied > 0 {
		r.params.Logger.Info(
			"indexed protocol events",
			slog.Uint64("from_block", from),
			slog.Uint64("to_block", to),
			slog.Int("events", applied),
		)
	}

	return nil
}

func (r *Runner) loadCheckpoint(ctx context.Context) (domain.IndexerCheckpoint, error) {
	checkpoint, err := r.params.Checkpoints.ByStream(ctx, domain.IndexerStreamProtocolEvents)
	if err == nil {
		return checkpoint, nil
	}

	if !errors.Is(err, domain.ErrNotFound) {
		return domain.IndexerCheckpoint{}, err
	}

	return domain.IndexerCheckpoint{
		StreamName:            domain.IndexerStreamProtocolEvents,
		ChainID:               r.params.Chain.ChainID,
		LastProcessedBlock:    0,
		LastProcessedLogIndex: -1,
	}, nil
}

func (r *Runner) reconcile(ctx context.Context, checkpoint *domain.IndexerCheckpoint) error {
	if checkpoint.LastProcessedBlock <= 0 {
		return nil
	}

	stored, err := r.params.Blocks.ByNumber(ctx, r.params.Chain.ChainID, checkpoint.LastProcessedBlock)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil
		}

		return err
	}

	onChain, err := r.params.Client.BlockByNumber(ctx, uint64(stored.BlockNumber))
	if err != nil {
		if errors.Is(err, domain.ErrBlockNotFound) {
			return r.rollback(ctx, checkpoint, stored.BlockNumber, "the indexed block is no longer on the chain")
		}

		return err
	}

	if onChain.Hash == stored.BlockHash {
		return nil
	}

	rewindTo := stored.BlockNumber - reorgLookback
	if rewindTo < 1 {
		rewindTo = 1
	}

	return r.rollback(ctx, checkpoint, rewindTo, "the indexed block hash no longer matches the chain")
}

func (r *Runner) rollback(
	ctx context.Context,
	checkpoint *domain.IndexerCheckpoint,
	fromBlock int64,
	reason string,
) error {
	removedEvents, err := r.params.Events.DeleteFrom(ctx, r.params.Chain.ChainID, fromBlock)
	if err != nil {
		return err
	}

	if _, err := r.params.Blocks.DeleteFrom(ctx, r.params.Chain.ChainID, fromBlock); err != nil {
		return err
	}

	notes := reason

	if err := r.params.Blocks.RecordReorg(ctx, &domain.ReorgEvent{
		ChainID:      r.params.Chain.ChainID,
		FromBlock:    fromBlock,
		ToBlock:      checkpoint.LastProcessedBlock,
		RowsReverted: int32(removedEvents),
		Notes:        &notes,
	}); err != nil {
		return err
	}

	checkpoint.LastProcessedBlock = fromBlock - 1
	checkpoint.LastProcessedLogIndex = -1
	checkpoint.LastProcessedBlockHash = nil

	if err := r.params.Checkpoints.Save(ctx, checkpoint); err != nil {
		return err
	}

	r.params.Logger.Warn(
		"rewound the indexer after a chain reorganisation",
		slog.Int64("from_block", fromBlock),
		slog.Int64("to_block", checkpoint.LastProcessedBlock),
		slog.Int64("events_removed", removedEvents),
		slog.String("reason", reason),
	)

	return nil
}

func (r *Runner) blockTime(ctx context.Context, cache map[uint64]time.Time, number uint64) (time.Time, error) {
	if known, ok := cache[number]; ok {
		return known, nil
	}

	block, err := r.params.Client.BlockByNumber(ctx, number)
	if err != nil {
		return time.Time{}, err
	}

	cache[number] = block.Time

	return block.Time, nil
}

func (r *Runner) recordProgress(ctx context.Context, checkpoint *domain.IndexerCheckpoint, to uint64) error {
	block, err := r.params.Client.BlockByNumber(ctx, to)
	if err != nil {
		return err
	}

	if err := r.params.Blocks.Insert(ctx, &domain.IndexedBlock{
		ChainID:     r.params.Chain.ChainID,
		BlockNumber: int64(block.Number),
		BlockHash:   block.Hash,
		ParentHash:  block.Hash,
		BlockTime:   block.Time,
		IsFinalized: true,
	}); err != nil {
		return fmt.Errorf("recording block %d failed: %w", block.Number, err)
	}

	checkpoint.ChainID = r.params.Chain.ChainID
	checkpoint.LastProcessedBlock = int64(block.Number)
	checkpoint.LastProcessedLogIndex = -1
	hash := block.Hash
	checkpoint.LastProcessedBlockHash = &hash

	return r.params.Checkpoints.Save(ctx, checkpoint)
}
