package indexer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/chain"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

type PositionTrackerParams struct {
	ChainID   int64
	Market    domain.Market
	Accounts  *chain.AccountReader
	Users     domain.UserRepository
	Positions domain.PositionRepository
}

type PositionTracker struct {
	params PositionTrackerParams
}

func NewPositionTracker(params PositionTrackerParams) *PositionTracker {
	return &PositionTracker{params: params}
}

func (t *PositionTracker) Refresh(ctx context.Context, address string, atBlock int64) error {
	user, err := t.params.Users.EnsureByAddress(ctx, t.params.ChainID, address)
	if err != nil {
		return fmt.Errorf("finding the holder %s failed: %w", address, err)
	}

	view, err := t.params.Accounts.ReadAccount(ctx, address)
	if err != nil {
		return err
	}

	now := time.Now().UTC()

	position := domain.Position{
		UserID:           user.ID,
		MarketID:         t.params.Market.ID,
		SupplyShares:     bigmath.FromBig(orZero(view.SupplyShares)),
		CollateralAmount: bigmath.FromBig(orZero(view.CollateralAmount)),
		DebtScaled:       bigmath.FromBig(orZero(view.DebtAmount)),
		IsLiquidatable:   view.IsLiquidatable,
		LastEventBlock:   atBlock,
	}

	if !view.PriceStale {
		collateralValue := bigmath.FromBig(orZero(view.CollateralValue))
		debtValue := bigmath.FromBig(orZero(view.DebtValue))
		price := bigmath.FromBig(orZero(view.CollateralPrice))

		position.CollateralValue = &collateralValue
		position.DebtValue = &debtValue
		position.ValuedAtPrice = &price
		position.ValuedAtBlock = &atBlock
		position.ValuedAt = &now
		position.HealthFactorBps = narrowHealth(view.HealthFactorBps, view.DebtAmount)
	}

	if view.IsLiquidatable {
		position.LiquidatableSince = &now
	}

	if err := t.params.Positions.Upsert(ctx, &position); err != nil {
		return fmt.Errorf("recording the position of %s failed: %w", address, err)
	}

	return nil
}

func (t *PositionTracker) Backfill(ctx context.Context, atBlock int64) (int, error) {
	addresses, err := t.params.Users.ListAddresses(ctx, 0)
	if err != nil {
		return 0, err
	}

	refreshed := 0

	for _, address := range addresses {
		if err := t.Refresh(ctx, address, atBlock); err != nil {
			return refreshed, err
		}

		refreshed++
	}

	return refreshed, nil
}

func narrowHealth(healthFactorBps *big.Int, debtAmount *big.Int) *int32 {
	if debtAmount == nil || debtAmount.Sign() <= 0 {
		return nil
	}

	return narrowBps(healthFactorBps)
}
