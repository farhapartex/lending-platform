package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
	"github.com/farhapartex/lending-platform/core-service/pkg/ethaddr"
)

const (
	DefaultTransactionPageSize = 25
	MaxTransactionPageSize     = 100

	DefaultActivitySize = 5
	MaxActivitySize     = 20
)

type TransactionServiceParams struct {
	Users        domain.UserRepository
	Transactions domain.TransactionRepository
	Checkpoints  domain.CheckpointRepository
	Now          func() time.Time
}

type transactionService struct {
	users        domain.UserRepository
	transactions domain.TransactionRepository
	checkpoints  domain.CheckpointRepository
	now          func() time.Time
}

func NewTransactionService(params TransactionServiceParams) domain.TransactionService {
	now := params.Now
	if now == nil {
		now = time.Now
	}

	return &transactionService{
		users:        params.Users,
		transactions: params.Transactions,
		checkpoints:  params.Checkpoints,
		now:          now,
	}
}

func (s *transactionService) ByID(
	ctx context.Context,
	address string,
	id int64,
) (domain.UserTransaction, error) {
	normalized, err := ethaddr.Normalize(address)
	if err != nil {
		return domain.UserTransaction{}, fmt.Errorf("%w: address %s", domain.ErrInvalidInput, err)
	}

	if id < 1 {
		return domain.UserTransaction{}, fmt.Errorf("%w: transaction id must be positive", domain.ErrInvalidInput)
	}

	user, err := s.users.ByAddress(ctx, normalized)
	if err != nil {
		return domain.UserTransaction{}, err
	}

	return s.transactions.ByID(ctx, user.ID, id)
}

func (s *transactionService) List(
	ctx context.Context,
	request domain.TransactionListRequest,
) (domain.TransactionPage, error) {
	normalized, err := ethaddr.Normalize(request.Address)
	if err != nil {
		return domain.TransactionPage{}, fmt.Errorf("%w: address %s", domain.ErrInvalidInput, err)
	}

	if err := validateKinds(request.Kinds); err != nil {
		return domain.TransactionPage{}, err
	}

	if err := validateWindow(request.From, request.To); err != nil {
		return domain.TransactionPage{}, err
	}

	pageSize := boundedSize(request.Limit, DefaultTransactionPageSize, MaxTransactionPageSize)

	owner, err := s.resolveOwner(ctx, normalized)
	if err != nil {
		return domain.TransactionPage{}, err
	}

	if !owner.found {
		return emptyPage(owner.asOf), nil
	}

	found, err := s.transactions.List(ctx, domain.TransactionQuery{
		UserID: owner.userID,
		Kinds:  request.Kinds,
		From:   request.From,
		To:     request.To,
		After:  request.After,
		Limit:  pageSize + 1,
	})
	if err != nil {
		return domain.TransactionPage{}, err
	}

	items, next := trimToPage(found, pageSize)

	return domain.TransactionPage{Items: items, NextCursor: next, AsOf: owner.asOf}, nil
}

func (s *transactionService) RecentActivity(
	ctx context.Context,
	address string,
	limit int,
) (domain.TransactionPage, error) {
	normalized, err := ethaddr.Normalize(address)
	if err != nil {
		return domain.TransactionPage{}, fmt.Errorf("%w: address %s", domain.ErrInvalidInput, err)
	}

	owner, err := s.resolveOwner(ctx, normalized)
	if err != nil {
		return domain.TransactionPage{}, err
	}

	if !owner.found {
		return emptyPage(owner.asOf), nil
	}

	items, err := s.transactions.List(ctx, domain.TransactionQuery{
		UserID: owner.userID,
		Limit:  boundedSize(limit, DefaultActivitySize, MaxActivitySize),
	})
	if err != nil {
		return domain.TransactionPage{}, err
	}

	return domain.TransactionPage{Items: items, AsOf: owner.asOf}, nil
}

type pageOwner struct {
	userID int64
	asOf   domain.IndexedAt
	found  bool
}

func (s *transactionService) resolveOwner(ctx context.Context, address string) (pageOwner, error) {
	asOf, err := s.indexedAt(ctx)
	if err != nil {
		return pageOwner{}, err
	}

	user, err := s.users.ByAddress(ctx, address)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return pageOwner{asOf: asOf}, nil
		}

		return pageOwner{}, err
	}

	return pageOwner{userID: user.ID, asOf: asOf, found: true}, nil
}

func (s *transactionService) indexedAt(ctx context.Context) (domain.IndexedAt, error) {
	now := s.now().UTC()

	if s.checkpoints == nil {
		return domain.IndexedAt{Time: now}, nil
	}

	checkpoint, err := s.checkpoints.ByStream(ctx, domain.IndexerStreamProtocolEvents)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return domain.IndexedAt{Time: now}, nil
		}

		return domain.IndexedAt{}, err
	}

	block := checkpoint.LastProcessedBlock

	return domain.IndexedAt{Block: &block, Time: checkpoint.UpdatedAt.UTC()}, nil
}

func emptyPage(asOf domain.IndexedAt) domain.TransactionPage {
	return domain.TransactionPage{Items: []domain.UserTransaction{}, AsOf: asOf}
}

func validateKinds(kinds []domain.TransactionKind) error {
	for _, kind := range kinds {
		if _, ok := domain.ParseTransactionKind(string(kind)); !ok {
			return fmt.Errorf("%w: %q is not a transaction kind", domain.ErrInvalidInput, kind)
		}
	}

	return nil
}

func validateWindow(from *time.Time, to *time.Time) error {
	if from == nil || to == nil {
		return nil
	}

	if from.After(*to) {
		return fmt.Errorf("%w: from must not be later than to", domain.ErrInvalidInput)
	}

	return nil
}

func boundedSize(limit int, fallback int, ceiling int) int {
	if limit < 1 {
		return fallback
	}

	if limit > ceiling {
		return ceiling
	}

	return limit
}

func trimToPage(found []domain.UserTransaction, pageSize int) ([]domain.UserTransaction, cursor.Key) {
	if len(found) <= pageSize {
		return found, cursor.Key{}
	}

	items := found[:pageSize]
	last := items[len(items)-1]

	return items, cursor.Key{Time: last.BlockTime, ID: last.ID}
}
