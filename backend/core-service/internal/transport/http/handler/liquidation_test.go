package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/handler"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
	"github.com/farhapartex/lending-platform/core-service/pkg/idmask"
)

type stubLiquidationService struct {
	page         domain.LiquidationPage
	receipt      domain.Liquidation
	listFailWith error
	byIDFailWith error
	lastRequest  domain.LiquidationListRequest
	lastID       int64
	listCount    int
	byIDCount    int
}

func (s *stubLiquidationService) List(
	_ context.Context,
	request domain.LiquidationListRequest,
) (domain.LiquidationPage, error) {
	s.lastRequest = request
	s.listCount++

	if s.listFailWith != nil {
		return domain.LiquidationPage{}, s.listFailWith
	}

	return s.page, nil
}

func (s *stubLiquidationService) ByID(_ context.Context, id int64) (domain.Liquidation, error) {
	s.lastID = id
	s.byIDCount++

	if s.byIDFailWith != nil {
		return domain.Liquidation{}, s.byIDFailWith
	}

	return s.receipt, nil
}

func newLiquidationRouter(
	t *testing.T,
	liquidations domain.LiquidationService,
) (*gin.Engine, *idmask.Masker) {
	t.Helper()

	gin.SetMode(gin.TestMode)

	masker := newMasker(t)

	liquidationHandler := handler.NewLiquidationHandler(handler.LiquidationHandlerParams{
		Liquidations: liquidations,
		Masker:       masker,
	})

	engine := gin.New()
	engine.GET("/liquidations/history", liquidationHandler.ListHistory)
	engine.GET("/liquidations/:liquidationId", liquidationHandler.GetReceipt)

	return engine, masker
}

func sampleReceipt(id int64) domain.Liquidation {
	health := int32(9_098)

	return domain.Liquidation{
		ID:                    id,
		DebtRepaid:            bigmath.FromInt64(5_100_000_000),
		CollateralSeized:      bigmath.MustFromString("1846551724137931035"),
		BonusAmount:           bigmath.FromInt64(25_500_000_000),
		ShortfallAmount:       bigmath.FromInt64(0),
		HealthFactorBeforeBps: &health,
		TriggerPrice:          bigmath.FromInt64(220_000_000_000),
		TriggerPriceDecimals:  8,
		BlockNumber:           4_218,
		BlockTime:             time.Date(2026, 8, 22, 9, 30, 0, 0, time.UTC),
		TxHash:                "0xabc",
		Borrower:              &domain.User{Address: "0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc"},
		Liquidator:            &domain.User{Address: "0x70997970c51812dc3a010c7d01b50e0d17dc79c8"},
		Market: &domain.Market{
			DebtAsset:       &domain.Asset{Symbol: "USDC", Decimals: 6},
			CollateralAsset: &domain.Asset{Symbol: "WETH", Decimals: 18},
		},
	}
}

func liquidationServicePage(count int, next cursor.Key, block *int64) domain.LiquidationPage {
	items := make([]domain.Liquidation, 0, count)

	for index := 0; index < count; index++ {
		items = append(items, sampleReceipt(int64(index+1)))
	}

	return domain.LiquidationPage{
		Items:      items,
		NextCursor: next,
		AsOf:       domain.IndexedAt{Block: block, Time: time.Date(2026, 8, 22, 9, 30, 0, 0, time.UTC)},
	}
}

func decodeLiquidationList(t *testing.T, body []byte) dto.LiquidationListResponse {
	t.Helper()

	var response dto.LiquidationListResponse
	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("could not decode the body: %v", err)
	}

	return response
}

func TestHistoryPathIsNotMistakenForAReceiptID(t *testing.T) {
	stub := &stubLiquidationService{page: liquidationServicePage(1, cursor.Key{}, nil)}
	engine, _ := newLiquidationRouter(t, stub)

	recorder := performGet(engine, "/liquidations/history")

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected the static route to win, got %d with body %s", recorder.Code, recorder.Body)
	}

	if stub.listCount != 1 {
		t.Fatalf("expected the list handler to run, got %d list calls", stub.listCount)
	}

	if stub.byIDCount != 0 {
		t.Fatal("history must not be treated as a liquidation id")
	}
}

func TestListHistoryReturnsAMaskedPage(t *testing.T) {
	stub := &stubLiquidationService{page: liquidationServicePage(2, cursor.Key{}, nil)}
	engine, masker := newLiquidationRouter(t, stub)

	response := decodeLiquidationList(t, performGet(engine, "/liquidations/history").Body.Bytes())

	if len(response.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(response.Items))
	}

	first, err := masker.Mask(idmask.KindLiquidation, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response.Items[0].ID != first {
		t.Fatalf("expected the masked id %q, got %q", first, response.Items[0].ID)
	}
}

