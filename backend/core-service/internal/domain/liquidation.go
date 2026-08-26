package domain

import (
	"context"
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
)

type Liquidation struct {
	ID                    int64       `gorm:"column:id;primaryKey;autoIncrement"`
	EventID               int64       `gorm:"column:event_id;not null"`
	MarketID              int64       `gorm:"column:market_id;not null"`
	BorrowerUserID        int64       `gorm:"column:borrower_user_id;not null"`
	LiquidatorUserID      int64       `gorm:"column:liquidator_user_id;not null"`
	DebtRepaid            bigmath.Int `gorm:"column:debt_repaid;type:numeric(78,0);not null"`
	CollateralSeized      bigmath.Int `gorm:"column:collateral_seized;type:numeric(78,0);not null"`
	BonusAmount           bigmath.Int `gorm:"column:bonus_amount;type:numeric(78,0);not null"`
	HealthFactorBeforeBps *int32      `gorm:"column:health_factor_before_bps"`
	TriggerPrice          bigmath.Int `gorm:"column:trigger_price;type:numeric(78,0);not null"`
	TriggerPriceDecimals  int16       `gorm:"column:trigger_price_decimals;not null"`
	ShortfallAmount       bigmath.Int `gorm:"column:shortfall_amount;type:numeric(78,0);not null"`
	BlockNumber           int64       `gorm:"column:block_number;not null"`
	BlockTime             time.Time   `gorm:"column:block_time;not null"`
	TxHash                string      `gorm:"column:tx_hash;not null"`
	CreatedAt             time.Time   `gorm:"column:created_at;not null"`

	Event      *ProtocolEvent `gorm:"foreignKey:EventID"`
	Market     *Market        `gorm:"foreignKey:MarketID"`
	Borrower   *User          `gorm:"foreignKey:BorrowerUserID"`
	Liquidator *User          `gorm:"foreignKey:LiquidatorUserID"`
}

func (Liquidation) TableName() string {
	return "liquidations"
}

type LiquidationQuery struct {
	MarketID *int64
	After    cursor.Key
	Limit    int
}

type LiquidationListRequest struct {
	MarketID *int64
	After    cursor.Key
	Limit    int
}

type LiquidationPage struct {
	Items      []Liquidation
	NextCursor cursor.Key
	AsOf       IndexedAt
}

type LiquidationRepository interface {
	List(ctx context.Context, query LiquidationQuery) ([]Liquidation, error)
	ByID(ctx context.Context, id int64) (Liquidation, error)
	Insert(ctx context.Context, liquidation *Liquidation) error
}

type LiquidationService interface {
	List(ctx context.Context, request LiquidationListRequest) (LiquidationPage, error)
	ByID(ctx context.Context, id int64) (Liquidation, error)
}
