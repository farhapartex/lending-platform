package service

import (
	"context"
	"fmt"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/ethaddr"
)

type TransactionServiceParams struct {
	Users        domain.UserRepository
	Transactions domain.TransactionRepository
}

type transactionService struct {
	users        domain.UserRepository
	transactions domain.TransactionRepository
}

func NewTransactionService(params TransactionServiceParams) domain.TransactionService {
	return &transactionService{
		users:        params.Users,
		transactions: params.Transactions,
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
