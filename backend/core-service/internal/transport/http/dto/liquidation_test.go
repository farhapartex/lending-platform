package dto_test

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"testing"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
	"github.com/farhapartex/lending-platform/core-service/pkg/queryparam"
)

const (
	borrowerAddress   = "0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc"
	liquidatorAddress = "0x70997970c51812dc3a010c7d01b50e0d17dc79c8"
)

func sampleLiquidation() domain.Liquidation {
	health := int32(9_098)

	return domain.Liquidation{
		ID:                    42,
		DebtRepaid:            bigmath.FromInt64(5_100_000_000),
		CollateralSeized:      bigmath.MustFromString("1846551724137931035"),
		BonusAmount:           bigmath.FromInt64(25_500_000_000),
		ShortfallAmount:       bigmath.FromInt64(0),
		HealthFactorBeforeBps: &health,
		TriggerPrice:          bigmath.FromInt64(220_000_000_000),
		TriggerPriceDecimals:  8,
		BlockNumber:           4_218,
		BlockTime:             listMoment,
		TxHash:                "0xabc",
		Borrower:              &domain.User{Address: borrowerAddress},
		Liquidator:            &domain.User{Address: liquidatorAddress},
		Market: &domain.Market{
			DebtAsset:       &domain.Asset{Symbol: "USDC", Decimals: 6},
			CollateralAsset: &domain.Asset{Symbol: "WETH", Decimals: 18},
		},
	}
}

func TestLiquidationResponseCarriesEveryReceiptField(t *testing.T) {
	response := dto.NewLiquidationResponse(sampleLiquidation(), "liq_abc")

	if response.ID != "liq_abc" {
		t.Fatalf("expected the masked id, got %q", response.ID)
	}

	if response.Borrower != borrowerAddress || response.Liquidator != liquidatorAddress {
		t.Fatalf("expected both parties, got %+v", response)
	}

	if response.HealthFactorBeforeBps == nil || *response.HealthFactorBeforeBps != 9_098 {
		t.Fatalf("expected the health factor before, got %v", response.HealthFactorBeforeBps)
	}

	if response.TxHash != "0xabc" || response.Block != 4_218 {
		t.Fatalf("unexpected chain details %+v", response)
	}

	if response.BlockTime != "2026-08-22T09:30:00Z" {
		t.Fatalf("unexpected block time %q", response.BlockTime)
	}
}

func TestLiquidationResponseUsesTheDebtAssetForTheRepayment(t *testing.T) {
	response := dto.NewLiquidationResponse(sampleLiquidation(), "liq_abc")

	if response.DebtRepaid.Amount != "5100000000" {
		t.Fatalf("unexpected debt amount %q", response.DebtRepaid.Amount)
	}

	if response.DebtRepaid.Decimals != 6 || response.DebtRepaid.Symbol != "USDC" {
		t.Fatalf("expected the debt asset units, got %+v", response.DebtRepaid)
	}
}

func TestLiquidationResponseUsesTheCollateralAssetForTheSeizure(t *testing.T) {
	response := dto.NewLiquidationResponse(sampleLiquidation(), "liq_abc")

	if response.CollateralSeized.Amount != "1846551724137931035" {
		t.Fatalf("unexpected collateral amount %q", response.CollateralSeized.Amount)
	}

	if response.CollateralSeized.Decimals != 18 || response.CollateralSeized.Symbol != "WETH" {
		t.Fatalf("expected the collateral asset units, got %+v", response.CollateralSeized)
	}
}

func TestLiquidationResponseTreatsBonusAndShortfallAsUsdValues(t *testing.T) {
	response := dto.NewLiquidationResponse(sampleLiquidation(), "liq_abc")

	for name, amount := range map[string]dto.Amount{
		"bonus":     response.BonusValue,
		"shortfall": response.ShortfallValue,
		"price":     response.TriggerPrice,
	} {
		if amount.Symbol != dto.UsdSymbol {
			t.Fatalf("%s is a usd value, not a token amount, got symbol %q", name, amount.Symbol)
		}

		if amount.Decimals != 8 {
			t.Fatalf("%s must carry the price scale so no client hardcodes it, got %d", name, amount.Decimals)
		}
	}

	if response.BonusValue.Amount != "25500000000" {
		t.Fatalf("unexpected bonus %q", response.BonusValue.Amount)
	}
}

func TestLiquidationResponseFollowsTheStoredPriceScale(t *testing.T) {
	liquidation := sampleLiquidation()
	liquidation.TriggerPriceDecimals = 18

	response := dto.NewLiquidationResponse(liquidation, "liq_abc")

	if response.TriggerPrice.Decimals != 18 || response.BonusValue.Decimals != 18 {
		t.Fatalf("values share the price scale, so both should follow it, got %+v", response)
	}
}

func TestLiquidationResponseSurvivesMissingRelations(t *testing.T) {
	liquidation := sampleLiquidation()
	liquidation.Market = nil
	liquidation.Borrower = nil
	liquidation.Liquidator = nil

	response := dto.NewLiquidationResponse(liquidation, "liq_abc")

	if response.Borrower != "" || response.Liquidator != "" {
		t.Fatalf("expected empty addresses rather than a panic, got %+v", response)
	}

	if response.DebtRepaid.Symbol != "" || response.DebtRepaid.Decimals != 0 {
		t.Fatalf("expected unknown units rather than a broken receipt, got %+v", response.DebtRepaid)
	}

	if response.DebtRepaid.Amount != "5100000000" {
		t.Fatalf("the amount itself is still known, got %q", response.DebtRepaid.Amount)
	}
}

