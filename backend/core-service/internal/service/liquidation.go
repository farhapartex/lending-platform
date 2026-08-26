package service

import (
	"context"
	"fmt"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
)

const (
	DefaultLiquidationPageSize = 25
	MaxLiquidationPageSize     = 100
)

type LiquidationServiceParams struct {
	Liquidations domain.LiquidationRepository
	Checkpoints  domain.CheckpointRepository
	Now          func() time.Time
}

type liquidationService struct {
	liquidations domain.LiquidationRepository
	checkpoints  domain.CheckpointRepository
	now          func() time.Time
}

func NewLiquidationService(params LiquidationServiceParams) domain.LiquidationService {
	return &liquidationService{
		liquidations: params.Liquidations,
		checkpoints:  params.Checkpoints,
		now:          clockOr(params.Now),
	}
}

func (s *liquidationService) List(
	ctx context.Context,
	request domain.LiquidationListRequest,
) (domain.LiquidationPage, error) {
	if request.MarketID != nil && *request.MarketID < 1 {
		return domain.LiquidationPage{}, fmt.Errorf("%w: market id must be positive", domain.ErrInvalidInput)
	}

	pageSize := boundedSize(request.Limit, DefaultLiquidationPageSize, MaxLiquidationPageSize)

	asOf, err := indexedAt(ctx, s.checkpoints, s.now)
	if err != nil {
		return domain.LiquidationPage{}, err
	}

	found, err := s.liquidations.List(ctx, domain.LiquidationQuery{
		MarketID: request.MarketID,
		After:    request.After,
		Limit:    pageSize + 1,
	})
	if err != nil {
		return domain.LiquidationPage{}, err
	}

	items, next := trimToPage(found, pageSize, liquidationKey)

	return domain.LiquidationPage{Items: items, NextCursor: next, AsOf: asOf}, nil
}

func (s *liquidationService) ByID(ctx context.Context, id int64) (domain.Liquidation, error) {
	if id < 1 {
		return domain.Liquidation{}, fmt.Errorf("%w: liquidation id must be positive", domain.ErrInvalidInput)
	}

	return s.liquidations.ByID(ctx, id)
}

func liquidationKey(liquidation domain.Liquidation) cursor.Key {
	return cursor.Key{Time: liquidation.BlockTime, ID: liquidation.ID}
}
