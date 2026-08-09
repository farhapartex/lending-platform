package domain

import (
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

type MarketStatus string

const (
	MarketStatusActive      MarketStatus = "active"
	MarketStatusDepositOnly MarketStatus = "deposit_only"
	MarketStatusDeprecated  MarketStatus = "deprecated"
)

type Market struct {
	ID                        int64        `gorm:"column:id;primaryKey;autoIncrement"`
	ChainID                   int64        `gorm:"column:chain_id;not null"`
	CollateralAssetID         int64        `gorm:"column:collateral_asset_id;not null"`
	DebtAssetID               int64        `gorm:"column:debt_asset_id;not null"`
	PoolAddress               string       `gorm:"column:pool_address;not null"`
	CollateralVaultAddress    string       `gorm:"column:collateral_vault_address;not null"`
	ControllerAddress         string       `gorm:"column:controller_address;not null"`
	LiquidationManagerAddress string       `gorm:"column:liquidation_manager_address;not null"`
	InterestRateModelAddress  string       `gorm:"column:interest_rate_model_address;not null"`
	OracleAdapterAddress      string       `gorm:"column:oracle_adapter_address;not null"`
	MaxLTVBps                 int32        `gorm:"column:max_ltv_bps;not null"`
	LiquidationThresholdBps   int32        `gorm:"column:liquidation_threshold_bps;not null"`
	LiquidationBonusBps       int32        `gorm:"column:liquidation_bonus_bps;not null"`
	ReserveFactorBps          int32        `gorm:"column:reserve_factor_bps;not null"`
	RecommendedLTVBps         int32        `gorm:"column:recommended_ltv_bps;not null"`
	KinkUtilizationBps        int32        `gorm:"column:kink_utilization_bps;not null"`
	MinDeposit                bigmath.Int  `gorm:"column:min_deposit;type:numeric(78,0);not null"`
	Status                    MarketStatus `gorm:"column:status;not null"`
	DeployedAtBlock           int64        `gorm:"column:deployed_at_block;not null"`
	CreatedAt                 time.Time    `gorm:"column:created_at;not null"`
	UpdatedAt                 time.Time    `gorm:"column:updated_at;not null"`

	CollateralAsset *Asset `gorm:"foreignKey:CollateralAssetID"`
	DebtAsset       *Asset `gorm:"foreignKey:DebtAssetID"`
}

func (Market) TableName() string {
	return "markets"
}

type MarketParameterHistory struct {
	ID             int64        `gorm:"column:id;primaryKey;autoIncrement"`
	MarketID       int64        `gorm:"column:market_id;not null"`
	ParameterName  string       `gorm:"column:parameter_name;not null"`
	OldValue       *bigmath.Int `gorm:"column:old_value;type:numeric(78,0)"`
	NewValue       bigmath.Int  `gorm:"column:new_value;type:numeric(78,0);not null"`
	ChangedAtBlock int64        `gorm:"column:changed_at_block;not null"`
	BlockTime      time.Time    `gorm:"column:block_time;not null"`
	TxHash         *string      `gorm:"column:tx_hash"`
	CreatedAt      time.Time    `gorm:"column:created_at;not null"`
}

func (MarketParameterHistory) TableName() string {
	return "market_parameter_history"
}
