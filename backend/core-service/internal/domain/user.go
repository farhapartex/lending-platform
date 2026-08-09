package domain

import (
	"time"
)

type User struct {
	ID              int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Address         string     `gorm:"column:address;not null"`
	AddressChecksum string     `gorm:"column:address_checksum;not null"`
	Nickname        *string    `gorm:"column:nickname"`
	Email           *string    `gorm:"column:email"`
	EmailVerifiedAt *time.Time `gorm:"column:email_verified_at"`
	FirstSeenAt     time.Time  `gorm:"column:first_seen_at;not null"`
	LastSeenAt      *time.Time `gorm:"column:last_seen_at"`
	CreatedAt       time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt       time.Time  `gorm:"column:updated_at;not null"`
}

func (User) TableName() string {
	return "users"
}

type UserNotificationPreference struct {
	ID                      int64     `gorm:"column:id;primaryKey;autoIncrement"`
	UserID                  int64     `gorm:"column:user_id;not null"`
	HealthAlertThresholdBps int32     `gorm:"column:health_alert_threshold_bps;not null"`
	InAppEnabled            bool      `gorm:"column:in_app_enabled;not null"`
	EmailEnabled            bool      `gorm:"column:email_enabled;not null"`
	RateChangeEnabled       bool      `gorm:"column:rate_change_enabled;not null"`
	CreatedAt               time.Time `gorm:"column:created_at;not null"`
	UpdatedAt               time.Time `gorm:"column:updated_at;not null"`

	User *User `gorm:"foreignKey:UserID"`
}

func (UserNotificationPreference) TableName() string {
	return "user_notification_preferences"
}

type EmailVerificationToken struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement"`
	UserID     int64      `gorm:"column:user_id;not null"`
	Email      string     `gorm:"column:email;not null"`
	TokenHash  string     `gorm:"column:token_hash;not null"`
	ExpiresAt  time.Time  `gorm:"column:expires_at;not null"`
	ConsumedAt *time.Time `gorm:"column:consumed_at"`
	CreatedAt  time.Time  `gorm:"column:created_at;not null"`

	User *User `gorm:"foreignKey:UserID"`
}

func (EmailVerificationToken) TableName() string {
	return "email_verification_tokens"
}