func TestListHistoryReturnsAnEmptyArrayNotNull(t *testing.T) {
	stub := &stubLiquidationService{page: liquidationServicePage(0, cursor.Key{}, nil)}
	engine, _ := newLiquidationRouter(t, stub)

	recorder := performGet(engine, "/liquidations/history")

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200 when nothing has been liquidated, got %d", recorder.Code)
	}

	if !contains(recorder.Body.String(), `"items":[]`) {
		t.Fatalf("expected an empty array, got %s", recorder.Body)
	}
}

func TestListHistoryForwardsCursorAndLimit(t *testing.T) {
	stub := &stubLiquidationService{page: liquidationServicePage(0, cursor.Key{}, nil)}
	engine, _ := newLiquidationRouter(t, stub)

	key := cursor.Encode(cursor.Key{Time: time.Date(2026, 8, 22, 9, 30, 0, 0, time.UTC), ID: 42})

	recorder := performGet(engine, "/liquidations/history?limit=50&cursor="+key)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d with body %s", recorder.Code, recorder.Body)
	}

	if stub.lastRequest.Limit != 50 || stub.lastRequest.After.ID != 42 {
		t.Fatalf("unexpected request %+v", stub.lastRequest)
	}
}

func TestListHistoryUnmasksTheMarketFilter(t *testing.T) {
	stub := &stubLiquidationService{page: liquidationServicePage(0, cursor.Key{}, nil)}
	engine, masker := newLiquidationRouter(t, stub)

	marketID := masker.MustMask(idmask.KindMarket, 7)

	recorder := performGet(engine, "/liquidations/history?market="+marketID)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d with body %s", recorder.Code, recorder.Body)
	}

	if stub.lastRequest.MarketID == nil || *stub.lastRequest.MarketID != 7 {
		t.Fatalf("expected the market to be unmasked to 7, got %v", stub.lastRequest.MarketID)
	}
}

func TestListHistoryWithoutAMarketFilterQueriesEverything(t *testing.T) {
	stub := &stubLiquidationService{page: liquidationServicePage(0, cursor.Key{}, nil)}
	engine, _ := newLiquidationRouter(t, stub)

	performGet(engine, "/liquidations/history")

	if stub.lastRequest.MarketID != nil {
		t.Fatalf("expected no market filter, got %v", stub.lastRequest.MarketID)
	}
}

func TestListHistoryRejectsABadMarketFilter(t *testing.T) {
	masker := newMasker(t)

	cases := []string{"7", "nonsense", masker.MustMask(idmask.KindLiquidation, 7)}

	for _, raw := range cases {
		t.Run(raw, func(t *testing.T) {
			stub := &stubLiquidationService{page: liquidationServicePage(0, cursor.Key{}, nil)}
			engine, _ := newLiquidationRouter(t, stub)

			recorder := performGet(engine, "/liquidations/history?market="+raw)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("expected 400 for market %q, got %d", raw, recorder.Code)
			}

			if stub.listCount != 0 {
				t.Fatal("expected a bad market id to be refused before reaching the service")
			}
		})
	}
}

func TestListHistoryNamesABadParameter(t *testing.T) {
	for param, query := range map[string]string{
		"limit":  "?limit=many",
		"cursor": "?cursor=not-a-cursor",
	} {
		t.Run(param, func(t *testing.T) {
			stub := &stubLiquidationService{page: liquidationServicePage(0, cursor.Key{}, nil)}
			engine, _ := newLiquidationRouter(t, stub)

			recorder := performGet(engine, "/liquidations/history"+query)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("expected 400, got %d", recorder.Code)
			}

			if !contains(decodeError(t, recorder.Body.Bytes()).Error.Message, param) {
				t.Fatalf("expected the message to name %q, got %s", param, recorder.Body)
			}
		})
	}
}

func TestListHistoryReportsAsOf(t *testing.T) {
	block := int64(4_218)
	stub := &stubLiquidationService{page: liquidationServicePage(1, cursor.Key{}, &block)}
	engine, _ := newLiquidationRouter(t, stub)

	response := decodeLiquidationList(t, performGet(engine, "/liquidations/history").Body.Bytes())

	if response.AsOf.Block == nil || *response.AsOf.Block != 4_218 {
		t.Fatalf("expected block 4218, got %v", response.AsOf.Block)
	}
}

func TestListHistoryMapsFailures(t *testing.T) {
	cases := map[string]struct {
		failWith error
		status   int
		code     string
	}{
		"invalid input": {failWith: domain.ErrInvalidInput, status: http.StatusBadRequest, code: dto.CodeBadRequest},
		"unexpected":    {failWith: errors.New("database is down"), status: http.StatusInternalServerError, code: dto.CodeInternalError},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			stub := &stubLiquidationService{listFailWith: testCase.failWith}
			engine, _ := newLiquidationRouter(t, stub)

			recorder := performGet(engine, "/liquidations/history")

			if recorder.Code != testCase.status {
				t.Fatalf("expected %d, got %d", testCase.status, recorder.Code)
			}

			body := decodeError(t, recorder.Body.Bytes())

			if body.Error.Code != testCase.code {
				t.Fatalf("expected %s, got %s", testCase.code, recorder.Body)
			}

			if contains(body.Error.Message, "database") {
				t.Fatalf("expected the internal detail to stay hidden, got %q", body.Error.Message)
			}
		})
	}
}

