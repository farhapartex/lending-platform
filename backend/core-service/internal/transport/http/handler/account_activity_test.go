package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/handler"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
	"github.com/farhapartex/lending-platform/core-service/pkg/idmask"
)

func newActivityRouter(t *testing.T, transactions domain.TransactionService) (*gin.Engine, *idmask.Masker) {
	t.Helper()

	gin.SetMode(gin.TestMode)

	masker := newMasker(t)

	accounts := handler.NewAccountHandler(handler.AccountHandlerParams{
		Transactions: transactions,
		Masker:       masker,
	})

	engine := gin.New()
	engine.GET("/accounts/:address/activity", accounts.GetActivity)

	return engine, masker
}

func decodeActivity(t *testing.T, body []byte) dto.ActivityResponse {
	t.Helper()

	var response dto.ActivityResponse
	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("could not decode the body: %v", err)
	}

	return response
}

func TestGetActivityReturnsMaskedItems(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(3, cursor.Key{}, nil)}
	engine, masker := newActivityRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/activity")

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d with body %s", recorder.Code, recorder.Body)
	}

	response := decodeActivity(t, recorder.Body.Bytes())

	if len(response.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(response.Items))
	}

	first, err := masker.Mask(idmask.KindTransaction, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response.Items[0].ID != first {
		t.Fatalf("expected the masked id %q, got %q", first, response.Items[0].ID)
	}
}

func TestGetActivityPassesTheAddressAndLimitThrough(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(0, cursor.Key{}, nil)}
	engine, _ := newActivityRouter(t, stub)

	performGet(engine, "/accounts/"+testAddress+"/activity?limit=8")

	if stub.lastAddress != testAddress {
		t.Fatalf("expected the address to reach the service, got %q", stub.lastAddress)
	}

	if stub.lastLimit != 8 {
		t.Fatalf("expected limit 8, got %d", stub.lastLimit)
	}
}

func TestGetActivityLeavesTheDefaultSizeToTheService(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(0, cursor.Key{}, nil)}
	engine, _ := newActivityRouter(t, stub)

	performGet(engine, "/accounts/"+testAddress+"/activity")

	if stub.lastLimit != 0 {
		t.Fatalf("expected the handler not to invent a default, got %d", stub.lastLimit)
	}
}

func TestGetActivityReturnsAnEmptyArrayNotNull(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(0, cursor.Key{}, nil)}
	engine, _ := newActivityRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/activity")

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200 for a wallet with no activity, got %d", recorder.Code)
	}

	if !contains(recorder.Body.String(), `"items":[]`) {
		t.Fatalf("expected an empty array, got %s", recorder.Body)
	}
}

func TestGetActivityEmitsNoCursor(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(2, cursor.Key{Time: listMoment, ID: 7}, nil)}
	engine, _ := newActivityRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/activity")

	if contains(recorder.Body.String(), "next_cursor") {
		t.Fatalf("the dashboard list does not page, got %s", recorder.Body)
	}
}

func TestGetActivityReportsAsOf(t *testing.T) {
	block := int64(4_218)
	stub := &stubTransactionService{page: samplePage(1, cursor.Key{}, &block)}
	engine, _ := newActivityRouter(t, stub)

	response := decodeActivity(t, performGet(engine, "/accounts/"+testAddress+"/activity").Body.Bytes())

	if response.AsOf.Block == nil || *response.AsOf.Block != 4_218 {
		t.Fatalf("expected block 4218, got %v", response.AsOf.Block)
	}

	if response.AsOf.Time != "2026-08-22T09:30:00Z" {
		t.Fatalf("unexpected as of time %q", response.AsOf.Time)
	}
}

func TestGetActivityNamesABadLimit(t *testing.T) {
	for _, raw := range []string{"many", "-5", "1.5"} {
		t.Run(raw, func(t *testing.T) {
			stub := &stubTransactionService{page: samplePage(0, cursor.Key{}, nil)}
			engine, _ := newActivityRouter(t, stub)

			recorder := performGet(engine, "/accounts/"+testAddress+"/activity?limit="+raw)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("expected 400, got %d with body %s", recorder.Code, recorder.Body)
			}

			body := decodeError(t, recorder.Body.Bytes())

			if body.Error.Code != dto.CodeBadRequest {
				t.Fatalf("expected BAD_REQUEST, got %s", recorder.Body)
			}

			if !contains(body.Error.Message, "limit") {
				t.Fatalf("expected the message to name the limit parameter, got %q", body.Error.Message)
			}

			if stub.activityCount != 0 {
				t.Fatal("expected a bad limit to be refused before reaching the service")
			}
		})
	}
}

func TestGetActivityIgnoresListOnlyParameters(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(1, cursor.Key{}, nil)}
	engine, _ := newActivityRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/activity?kind=teleport&cursor=nonsense&from=yesterday")

	if recorder.Code != http.StatusOK {
		t.Fatalf("this endpoint takes no filters, so they should be ignored rather than refused, got %d", recorder.Code)
	}
}

func TestGetActivityMapsInvalidInput(t *testing.T) {
	stub := &stubTransactionService{activityFailWith: domain.ErrInvalidInput}
	engine, _ := newActivityRouter(t, stub)

	recorder := performGet(engine, "/accounts/not-an-address/activity")

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", recorder.Code)
	}

	if decodeError(t, recorder.Body.Bytes()).Error.Code != dto.CodeBadRequest {
		t.Fatalf("expected BAD_REQUEST, got %s", recorder.Body)
	}
}

func TestGetActivityMapsUnexpectedFailures(t *testing.T) {
	stub := &stubTransactionService{activityFailWith: errors.New("database is down")}
	engine, _ := newActivityRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/activity")

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

func TestGetActivityFailsWhenAStoredIDCannotBeMasked(t *testing.T) {
	page := samplePage(2, cursor.Key{}, nil)
	page.Items[1].ID = 0

	stub := &stubTransactionService{page: page}
	engine, _ := newActivityRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/activity")

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500 rather than an item with a blank id, got %d", recorder.Code)
	}
}

func TestGetActivityResponseIsJSON(t *testing.T) {
	stub := &stubTransactionService{page: samplePage(1, cursor.Key{}, nil)}
	engine, _ := newActivityRouter(t, stub)

	recorder := performGet(engine, "/accounts/"+testAddress+"/activity")

	if contentType := recorder.Header().Get("Content-Type"); !contains(contentType, "application/json") {
		t.Fatalf("expected a json content type, got %q", contentType)
	}
}
