package domain

import (
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/jsonb"
)

type NotificationKind string

const (
	NotificationKindHealthWarning NotificationKind = "health_warning"
	NotificationKindLiquidation   NotificationKind = "liquidation"
	NotificationKindRateChange    NotificationKind = "rate_change"
	NotificationKindAnnouncement  NotificationKind = "announcement"
)

type NotificationSeverity string

const (
	NotificationSeverityInfo     NotificationSeverity = "info"
	NotificationSeverityCaution  NotificationSeverity = "caution"
	NotificationSeverityCritical NotificationSeverity = "critical"
)

type EmailStatus string

const (
	EmailStatusPending   EmailStatus = "pending"
	EmailStatusSending   EmailStatus = "sending"
	EmailStatusSent      EmailStatus = "sent"
	EmailStatusFailed    EmailStatus = "failed"
	EmailStatusCancelled EmailStatus = "cancelled"
)

type Notification struct {
	ID        int64                `gorm:"column:id;primaryKey;autoIncrement"`
	UserID    int64                `gorm:"column:user_id;not null"`
	Kind      NotificationKind     `gorm:"column:kind;not null"`
	Severity  NotificationSeverity `gorm:"column:severity;not null"`
	Title     string               `gorm:"column:title;not null"`
	Body      string               `gorm:"column:body;not null"`
	Payload   jsonb.Document       `gorm:"column:payload;type:jsonb"`
	DedupeKey *string              `gorm:"column:dedupe_key"`
	ReadAt    *time.Time           `gorm:"column:read_at"`
	CreatedAt time.Time            `gorm:"column:created_at;not null"`

	User *User `gorm:"foreignKey:UserID"`
}

func (Notification) TableName() string {
	return "notifications"
}

type EmailMessage struct {
	ID                int64          `gorm:"column:id;primaryKey;autoIncrement"`
	UserID            *int64         `gorm:"column:user_id"`
	NotificationID    *int64         `gorm:"column:notification_id"`
	ToEmail           string         `gorm:"column:to_email;not null"`
	Template          string         `gorm:"column:template;not null"`
	Payload           jsonb.Document `gorm:"column:payload;type:jsonb;not null"`
	Status            EmailStatus    `gorm:"column:status;not null"`
	Attempts          int16          `gorm:"column:attempts;not null"`
	LastError         *string        `gorm:"column:last_error"`
	ProviderMessageID *string        `gorm:"column:provider_message_id"`
	ScheduledAt       time.Time      `gorm:"column:scheduled_at;not null"`
	SentAt            *time.Time     `gorm:"column:sent_at"`
	CreatedAt         time.Time      `gorm:"column:created_at;not null"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;not null"`

	User         *User         `gorm:"foreignKey:UserID"`
	Notification *Notification `gorm:"foreignKey:NotificationID"`
}

func (EmailMessage) TableName() string {
	return "email_messages"
}
