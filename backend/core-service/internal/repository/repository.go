package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/ethaddr"
)

func translate(err error, what string) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("%w: %s", domain.ErrNotFound, what)
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return fmt.Errorf("%w: %s", domain.ErrAlreadyExists, what)
	}

	return fmt.Errorf("%s: %w", what, err)
}

func normalizeAddress(raw string, field string) (string, error) {
	normalized, err := ethaddr.Normalize(raw)
	if err != nil {
		return "", fmt.Errorf("%w: %s %s", domain.ErrInvalidInput, field, err)
	}

	return normalized, nil
}

func boundedLimit(limit int, fallback int, ceiling int) int {
	if limit < 1 {
		return fallback
	}

	if limit > ceiling {
		return ceiling
	}

	return limit
}

func requirePositiveID(id int64, field string) error {
	if id < 1 {
		return fmt.Errorf("%w: %s must be a positive integer", domain.ErrInvalidInput, field)
	}

	return nil
}
