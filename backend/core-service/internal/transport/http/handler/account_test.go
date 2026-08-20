package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/handler"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
	"github.com/farhapartex/lending-platform/core-service/pkg/idmask"
)

const testAddress = "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"

type stubTransactionService struct {
	transaction domain.UserTransaction
	failWith    error
	lastAddress string
	lastID      int64
	callCount   int
}

func (s *stubTransactionService) ByID(
	_ context.Context,
	address string,
	id int64,
) (domain.UserTransaction, error) {
	s.lastAddress = address
	s.lastID = id
	s.callCount++

	if s.failWith != nil {
		return domain.UserTransaction{}, s.failWith
	}

	return s.transaction, nil
}

func newMasker(t *testing.T) *idmask.Masker {
	t.Helper()

	masker, err := idmask.New("handler-test-secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	return masker
}

func newRouter(t *testing.T, transactions domain.TransactionService) (*gin.Engine, *idmask.Masker) {
	t.Helper()

	gin.SetMode(gin.TestMode)

	masker := newMasker(t)

	accounts := handler.NewAccountHandler(handler.AccountHandlerParams{
		Transactions: transactions,
		Masker:       masker,
	})

	engine := gin.New()
	engine.GET("/accounts/:address/transactions/:transactionId", accounts.GetTransaction)

	return engine, masker
}

func sampleTransaction() domain.UserTransaction {
	health := int32(12661)

	return domain.UserTransaction{
		ID:                   77,
		UserID:               9,
		Kind:                 domain.TransactionKindBorrow,
		Amount:               bigmath.MustFromString("5100000000"),
		HealthFactorAfterBps: &health,
		BlockNumber:          42,
		BlockTime:            time.Date(2026, 8, 20, 3, 20, 11, 0, time.UTC),
		TxHash:               "0x9f2ccafe",
		LogIndex:             3,
		Asset:                &domain.Asset{Symbol: "USDC", Decimals: 6},
	}
}

func performGet(engine *gin.Engine, path string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, path, nil)

	engine.ServeHTTP(recorder, request)

	return recorder
}

func decodeError(t *testing.T, body []byte) dto.ErrorResponse {
	t.Helper()

	var response dto.ErrorResponse
	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("could not decode the error body: %v", err)
	}

	return response
}

func TestGetTransactionReturnsTheMaskedTransaction(t *testing.T) {
	stub := &stubTransactionService{transaction: sampleTransaction()}
	engine, masker := newRouter(t, stub)

	publicID := masker.MustMask(idmask.KindTransaction, 77)

	recorder := performGet(engine, "/accounts/"+testAddress+"/transactions/"+publicID)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d with body %s", recorder.Code, recorder.Body)
	}

	var response dto.TransactionResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("could not decode the body: %v", err)
	}

	if response.ID != publicID {
		t.Fatalf("expected the masked id %q, got %q", publicID, response.ID)
	}

	if response.Kind != "borrow" || response.Amount.Amount != "5100000000" {
		t.Fatalf("unexpected payload %+v", response)
	}

	if stub.lastID != 77 {
		t.Fatalf("expected the handler to unmask to 77, got %d", stub.lastID)
	}

	if stub.lastAddress != testAddress {
		t.Fatalf("expected the address to be passed through, got %q", stub.lastAddress)
	}
}

func TestGetTransactionRejectsAMalformedIdentifier(t *testing.T) {
	stub := &stubTransactionService{transaction: sampleTransaction()}
	engine, _ := newRouter(t, stub)

	cases := []string{"77", "txn_", "not-a-token", "mkt_zuxjejolmpb2ynh4g32q", "txn_!!!!"}

	for _, identifier := range cases {
		t.Run(identifier, func(t *testing.T) {
			recorder := performGet(engine, "/accounts/"+testAddress+"/transactions/"+identifier)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("expected 400, got %d with body %s", recorder.Code, recorder.Body)
			}

			if decodeError(t, recorder.Body.Bytes()).Error.Code != dto.CodeBadRequest {
				t.Fatalf("expected BAD_REQUEST, got %s", recorder.Body)
			}
		})
	}

	if stub.callCount != 0 {
		t.Fatal("expected a malformed identifier to be rejected before reaching the service")
	}
}

func TestGetTransactionRejectsAMarketIdentifier(t *testing.T) {
	stub := &stubTransactionService{transaction: sampleTransaction()}
	engine, masker := newRouter(t, stub)

	marketID := masker.MustMask(idmask.KindMarket, 77)

	recorder := performGet(engine, "/accounts/"+testAddress+"/transactions/"+marketID)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected a market id to be refused with 400, got %d", recorder.Code)
	}
}

func TestGetTransactionMapsNotFound(t *testing.T) {
	stub := &stubTransactionService{failWith: domain.ErrNotFound}
	engine, masker := newRouter(t, stub)

	recorder := performGet(
		engine,
		"/accounts/"+testAddress+"/transactions/"+masker.MustMask(idmask.KindTransaction, 77),
	)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", recorder.Code)
	}

	if decodeError(t, recorder.Body.Bytes()).Error.Code != dto.CodeNotFound {
		t.Fatalf("expected NOT_FOUND, got %s", recorder.Body)
	}
}

func TestGetTransactionMapsInvalidInput(t *testing.T) {
	stub := &stubTransactionService{failWith: domain.ErrInvalidInput}
	engine, masker := newRouter(t, stub)

	recorder := performGet(
		engine,
		"/accounts/not-an-address/transactions/"+masker.MustMask(idmask.KindTransaction, 77),
	)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", recorder.Code)
	}

	if decodeError(t, recorder.Body.Bytes()).Error.Code != dto.CodeBadRequest {
		t.Fatalf("expected BAD_REQUEST, got %s", recorder.Body)
	}
}

func TestGetTransactionMapsUnexpectedFailures(t *testing.T) {
	stub := &stubTransactionService{failWith: errors.New("database is down")}
	engine, masker := newRouter(t, stub)

	recorder := performGet(
		engine,
		"/accounts/"+testAddress+"/transactions/"+masker.MustMask(idmask.KindTransaction, 77),
	)

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

func TestGetTransactionFailsWhenTheStoredIDCannotBeMasked(t *testing.T) {
	broken := sampleTransaction()
	broken.ID = 0

	stub := &stubTransactionService{transaction: broken}
	engine, masker := newRouter(t, stub)

	recorder := performGet(
		engine,
		"/accounts/"+testAddress+"/transactions/"+masker.MustMask(idmask.KindTransaction, 77),
	)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500 for an unmaskable id, got %d", recorder.Code)
	}
}

func TestGetTransactionResponseIsJSON(t *testing.T) {
	stub := &stubTransactionService{transaction: sampleTransaction()}
	engine, masker := newRouter(t, stub)

	recorder := performGet(
		engine,
		"/accounts/"+testAddress+"/transactions/"+masker.MustMask(idmask.KindTransaction, 77),
	)

	if contentType := recorder.Header().Get("Content-Type"); !contains(contentType, "application/json") {
		t.Fatalf("expected a json content type, got %q", contentType)
	}
}

func contains(haystack, needle string) bool {
	for index := 0; index+len(needle) <= len(haystack); index++ {
		if haystack[index:index+len(needle)] == needle {
			return true
		}
	}

	return false
}
