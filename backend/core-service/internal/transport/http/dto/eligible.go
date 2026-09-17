package dto

import (
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

const priceValueDecimals int16 = 8

type EligiblePositionResponse struct {
	Borrower         string  `json:"borrower"`
	CollateralAmount Amount  `json:"collateral_amount"`
	DebtAmount       Amount  `json:"debt_amount"`
	CollateralValue  *Amount `json:"collateral_value"`
	DebtValue        *Amount `json:"debt_value"`
	HealthFactorBps  *int32  `json:"health_factor_bps"`
	LastEventBlock   int64   `json:"last_event_block"`
}

type EligibleListResponse struct {
	Items []EligiblePositionResponse `json:"items"`
	AsOf  AsOfResponse               `json:"as_of"`
}

func NewEligibleListResponse(
	page domain.EligiblePage,
	collateralDecimals int16,
	collateralSymbol string,
	debtDecimals int16,
	debtSymbol string,
) EligibleListResponse {
	items := make([]EligiblePositionResponse, 0, len(page.Items))

	for _, entry := range page.Items {
		item := EligiblePositionResponse{
			Borrower:         entry.Borrower.AddressChecksum,
			CollateralAmount: NewAmount(entry.Position.CollateralAmount, collateralDecimals, collateralSymbol),
			DebtAmount:       NewAmount(entry.Position.DebtScaled, debtDecimals, debtSymbol),
			HealthFactorBps:  entry.Position.HealthFactorBps,
			LastEventBlock:   entry.Position.LastEventBlock,
		}

		if entry.Borrower.AddressChecksum == "" {
			item.Borrower = entry.Borrower.Address
		}

		if entry.Position.CollateralValue != nil {
			value := NewAmount(*entry.Position.CollateralValue, priceValueDecimals, UsdSymbol)
			item.CollateralValue = &value
		}

		if entry.Position.DebtValue != nil {
			value := NewAmount(*entry.Position.DebtValue, priceValueDecimals, UsdSymbol)
			item.DebtValue = &value
		}

		items = append(items, item)
	}

	return EligibleListResponse{Items: items, AsOf: newAsOfResponse(page.AsOf)}
}
