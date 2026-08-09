package domain

import (
	"time"
)

type Asset struct {
	ID                int64  `gorm:"column:id;primaryKey;autoIncrement"`
	ChainID           int64  `gorm:"column:chain_id;not null"`
	Address           string `gorm:"column:address;not null"`
	AddressChecksum   string `gorm:"column:address_checksum;not null"`
	Symbol            string `gorm:"column:symbol;not null"`
	Name              string `gorm:"column:name;not null"`
	Decimals          int16  `gorm:"column:decimals;not null"`
	IsCollateral      bool   `gorm:"column:is_collateral;not null"`
	IsBorrowable      bool   `gorm:"column:is_borrowable;not null"`
	PriceFeedAddress  *string
	PriceFeedDecimals *int16    `gorm:"column:price_feed_decimals"`
	CreatedAt         time.Time `gorm:"column:created_at;not null"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null"`
}

func (Asset) TableName() string {
	return "assets"
}
