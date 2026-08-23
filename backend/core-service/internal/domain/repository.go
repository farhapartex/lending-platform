package domain

import (
	"context"
	"errors"
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
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

type TransactionQuery struct {
	UserID int64
	Kinds  []TransactionKind
	From   *time.Time
	To     *time.Time
	After  cursor.Key
	Limit  int
}

type TransactionRepository interface {
	List(ctx context.Context, query TransactionQuery) ([]UserTransaction, error)
	ByID(ctx context.Context, userID int64, id int64) (UserTransaction, error)
	Insert(ctx context.Context, transaction *UserTransaction) error
}

type CheckpointRepository interface {
	ByStream(ctx context.Context, stream string) (IndexerCheckpoint, error)
}

type UserRepository interface {
	ByAddress(ctx context.Context, address string) (User, error)
	EnsureByAddress(ctx context.Context, chainID int64, address string) (User, error)
}

type MarketSnapshotRepository interface {
	Insert(ctx context.Context, snapshot *MarketSnapshot) error
	Latest(ctx context.Context, marketID int64) (MarketSnapshot, error)
	Since(ctx context.Context, marketID int64, since time.Time, limit int) ([]MarketSnapshot, error)
}
