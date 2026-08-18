package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrNotFound      = errors.New("record was not found")
	ErrAlreadyExists = errors.New("record already exists")
	ErrInvalidInput  = errors.New("input is not usable")
)

type AssetRepository interface {
	ByID(ctx context.Context, id int64) (Asset, error)
	ByAddress(ctx context.Context, chainID int64, address string) (Asset, error)
	List(ctx context.Context, chainID int64) ([]Asset, error)
	Upsert(ctx context.Context, asset *Asset) error
}

type MarketRepository interface {
	ByID(ctx context.Context, id int64) (Market, error)
	ByPoolAddress(ctx context.Context, chainID int64, poolAddress string) (Market, error)
	List(ctx context.Context, chainID int64) ([]Market, error)
	Upsert(ctx context.Context, market *Market) error
}

type MarketSnapshotRepository interface {
	Insert(ctx context.Context, snapshot *MarketSnapshot) error
	Latest(ctx context.Context, marketID int64) (MarketSnapshot, error)
	Since(ctx context.Context, marketID int64, since time.Time, limit int) ([]MarketSnapshot, error)
}
