package repository

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/ethaddr"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) ByAddress(ctx context.Context, address string) (domain.User, error) {
	normalized, err := normalizeAddress(address, "user address")
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User

	query := r.db.WithContext(ctx).First(&user, "address = ?", normalized)
	if query.Error != nil {
		return domain.User{}, translate(query.Error, "user by address")
	}

	return user, nil
}

func (r *userRepository) EnsureByAddress(
	ctx context.Context,
	chainID int64,
	address string,
) (domain.User, error) {
	normalized, checksum, err := ethaddr.NormalizeWithChecksum(address)
	if err != nil {
		return domain.User{}, fmt.Errorf("%w: user address %s", domain.ErrInvalidInput, err)
	}

	now := time.Now().UTC()

	user := domain.User{
		Address:         normalized,
		AddressChecksum: checksum,
		FirstSeenAt:     now,
		LastSeenAt:      &now,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	insert := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "address"}},
			DoUpdates: clause.AssignmentColumns([]string{"last_seen_at", "updated_at"}),
		}).
		Create(&user)

	if insert.Error != nil {
		return domain.User{}, translate(insert.Error, "user ensure")
	}

	return user, nil
}
