package handler_test

import (
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

var listMoment = time.Date(2026, 8, 22, 9, 30, 0, 0, time.UTC)

func newListRouter(t *testing.T, transactions domain.TransactionService) (*gin.Engine, *idmask.Masker) {
	t.Helper()

	gin.SetMode(gin.TestMode)

	masker := newMasker(t)

	accounts := handler.NewAccountHandler(handler.AccountHandlerParams{
		Transactions: transactions,
		Masker:       masker,
	})

	engine := gin.New()
	engine.GET("/accounts/:address/transactions", accounts.ListTransactions)

	return engine, masker
}

func samplePage(count int, next cursor.Key, block *int64) domain.TransactionPage {
	items := make([]domain.UserTransaction, 0, count)

	for index := 0; index < count; index++ {
		items = append(items, domain.UserTransaction{
			ID:          int64(index + 1),
			Kind:        domain.TransactionKindBorrow,
			Amount:      bigmath.FromInt64(1_000),
			BlockNumber: int64(500 + index),
			BlockTime:   listMoment,
			TxHash:      "0xabc",
			Asset:       &domain.Asset{Symbol: "USDC", Decimals: 6},
		})
	}

	return domain.TransactionPage{
		Items:      items,
		NextCursor: next,
		AsOf:       domain.IndexedAt{Block: block, Time: listMoment},
	}
}

func decodeList(t *testing.T, body []byte) dto.TransactionListResponse {
	t.Helper()

	var response dto.TransactionListResponse
	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("could not decode the body: %v", err)
	}

	return response
}

func TestListTransactionsReturnsAMaskedPage(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(2, cursor.Key{}, nil)}
	engine, masker := newListRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/transactions")

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d with body %s", recorder.Code, recorder.Body)
	}

	response := decodeList(t, recorder.Body.Bytes())

	if len(response.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(response.Items))
	}

	first, err := masker.Mask(idmask.KindTransaction, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response.Items[0].ID != first {
		t.Fatalf("expected the masked id %q, got %q", first, response.Items[0].ID)
	}
}

func TestListTransactionsPassesTheAddressThrough(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(0, cursor.Key{}, nil)}
	engine, _ := newListRouter(t, stub)

	performGet(engine, "/accounts/"+testAddress+"/transactions")

	if stub.lastRequest.Address != testAddress {
		t.Fatalf("expected the address to reach the service, got %q", stub.lastRequest.Address)
	}
}

func TestListTransactionsForwardsEveryFilter(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(0, cursor.Key{}, nil)}
	engine, _ := newListRouter(t, stub)

	key := cursor.Encode(cursor.Key{Time: listMoment, ID: 42})

	path := "/accounts/" + testAddress + "/transactions" +
		"?kind=borrow,repay&from=2026-08-01T00:00:00Z&to=2026-08-22T00:00:00Z&limit=50&cursor=" + key

	if recorder := performGet(engine, path); recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d with body %s", recorder.Code, recorder.Body)
	}

	request := stub.lastRequest

	if len(request.Kinds) != 2 {
		t.Fatalf("expected two kinds, got %v", request.Kinds)
	}

	if request.From == nil || request.To == nil {
		t.Fatalf("expected both bounds, got %v and %v", request.From, request.To)
	}

	if request.Limit != 50 {
		t.Fatalf("expected limit 50, got %d", request.Limit)
	}

	if request.After.ID != 42 {
		t.Fatalf("expected the cursor to be decoded, got %+v", request.After)
	}
}

func TestListTransactionsReturnsAnEmptyArrayNotNull(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(0, cursor.Key{}, nil)}
	engine, _ := newListRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/transactions")

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200 for a wallet with no history, got %d", recorder.Code)
	}

	if !contains(recorder.Body.String(), `"items":[]`) {
		t.Fatalf("expected an empty array so the client can render without a null check, got %s", recorder.Body)
	}
}

