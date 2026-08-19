package domain

import (
	"context"
	"math/big"
)

type MarketView struct {
	TotalSupplied           *big.Int
	TotalBorrowed           *big.Int
	AvailableLiquidity      *big.Int
	UtilizationBps          int32
	SupplyRatePerSecond     *big.Int
	BorrowRatePerSecond     *big.Int
	SupplyAprBps            int32
	BorrowAprBps            int32
	SupplyIndex             *big.Int
	BorrowIndex             *big.Int
	MaxLtvBps               int32
	LiquidationThresholdBps int32
	LiquidationBonusBps     int32
	KinkUtilizationBps      int32
	ReserveFactorBps        int32
	MinDeposit              *big.Int
	AccruedReserves         *big.Int
	DepositsPaused          bool
	BorrowPaused            bool
}

type MarketViewReader interface {
	ReadMarketView(ctx context.Context) (MarketView, error)
}
