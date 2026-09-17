package dto

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/ethaddr"
)

const (
	ParamMarket   = "market"
	ParamBorrower = "borrower"
)

func ParseLiquidationListRequest(
	marketID *int64,
	values url.Values,
) (domain.LiquidationListRequest, error) {
	after, err := parseCursor(values)
	if err != nil {
		return domain.LiquidationListRequest{}, err
	}

	limit, err := ParseLimit(values)
	if err != nil {
		return domain.LiquidationListRequest{}, err
	}

	borrower := strings.TrimSpace(values.Get(ParamBorrower))

	if borrower != "" {
		normalized, err := ethaddr.Normalize(borrower)
		if err != nil {
			return domain.LiquidationListRequest{}, fmt.Errorf(
				"%w: %s is not a usable address", domain.ErrInvalidInput, ParamBorrower,
			)
		}

		borrower = normalized
	}

	return domain.LiquidationListRequest{
		MarketID:        marketID,
		BorrowerAddress: borrower,
		After:           after,
		Limit:           limit,
	}, nil
}
