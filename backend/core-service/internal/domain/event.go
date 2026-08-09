package domain

import (
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/jsonb"
)

type EventType string

const (
	EventTypeDeposit             EventType = "deposit"
	EventTypeWithdraw            EventType = "withdraw"
	EventTypeBorrow              EventType = "borrow"
	EventTypeRepay               EventType = "repay"
	EventTypeCollateralAdded     EventType = "collateral_added"
	EventTypeCollateralWithdrawn EventType = "collateral_withdrawn"
	EventTypeLiquidation         EventType = "liquidation"
	EventTypeInterestAccrued     EventType = "interest_accrued"
	EventTypeParameterChanged    EventType = "parameter_changed"
)

type ProtocolEvent struct {
	ID              int64          `gorm:"column:id;primaryKey;autoIncrement"`
	ChainID         int64          `gorm:"column:chain_id;not null"`
	MarketID        *int64         `gorm:"column:market_id"`
	EventType       EventType      `gorm:"column:event_type;not null"`
	ContractAddress string         `gorm:"column:contract_address;not null"`
	BlockNumber     int64          `gorm:"column:block_number;not null"`
	BlockHash       string         `gorm:"column:block_hash;not null"`
	BlockTime       time.Time      `gorm:"column:block_time;not null"`
	TxHash          string         `gorm:"column:tx_hash;not null"`
	TxIndex         int32          `gorm:"column:tx_index;not null"`
	LogIndex        int32          `gorm:"column:log_index;not null"`
	ActorUserID     *int64         `gorm:"column:actor_user_id"`
	Payload         jsonb.Document `gorm:"column:payload;type:jsonb;not null"`
	CreatedAt       time.Time      `gorm:"column:created_at;not null"`

	Market    *Market `gorm:"foreignKey:MarketID"`
	ActorUser *User   `gorm:"foreignKey:ActorUserID"`
}

func (ProtocolEvent) TableName() string {
	return "protocol_events"
}
