package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrChainUnreachable = errors.New("chain node is unreachable")
	ErrChainIDMismatch  = errors.New("chain id does not match the configured network")
	ErrBlockNotFound    = errors.New("block was not found")
)

type BlockRef struct {
	Number uint64
	Hash   string
	Time   time.Time
}

func (b BlockRef) IsZero() bool {
	return b.Number == 0 && b.Hash == ""
}

type ChainReader interface {
	ChainID(ctx context.Context) (int64, error)
	HeadBlock(ctx context.Context) (uint64, error)
	BlockByNumber(ctx context.Context, number uint64) (BlockRef, error)
}