func TestListTransactionsReportsTheNextCursor(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(2, cursor.Key{Time: listMoment, ID: 7}, nil)}
	engine, _ := newListRouter(t, stub)

	response := decodeList(t, performGet(engine, "/accounts/"+testAddress+"/transactions").Body.Bytes())

	if response.NextCursor == nil {
		t.Fatal("expected a cursor")
	}

	decoded, err := cursor.Decode(*response.NextCursor)
	if err != nil {
		t.Fatalf("expected a usable cursor, got %v", err)
	}

	if decoded.ID != 7 {
		t.Fatalf("expected the cursor to point at row 7, got %d", decoded.ID)
	}
}

func TestListTransactionsReportsAsOf(t *testing.T) {
	block := int64(4_218)
	stub := &stubTransactionService{page: samplePage(1, cursor.Key{}, &block)}
	engine, _ := newListRouter(t, stub)

	response := decodeList(t, performGet(engine, "/accounts/"+testAddress+"/transactions").Body.Bytes())

	if response.AsOf.Block == nil || *response.AsOf.Block != 4_218 {
		t.Fatalf("expected block 4218, got %v", response.AsOf.Block)
	}

	if response.AsOf.Time != "2026-08-22T09:30:00Z" {
		t.Fatalf("unexpected as of time %q", response.AsOf.Time)
	}
}

func TestListTransactionsNamesTheParameterThatWasWrong(t *testing.T) {
	cases := map[string]string{
		"kind":   "?kind=teleport",
		"from":   "?from=2026-08-01",
		"to":     "?to=yesterday",
		"limit":  "?limit=many",
		"cursor": "?cursor=not-a-cursor",
	}

	for param, query := range cases {
		t.Run(param, func(t *testing.T) {
			stub := &stubTransactionService{page: samplePage(0, cursor.Key{}, nil)}
			engine, _ := newListRouter(t, stub)

			recorder := performGet(engine, "/accounts/"+testAddress+"/transactions"+query)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("expected 400, got %d with body %s", recorder.Code, recorder.Body)
			}

			body := decodeError(t, recorder.Body.Bytes())

			if body.Error.Code != dto.CodeBadRequest {
				t.Fatalf("expected BAD_REQUEST, got %s", recorder.Body)
			}

			if !contains(body.Error.Message, param) {
				t.Fatalf("expected the message to name %q, got %q", param, body.Error.Message)
			}

			if stub.listCount != 0 {
				t.Fatal("expected a bad parameter to be refused before reaching the service")
			}
		})
	}
}

func TestListTransactionsMapsInvalidInput(t *testing.T) {
	stub := &stubTransactionService{listFailWith: domain.ErrInvalidInput}
	engine, _ := newListRouter(t, stub)

	recorder := performGet(engine, "/accounts/not-an-address/transactions")

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", recorder.Code)
	}

	if decodeError(t, recorder.Body.Bytes()).Error.Code != dto.CodeBadRequest {
		t.Fatalf("expected BAD_REQUEST, got %s", recorder.Body)
	}
}

func TestListTransactionsMapsUnexpectedFailures(t *testing.T) {
	stub := &stubTransactionService{listFailWith: errors.New("database is down")}
	engine, _ := newListRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/transactions")

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", recorder.Code)
	}

	body := decodeError(t, recorder.Body.Bytes())

	if body.Error.Code != dto.CodeInternalError {
		t.Fatalf("expected INTERNAL_ERROR, got %s", recorder.Body)
	}

	if contains(body.Error.Message, "database") {
		t.Fatalf("expected the internal detail to stay hidden, got %q", body.Error.Message)
	}
}

func TestListTransactionsFailsWhenAStoredIDCannotBeMasked(t *testing.T) {
	page := samplePage(2, cursor.Key{}, nil)
	page.Items[1].ID = 0

	stub := &stubTransactionService{page: page}
	engine, _ := newListRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/transactions")

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500 rather than a page with a blank id, got %d", recorder.Code)
	}
}

func TestListTransactionsResponseIsJSON(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(1, cursor.Key{}, nil)}
	engine, _ := newListRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/transactions")

	if contentType := recorder.Header().Get("Content-Type"); !contains(contentType, "application/json") {
		t.Fatalf("expected a json content type, got %q", contentType)
	}
}
