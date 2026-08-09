package domain

import (
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
	"github.com/farhapartex/lending-platform/core-service/pkg/jsonb"
)

type FaucetRequestStatus string

const (
	FaucetRequestStatusPending   FaucetRequestStatus = "pending"
	FaucetRequestStatusSending   FaucetRequestStatus = "sending"
	FaucetRequestStatusCompleted FaucetRequestStatus = "completed"
	FaucetRequestStatusFailed    FaucetRequestStatus = "failed"
)

type FaucetRequest struct {
	ID          int64               `gorm:"column:id;primaryKey;autoIncrement"`
	UserID      int64               `gorm:"column:user_id;not null"`
	AssetID     int64               `gorm:"column:asset_id;not null"`
	Amount      bigmath.Int         `gorm:"column:amount;type:numeric(78,0);not null"`
	Status      FaucetRequestStatus `gorm:"column:status;not null"`
	TxHash      *string             `gorm:"column:tx_hash"`
	RequestIP   *string             `gorm:"column:request_ip;type:inet"`
	Error       *string             `gorm:"column:error"`
	CreatedAt   time.Time           `gorm:"column:created_at;not null"`
	CompletedAt *time.Time          `gorm:"column:completed_at"`

	User  *User  `gorm:"foreignKey:UserID"`
	Asset *Asset `gorm:"foreignKey:AssetID"`
}

func (FaucetRequest) TableName() string {
	return "faucet_requests"
}

type IdempotencyKey struct {
	ID             int64          `gorm:"column:id;primaryKey;autoIncrement"`
	Key            string         `gorm:"column:key;not null"`
	Scope          string         `gorm:"column:scope;not null"`
	RequestHash    string         `gorm:"column:request_hash;not null"`
	ResponseStatus *int16         `gorm:"column:response_status"`
	ResponseBody   jsonb.Document `gorm:"column:response_body;type:jsonb"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null"`
	ExpiresAt      time.Time      `gorm:"column:expires_at;not null"`
}

func (IdempotencyKey) TableName() string {
	return "idempotency_keys"
}

type AuditLog struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement"`
	Actor      string         `gorm:"column:actor;not null"`
	Action     string         `gorm:"column:action;not null"`
	EntityType string         `gorm:"column:entity_type;not null"`
	EntityID   *int64         `gorm:"column:entity_id"`
	Before     jsonb.Document `gorm:"column:before;type:jsonb"`
	After      jsonb.Document `gorm:"column:after;type:jsonb"`
	RequestIP  *string        `gorm:"column:request_ip;type:inet"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null"`
}

func (AuditLog) TableName() string {
	return "audit_logs"
}
