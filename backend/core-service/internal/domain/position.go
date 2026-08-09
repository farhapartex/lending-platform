package domain

import (
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

type Position struct {
	ID                int64        `gorm:"column:id;primaryKey;autoIncrement"`
	UserID            int64        `gorm:"column:user_id;not null"`
	MarketID          int64        `gorm:"column:market_id;not null"`
	SupplyShares      bigmath.Int  `gorm:"column:supply_shares;type:numeric(78,0);not null"`
	CollateralAmount  bigmath.Int  `gorm:"column:collateral_amount;type:numeric(78,0);not null"`
	DebtScaled        bigmath.Int  `gorm:"column:debt_scaled;type:numeric(78,0);not null"`
	LastSupplyIndex   *bigmath.Int `gorm:"column:last_supply_index;type:numeric(78,0)"`
	LastBorrowIndex   *bigmath.Int `gorm:"column:last_borrow_index;type:numeric(78,0)"`
	HealthFactorBps   *int32       `gorm:"column:health_factor_bps"`
	CollateralValue   *bigmath.Int `gorm:"column:collateral_value;type:numeric(78,0)"`
	DebtValue         *bigmath.Int `gorm:"column:debt_value;type:numeric(78,0)"`
	IsLiquidatable    bool         `gorm:"column:is_liquidatable;not null"`
	LiquidatableSince *time.Time   `gorm:"column:liquidatable_since"`
	ValuedAtPrice     *bigmath.Int `gorm:"column:valued_at_price;type:numeric(78,0)"`
	ValuedAtBlock     *int64       `gorm:"column:valued_at_block"`
	ValuedAt          *time.Time   `gorm:"column:valued_at"`
	LastEventBlock    int64        `gorm:"column:last_event_block;not null"`
	CreatedAt         time.Time    `gorm:"column:created_at;not null"`
	UpdatedAt         time.Time    `gorm:"column:updated_at;not null"`

	User   *User   `gorm:"foreignKey:UserID"`
	Market *Market `gorm:"foreignKey:MarketID"`
}

func (Position) TableName() string {
	return "positions"
}
