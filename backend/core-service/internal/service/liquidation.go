package service

import (
	"context"
	"errors"
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
	Positions    domain.PositionRepository
	Users        domain.UserRepository
	Checkpoints  domain.CheckpointRepository
	Now          func() time.Time
}

type liquidationService struct {
	liquidations domain.LiquidationRepository
	positions    domain.PositionRepository
	users        domain.UserRepository
	checkpoints  domain.CheckpointRepository
	now          func() time.Time
}

func NewLiquidationService(params LiquidationServiceParams) domain.LiquidationService {
	return &liquidationService{
		liquidations: params.Liquidations,
		positions:    params.Positions,
		users:        params.Users,
		checkpoints:  params.Checkpoints,
		now:          clockOr(params.Now),
	}
}

func (s *liquidationService) Eligible(
	ctx context.Context,
	request domain.EligibleRequest,
) (domain.EligiblePage, error) {
	if s.positions == nil {
		return domain.EligiblePage{}, fmt.Errorf("%w: positions are not available", domain.ErrNotFound)
	}

	if request.MarketID != nil && *request.MarketID < 1 {
		return domain.EligiblePage{}, fmt.Errorf("%w: market id must be positive", domain.ErrInvalidInput)
	}

	asOf, err := indexedAt(ctx, s.checkpoints, s.now)
	if err != nil {
		return domain.EligiblePage{}, err
	}

	pageSize := boundedSize(request.Limit, DefaultLiquidationPageSize, MaxLiquidationPageSize)

	found, err := s.positions.Liquidatable(ctx, request.MarketID, pageSize)
	if err != nil {
		return domain.EligiblePage{}, err
	}

	items := make([]domain.EligiblePosition, 0, len(found))

	for _, position := range found {
		item := domain.EligiblePosition{Position: position}

		if position.User != nil {
			item.Borrower = *position.User
		}

		items = append(items, item)
	}

	return domain.EligiblePage{Items: items, AsOf: asOf}, nil
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

	var borrowerID *int64

	if request.BorrowerAddress != "" {
		borrower, err := s.users.ByAddress(ctx, request.BorrowerAddress)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return domain.LiquidationPage{Items: []domain.Liquidation{}, AsOf: asOf}, nil
			}

			return domain.LiquidationPage{}, err
		}

		borrowerID = &borrower.ID
	}

	found, err := s.liquidations.List(ctx, domain.LiquidationQuery{
		MarketID:   request.MarketID,
		BorrowerID: borrowerID,
		After:      request.After,
		Limit:      pageSize + 1,
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