func TestListHistoryFailsWhenAStoredIDCannotBeMasked(t *testing.T) {
	page := liquidationServicePage(2, cursor.Key{}, nil)
	page.Items[1].ID = 0

	stub := &stubLiquidationService{page: page}
	engine, _ := newLiquidationRouter(t, stub)

	recorder := performGet(engine, "/liquidations/history")

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500 rather than an item with a blank id, got %d", recorder.Code)
	}
}

func TestGetReceiptReturnsTheFullRecord(t *testing.T) {
	stub := &stubLiquidationService{receipt: sampleReceipt(42)}
	engine, masker := newLiquidationRouter(t, stub)

	publicID := masker.MustMask(idmask.KindLiquidation, 42)

	recorder := performGet(engine, "/liquidations/"+publicID)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d with body %s", recorder.Code, recorder.Body)
	}

	var response dto.LiquidationResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("could not decode the body: %v", err)
	}

	if response.ID != publicID {
		t.Fatalf("expected the masked id %q, got %q", publicID, response.ID)
	}

	if response.DebtRepaid.Symbol != "USDC" || response.CollateralSeized.Symbol != "WETH" {
		t.Fatalf("expected each amount in its own asset, got %+v", response)
	}

	if response.TriggerPrice.Decimals != 8 {
		t.Fatalf("expected the price scale to be reported, got %d", response.TriggerPrice.Decimals)
	}

	if stub.lastID != 42 {
		t.Fatalf("expected the handler to unmask to 42, got %d", stub.lastID)
	}
}

func TestGetReceiptRejectsAMalformedIdentifier(t *testing.T) {
	masker := newMasker(t)

	cases := []string{"42", "liq_", "not-a-token", masker.MustMask(idmask.KindTransaction, 42)}

	for _, raw := range cases {
		t.Run(raw, func(t *testing.T) {
			stub := &stubLiquidationService{receipt: sampleReceipt(42)}
			engine, _ := newLiquidationRouter(t, stub)

			recorder := performGet(engine, "/liquidations/"+raw)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("expected 400 for %q, got %d", raw, recorder.Code)
			}

			if stub.byIDCount != 0 {
				t.Fatal("expected a malformed id to be refused before reaching the service")
			}
		})
	}
}

func TestGetReceiptMapsNotFound(t *testing.T) {
	stub := &stubLiquidationService{byIDFailWith: domain.ErrNotFound}
	engine, masker := newLiquidationRouter(t, stub)

	recorder := performGet(engine, "/liquidations/"+masker.MustMask(idmask.KindLiquidation, 42))

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", recorder.Code)
	}

	if decodeError(t, recorder.Body.Bytes()).Error.Code != dto.CodeNotFound {
		t.Fatalf("expected NOT_FOUND, got %s", recorder.Body)
	}
}

func TestGetReceiptMapsUnexpectedFailures(t *testing.T) {
	stub := &stubLiquidationService{byIDFailWith: errors.New("database is down")}
	engine, masker := newLiquidationRouter(t, stub)

	recorder := performGet(engine, "/liquidations/"+masker.MustMask(idmask.KindLiquidation, 42))

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", recorder.Code)
	}

	if contains(decodeError(t, recorder.Body.Bytes()).Error.Message, "database") {
		t.Fatalf("expected the internal detail to stay hidden, got %s", recorder.Body)
	}
}

func TestGetReceiptFailsWhenTheStoredIDCannotBeMasked(t *testing.T) {
	stub := &stubLiquidationService{receipt: sampleReceipt(0)}
	engine, masker := newLiquidationRouter(t, stub)

	recorder := performGet(engine, "/liquidations/"+masker.MustMask(idmask.KindLiquidation, 42))

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500 for an unmaskable id, got %d", recorder.Code)
	}
}

func TestLiquidationResponsesAreJSON(t *testing.T) {
	stub := &stubLiquidationService{page: liquidationServicePage(1, cursor.Key{}, nil), receipt: sampleReceipt(42)}
	engine, masker := newLiquidationRouter(t, stub)

	for _, path := range []string{
		"/liquidations/history",
		"/liquidations/" + masker.MustMask(idmask.KindLiquidation, 42),
	} {
		recorder := performGet(engine, path)

		if contentType := recorder.Header().Get("Content-Type"); !contains(contentType, "application/json") {
			t.Fatalf("expected json for %s, got %q", path, contentType)
		}
	}
}
