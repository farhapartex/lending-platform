package dto

import (
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

const UsdSymbol = "USD"

type LiquidationResponse struct {
	ID                    string `json:"id"`
	Borrower              string `json:"borrower"`
	Liquidator            string `json:"liquidator"`
	DebtRepaid            Amount `json:"debt_repaid"`
	CollateralSeized      Amount `json:"collateral_seized"`
	BonusValue            Amount `json:"bonus_value"`
	ShortfallValue        Amount `json:"shortfall_value"`
	HealthFactorBeforeBps *int32 `json:"health_factor_before_bps"`
	TriggerPrice          Amount `json:"trigger_price"`
	TxHash                string `json:"tx_hash"`
	Block                 int64  `json:"block"`
	BlockTime             string `json:"block_time"`
}

type LiquidationListResponse struct {
	Items      []LiquidationResponse `json:"items"`
	NextCursor *string               `json:"next_cursor"`
	AsOf       AsOfResponse          `json:"as_of"`
}

func NewLiquidationResponse(liquidation domain.Liquidation, publicID string) LiquidationResponse {
	debtDecimals, debtSymbol := assetUnits(debtAssetOf(liquidation.Market))
	collateralDecimals, collateralSymbol := assetUnits(collateralAssetOf(liquidation.Market))
	valueDecimals := liquidation.TriggerPriceDecimals

	return LiquidationResponse{
		ID:                    publicID,
		Borrower:              addressOf(liquidation.Borrower),
		Liquidator:            addressOf(liquidation.Liquidator),
		DebtRepaid:            NewAmount(liquidation.DebtRepaid, debtDecimals, debtSymbol),
		CollateralSeized:      NewAmount(liquidation.CollateralSeized, collateralDecimals, collateralSymbol),
		BonusValue:            NewAmount(liquidation.BonusAmount, valueDecimals, UsdSymbol),
		ShortfallValue:        NewAmount(liquidation.ShortfallAmount, valueDecimals, UsdSymbol),
		HealthFactorBeforeBps: liquidation.HealthFactorBeforeBps,
		TriggerPrice:          NewAmount(liquidation.TriggerPrice, valueDecimals, UsdSymbol),
		TxHash:                liquidation.TxHash,
		Block:                 liquidation.BlockNumber,
		BlockTime:             liquidation.BlockTime.UTC().Format(time.RFC3339),
	}
}

func NewLiquidationListResponse(
	page domain.LiquidationPage,
	publicID PublicIDFunc,
) (LiquidationListResponse, error) {
	items := make([]LiquidationResponse, 0, len(page.Items))

	for _, liquidation := range page.Items {
		masked, err := publicID(liquidation.ID)
		if err != nil {
			return LiquidationListResponse{}, err
		}

		items = append(items, NewLiquidationResponse(liquidation, masked))
	}

	return LiquidationListResponse{
		Items:      items,
		NextCursor: encodeNextCursor(page.NextCursor),
		AsOf:       newAsOfResponse(page.AsOf),
	}, nil
}

func debtAssetOf(market *domain.Market) *domain.Asset {
	if market == nil {
		return nil
	}

	return market.DebtAsset
}

func collateralAssetOf(market *domain.Market) *domain.Asset {
	if market == nil {
		return nil
	}

	return market.CollateralAsset
}

func addressOf(user *domain.User) string {
	if user == nil {
		return ""
	}

	return user.Address
}