func TestLiquidationResponseSurvivesAMarketWithoutAssets(t *testing.T) {
	liquidation := sampleLiquidation()
	liquidation.Market = &domain.Market{}

	response := dto.NewLiquidationResponse(liquidation, "liq_abc")

	if response.DebtRepaid.Symbol != "" || response.CollateralSeized.Symbol != "" {
		t.Fatalf("expected unknown units, got %+v", response)
	}
}

func TestLiquidationResponseKeepsANullHealthFactor(t *testing.T) {
	liquidation := sampleLiquidation()
	liquidation.HealthFactorBeforeBps = nil

	response := dto.NewLiquidationResponse(liquidation, "liq_abc")

	if response.HealthFactorBeforeBps != nil {
		t.Fatalf("expected null rather than zero, got %d", *response.HealthFactorBeforeBps)
	}
}

func liquidationPage(count int, next cursor.Key, block *int64) domain.LiquidationPage {
	items := make([]domain.Liquidation, 0, count)

	for index := 0; index < count; index++ {
		row := sampleLiquidation()
		row.ID = int64(index + 1)
		items = append(items, row)
	}

	return domain.LiquidationPage{
		Items:      items,
		NextCursor: next,
		AsOf:       domain.IndexedAt{Block: block, Time: listMoment},
	}
}

func TestLiquidationListResponseMasksEveryItem(t *testing.T) {
	response, err := dto.NewLiquidationListResponse(liquidationPage(3, cursor.Key{}, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(response.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(response.Items))
	}

	for _, item := range response.Items {
		if !strings.HasPrefix(item.ID, "txn_") {
			t.Fatalf("expected a masked id, got %q", item.ID)
		}
	}
}

func TestLiquidationListResponseEncodesTheCursorOnlyWhenPresent(t *testing.T) {
	without, err := dto.NewLiquidationListResponse(liquidationPage(1, cursor.Key{}, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if without.NextCursor != nil {
		t.Fatalf("expected a null cursor, got %q", *without.NextCursor)
	}

	with, err := dto.NewLiquidationListResponse(
		liquidationPage(1, cursor.Key{Time: listMoment, ID: 7}, nil),
		maskAsHex,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if with.NextCursor == nil {
		t.Fatal("expected a cursor")
	}

	decoded, err := cursor.Decode(*with.NextCursor)
	if err != nil || decoded.ID != 7 {
		t.Fatalf("expected the cursor to round trip, got %+v and %v", decoded, err)
	}
}

func TestLiquidationListResponseReportsAsOf(t *testing.T) {
	block := int64(4_218)

	response, err := dto.NewLiquidationListResponse(liquidationPage(1, cursor.Key{}, &block), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response.AsOf.Block == nil || *response.AsOf.Block != 4_218 {
		t.Fatalf("expected block 4218, got %v", response.AsOf.Block)
	}
}

func TestLiquidationListResponseSerialisesAnEmptyPageAsAnArray(t *testing.T) {
	response, err := dto.NewLiquidationListResponse(liquidationPage(0, cursor.Key{}, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	encoded, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("could not encode: %v", err)
	}

	if !strings.Contains(string(encoded), `"items":[]`) {
		t.Fatalf("expected an empty array rather than null, got %s", encoded)
	}
}

func TestLiquidationListResponseSurfacesAMaskingFailure(t *testing.T) {
	if _, err := dto.NewLiquidationListResponse(liquidationPage(2, cursor.Key{}, nil), failingMask); err == nil {
		t.Fatal("expected a masking failure to surface rather than emit a blank id")
	}
}

func TestParseLiquidationListRequestReadsCursorAndLimit(t *testing.T) {
	marketID := int64(7)
	key := cursor.Key{Time: listMoment, ID: 42}

	values := url.Values{}
	values.Set(dto.ParamCursor, cursor.Encode(key))
	values.Set(dto.ParamLimit, "50")

	request, err := dto.ParseLiquidationListRequest(&marketID, values)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if request.MarketID == nil || *request.MarketID != 7 {
		t.Fatalf("expected the market to pass through, got %v", request.MarketID)
	}

	if request.After.ID != 42 || request.Limit != 50 {
		t.Fatalf("unexpected request %+v", request)
	}
}

func TestParseLiquidationListRequestDefaultsToEverything(t *testing.T) {
	request, err := dto.ParseLiquidationListRequest(nil, url.Values{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if request.MarketID != nil {
		t.Fatalf("expected no market filter, got %v", request.MarketID)
	}

	if !request.After.IsZero() || request.Limit != 0 {
		t.Fatalf("expected an unfiltered first page, got %+v", request)
	}
}

func TestParseLiquidationListRequestRejectsBadInput(t *testing.T) {
	cases := map[string]url.Values{
		"bad cursor":     {dto.ParamCursor: []string{"not-a-cursor"}},
		"bad limit":      {dto.ParamLimit: []string{"many"}},
		"negative limit": {dto.ParamLimit: []string{"-5"}},
	}

	for name, values := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := dto.ParseLiquidationListRequest(nil, values)

			if !errors.Is(err, queryparam.ErrInvalid) {
				t.Fatalf("expected ErrInvalid, got %v", err)
			}

			var paramError *queryparam.ParamError
			if !errors.As(err, &paramError) || paramError.Param == "" {
				t.Fatalf("expected the error to name the parameter, got %v", err)
			}
		})
	}
}
