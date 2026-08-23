package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

const (
	defaultTransactionLimit = 25
	maxTransactionLimit     = 200
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) List(
	ctx context.Context,
	query domain.TransactionQuery,
) ([]domain.UserTransaction, error) {
	if err := requirePositiveID(query.UserID, "user id"); err != nil {
		return nil, err
	}

	transactions := make([]domain.UserTransaction, 0)

	statement := r.db.WithContext(ctx).
		Preload("Asset").
		Where("user_id = ?", query.UserID)

	statement = applyKindFilter(statement, query.Kinds)
	statement = applyTimeWindow(statement, query)
	statement = applyKeysetCursor(statement, query)

	result := statement.
		Order("block_time DESC, id DESC").
		Limit(boundedLimit(query.Limit, defaultTransactionLimit, maxTransactionLimit)).
		Find(&transactions)

	if result.Error != nil {
		return nil, translate(result.Error, "transaction list")
	}

	return transactions, nil
}

func (r *transactionRepository) ByID(
	ctx context.Context,
	userID int64,
	id int64,
) (domain.UserTransaction, error) {
	if err := requirePositiveID(userID, "user id"); err != nil {
		return domain.UserTransaction{}, err
	}

	if err := requirePositiveID(id, "transaction id"); err != nil {
		return domain.UserTransaction{}, err
	}

	var transaction domain.UserTransaction

	statement := r.db.WithContext(ctx).
		Preload("Asset").
		First(&transaction, "id = ? AND user_id = ?", id, userID)

	if statement.Error != nil {
		return domain.UserTransaction{}, translate(statement.Error, "transaction by id")
	}

	return transaction, nil
}

func (r *transactionRepository) Insert(ctx context.Context, transaction *domain.UserTransaction) error {
	if transaction == nil {
		return domain.ErrInvalidInput
	}

	if err := requirePositiveID(transaction.UserID, "user id"); err != nil {
		return err
	}

	return translate(
		r.db.WithContext(ctx).Omit("Event", "User", "Market", "Asset").Create(transaction).Error,
		"transaction insert",
	)
}

func applyKindFilter(statement *gorm.DB, kinds []domain.TransactionKind) *gorm.DB {
	if len(kinds) == 0 {
		return statement
	}

	return statement.Where("kind IN ?", kinds)
}

func applyTimeWindow(statement *gorm.DB, query domain.TransactionQuery) *gorm.DB {
	if query.From != nil {
		statement = statement.Where("block_time >= ?", query.From.UTC())
	}

	if query.To != nil {
		statement = statement.Where("block_time <= ?", query.To.UTC())
	}

	return statement
}

func applyKeysetCursor(statement *gorm.DB, query domain.TransactionQuery) *gorm.DB {
	if query.After.IsZero() {
		return statement
	}

	return statement.Where(
		"(block_time, id) < (?, ?)",
		query.After.Time.UTC(),
		query.After.ID,
	)
}
