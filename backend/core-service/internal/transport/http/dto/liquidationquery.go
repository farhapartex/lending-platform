package dto

import (
	"net/url"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

const ParamMarket = "market"

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

	return domain.LiquidationListRequest{
		MarketID: marketID,
		After:    after,
		Limit:    limit,
	}, nil
}
