package indexer

import (
	"context"
	"fmt"

	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
	"github.com/farhapartex/lending-platform/core-service/pkg/jsonb"
)

type IngestorParams struct {
	ChainID      int64
	Market       domain.Market
	Users        domain.UserRepository
	Events       domain.ProtocolEventRepository
	Transactions domain.TransactionRepository
	Liquidations domain.LiquidationRepository
}

type Ingestor struct {
	params IngestorParams
}

func NewIngestor(params IngestorParams) *Ingestor {
	return &Ingestor{params: params}
}

func (i *Ingestor) Apply(ctx context.Context, event DecodedEvent, blockTime time.Time) error {
	actor, err := i.params.Users.EnsureByAddress(ctx, i.params.ChainID, event.Actor)
	if err != nil {
		return fmt.Errorf("recording the actor %s failed: %w", event.Actor, err)
	}

	payload, err := jsonb.Marshal(event.Payload)
	if err != nil {
		return fmt.Errorf("encoding the payload of %s failed: %w", event.TxHash, err)
	}

	marketID := i.params.Market.ID

	record := domain.ProtocolEvent{
		ChainID:         i.params.ChainID,
		MarketID:        &marketID,
		EventType:       event.Type,
		ContractAddress: event.ContractAddress,
		BlockNumber:     int64(event.BlockNumber),
		BlockHash:       event.BlockHash,
		BlockTime:       blockTime,
		TxHash:          event.TxHash,
		TxIndex:         int32(event.TxIndex),
		LogIndex:        int32(event.LogIndex),
		ActorUserID:     &actor.ID,
		Payload:         payload,
	}

	if err := i.params.Events.Insert(ctx, &record); err != nil {
		return fmt.Errorf("recording the event at %s log %d failed: %w", event.TxHash, event.LogIndex, err)
	}

	if err := i.applyTransaction(ctx, event, record, actor, blockTime); err != nil {
		return err
	}

	if event.Liquidation == nil {
		return nil
	}

	return i.applyLiquidation(ctx, event, record, actor, blockTime)
}

func (i *Ingestor) applyTransaction(
	ctx context.Context,
	event DecodedEvent,
	record domain.ProtocolEvent,
	actor domain.User,
	blockTime time.Time,
) error {
	if event.Kind == "" || event.Amount == nil || event.Amount.Sign() <= 0 {
		return nil
	}

	transaction := domain.UserTransaction{
		EventID:              record.ID,
		UserID:               actor.ID,
		MarketID:             i.params.Market.ID,
		AssetID:              i.assetFor(event.Kind),
		Kind:                 event.Kind,
		Amount:               bigmath.FromBig(event.Amount),
		HealthFactorAfterBps: event.HealthFactorBps,
		BlockNumber:          int64(event.BlockNumber),
		BlockTime:            blockTime,
		TxHash:               event.TxHash,
		LogIndex:             int32(event.LogIndex),
	}

	if err := i.params.Transactions.Insert(ctx, &transaction); err != nil {
		return fmt.Errorf("recording the transaction at %s failed: %w", event.TxHash, err)
	}

	return nil
}

func (i *Ingestor) applyLiquidation(
	ctx context.Context,
	event DecodedEvent,
	record domain.ProtocolEvent,
	borrower domain.User,
	blockTime time.Time,
) error {
	detail := event.Liquidation

	liquidator, err := i.params.Users.EnsureByAddress(ctx, i.params.ChainID, detail.Liquidator)
	if err != nil {
		return fmt.Errorf("recording the liquidator %s failed: %w", detail.Liquidator, err)
	}

	liquidation := domain.Liquidation{
		EventID:               record.ID,
		MarketID:              i.params.Market.ID,
		BorrowerUserID:        borrower.ID,
		LiquidatorUserID:      liquidator.ID,
		DebtRepaid:            bigmath.FromBig(orZero(detail.DebtRepaid)),
		CollateralSeized:      bigmath.FromBig(orZero(detail.CollateralSeized)),
		BonusAmount:           bigmath.FromBig(orZero(detail.BonusAmount)),
		HealthFactorBeforeBps: event.HealthFactorBps,
		TriggerPrice:          bigmath.FromBig(orZero(detail.TriggerPrice)),
		TriggerPriceDecimals:  detail.TriggerPriceDecimals,
		ShortfallAmount:       bigmath.FromBig(orZero(detail.Shortfall)),
		BlockNumber:           int64(event.BlockNumber),
		BlockTime:             blockTime,
		TxHash:                event.TxHash,
	}

	if err := i.params.Liquidations.Insert(ctx, &liquidation); err != nil {
		return fmt.Errorf("recording the liquidation at %s failed: %w", event.TxHash, err)
	}

	return nil
}

func (i *Ingestor) assetFor(kind domain.TransactionKind) int64 {
	switch kind {
	case domain.TransactionKindCollateralAdded, domain.TransactionKindCollateralWithdrawn:
		return i.params.Market.CollateralAssetID
	default:
		return i.params.Market.DebtAssetID
	}
}
