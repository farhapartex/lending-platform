package domain

import (
	"context"
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
)

type TransactionKind string

const (
	TransactionKindDeposit             TransactionKind = "deposit"
	TransactionKindWithdraw            TransactionKind = "withdraw"
	TransactionKindBorrow              TransactionKind = "borrow"
	TransactionKindRepay               TransactionKind = "repay"
	TransactionKindCollateralAdded     TransactionKind = "collateral_added"
	TransactionKindCollateralWithdrawn TransactionKind = "collateral_withdrawn"
	TransactionKindLiquidation         TransactionKind = "liquidation"
)

type UserTransaction struct {
	ID                   int64           `gorm:"column:id;primaryKey;autoIncrement"`
	EventID              int64           `gorm:"column:event_id;not null"`
	UserID               int64           `gorm:"column:user_id;not null"`
	MarketID             int64           `gorm:"column:market_id;not null"`
	AssetID              int64           `gorm:"column:asset_id;not null"`
	Kind                 TransactionKind `gorm:"column:kind;not null"`
	Amount               bigmath.Int     `gorm:"column:amount;type:numeric(78,0);not null"`
	HealthFactorAfterBps *int32          `gorm:"column:health_factor_after_bps"`
	BlockNumber          int64           `gorm:"column:block_number;not null"`
	BlockTime            time.Time       `gorm:"column:block_time;not null"`
	TxHash               string          `gorm:"column:tx_hash;not null"`
	LogIndex             int32           `gorm:"column:log_index;not null"`
	CreatedAt            time.Time       `gorm:"column:created_at;not null"`

	Event  *ProtocolEvent `gorm:"foreignKey:EventID"`
	User   *User          `gorm:"foreignKey:UserID"`
	Market *Market        `gorm:"foreignKey:MarketID"`
	Asset  *Asset         `gorm:"foreignKey:AssetID"`
}

func (UserTransaction) TableName() string {
	return "user_transactions"
}

var AllTransactionKinds = []TransactionKind{
	TransactionKindDeposit,
	TransactionKindWithdraw,
	TransactionKindBorrow,
	TransactionKindRepay,
	TransactionKindCollateralAdded,
	TransactionKindCollateralWithdrawn,
	TransactionKindLiquidation,
}

func ParseTransactionKind(raw string) (TransactionKind, bool) {
	for _, kind := range AllTransactionKinds {
		if string(kind) == raw {
			return kind, true
		}
	}

	return "", false
}

type TransactionListRequest struct {
	Address string
	Kinds   []TransactionKind
	From    *time.Time
	To      *time.Time
	After   cursor.Key
	Limit   int
}

type TransactionPage struct {
	Items      []UserTransaction
	NextCursor cursor.Key
	AsOf       IndexedAt
}

type IndexedAt struct {
	Block *int64
	Time  time.Time
}

type TransactionService interface {
	ByID(ctx context.Context, address string, id int64) (UserTransaction, error)
	List(ctx context.Context, request TransactionListRequest) (TransactionPage, error)
	RecentActivity(ctx context.Context, address string, limit int) (TransactionPage, error)
}
