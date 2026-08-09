package domain

import (
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

type PriceSource string

const (
	PriceSourceChainlink PriceSource = "chainlink"
	PriceSourceFallback  PriceSource = "fallback"
	PriceSourceManual    PriceSource = "manual"
)

type MarketSnapshot struct {
	ID                 int64       `gorm:"column:id;primaryKey;autoIncrement"`
	MarketID           int64       `gorm:"column:market_id;not null"`
	CapturedAt         time.Time   `gorm:"column:captured_at;not null"`
	BlockNumber        int64       `gorm:"column:block_number;not null"`
	TotalSupplied      bigmath.Int `gorm:"column:total_supplied;type:numeric(78,0);not null"`
	TotalBorrowed      bigmath.Int `gorm:"column:total_borrowed;type:numeric(78,0);not null"`
	AvailableLiquidity bigmath.Int `gorm:"column:available_liquidity;type:numeric(78,0);not null"`
	UtilizationBps     int32       `gorm:"column:utilization_bps;not null"`
	SupplyRateBps      int32       `gorm:"column:supply_rate_bps;not null"`
	BorrowRateBps      int32       `gorm:"column:borrow_rate_bps;not null"`
	SupplyIndex        bigmath.Int `gorm:"column:supply_index;type:numeric(78,0);not null"`
	BorrowIndex        bigmath.Int `gorm:"column:borrow_index;type:numeric(78,0);not null"`
	PositionsAtRisk    int32       `gorm:"column:positions_at_risk;not null"`
	CreatedAt          time.Time   `gorm:"column:created_at;not null"`

	Market *Market `gorm:"foreignKey:MarketID"`
}

func (MarketSnapshot) TableName() string {
	return "market_snapshots"
}

type PriceObservation struct {
	ID            int64        `gorm:"column:id;primaryKey;autoIncrement"`
	AssetID       int64        `gorm:"column:asset_id;not null"`
	Price         bigmath.Int  `gorm:"column:price;type:numeric(78,0);not null"`
	PriceDecimals int16        `gorm:"column:price_decimals;not null"`
	Source        PriceSource  `gorm:"column:source;not null"`
	RoundID       *bigmath.Int `gorm:"column:round_id;type:numeric(78,0)"`
	ObservedAt    time.Time    `gorm:"column:observed_at;not null"`
	FetchedAt     time.Time    `gorm:"column:fetched_at;not null"`
	IsStale       bool         `gorm:"column:is_stale;not null"`
	CreatedAt     time.Time    `gorm:"column:created_at;not null"`

	Asset *Asset `gorm:"foreignKey:AssetID"`
}

func (PriceObservation) TableName() string {
	return "price_observations"
}
