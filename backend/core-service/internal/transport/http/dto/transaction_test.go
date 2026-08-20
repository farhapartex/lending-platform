package dto_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

func sampleTransaction() domain.UserTransaction {
	health := int32(12661)

	return domain.UserTransaction{
		ID:                   77,
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

func TestNewTransactionResponse(t *testing.T) {
	response := dto.NewTransactionResponse(sampleTransaction(), "txn_mvtep746x22iqzn52kca")

	if response.ID != "txn_mvtep746x22iqzn52kca" {
		t.Fatalf("expected the masked id, got %q", response.ID)
	}

	if response.Kind != "borrow" {
		t.Fatalf("expected borrow, got %q", response.Kind)
	}

	if response.Amount.Amount != "5100000000" || response.Amount.Decimals != 6 || response.Amount.Symbol != "USDC" {
		t.Fatalf("unexpected amount %+v", response.Amount)
	}

	if response.HealthFactorAfterBps == nil || *response.HealthFactorAfterBps != 12661 {
		t.Fatalf("expected a health factor of 12661, got %v", response.HealthFactorAfterBps)
	}

	if response.TxHash != "0x9f2ccafe" || response.Block != 42 || response.LogIndex != 3 {
		t.Fatalf("unexpected chain fields %+v", response)
	}

	if response.BlockTime != "2026-08-20T03:20:11Z" {
		t.Fatalf("expected an RFC 3339 timestamp, got %q", response.BlockTime)
	}

	if response.Status != dto.TransactionStatusConfirmed {
		t.Fatalf("expected confirmed, got %q", response.Status)
	}
}

func TestNewTransactionResponseNeverLeaksTheDatabaseID(t *testing.T) {
	encoded, err := json.Marshal(dto.NewTransactionResponse(sampleTransaction(), "txn_masked"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if decoded["id"] != "txn_masked" {
		t.Fatalf("expected the masked id, got %v", decoded["id"])
	}

	if _, present := decoded["user_id"]; present {
		t.Fatal("expected no user id in the response")
	}

	if _, present := decoded["event_id"]; present {
		t.Fatal("expected no event id in the response")
	}
}

func TestNewTransactionResponseConvertsTimeToUTC(t *testing.T) {
	zone := time.FixedZone("UTC+6", 6*60*60)

	transaction := sampleTransaction()
	transaction.BlockTime = time.Date(2026, 8, 20, 9, 20, 11, 0, zone)

	response := dto.NewTransactionResponse(transaction, "txn_masked")

	if response.BlockTime != "2026-08-20T03:20:11Z" {
		t.Fatalf("expected the timestamp in UTC, got %q", response.BlockTime)
	}
}

func TestNewTransactionResponseHandlesAMissingHealthFactor(t *testing.T) {
	transaction := sampleTransaction()
	transaction.HealthFactorAfterBps = nil

	response := dto.NewTransactionResponse(transaction, "txn_masked")

	if response.HealthFactorAfterBps != nil {
		t.Fatalf("expected nil, got %v", response.HealthFactorAfterBps)
	}

	encoded, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	value, present := decoded["health_factor_after_bps"]
	if !present {
		t.Fatal("expected the field to be present so the client can distinguish absent from zero")
	}

	if value != nil {
		t.Fatalf("expected null, got %v", value)
	}
}

func TestNewTransactionResponseHandlesAMissingAsset(t *testing.T) {
	transaction := sampleTransaction()
	transaction.Asset = nil

	response := dto.NewTransactionResponse(transaction, "txn_masked")

	if response.Amount.Amount != "5100000000" {
		t.Fatalf("expected the raw amount to survive, got %q", response.Amount.Amount)
	}

	if response.Amount.Decimals != 0 || response.Amount.Symbol != "" {
		t.Fatalf("expected empty units rather than a guess, got %+v", response.Amount)
	}
}

func TestNewTransactionResponseCoversEveryKind(t *testing.T) {
	kinds := []domain.TransactionKind{
		domain.TransactionKindDeposit,
		domain.TransactionKindWithdraw,
		domain.TransactionKindBorrow,
		domain.TransactionKindRepay,
		domain.TransactionKindCollateralAdded,
		domain.TransactionKindCollateralWithdrawn,
		domain.TransactionKindLiquidation,
	}

	for _, kind := range kinds {
		t.Run(string(kind), func(t *testing.T) {
			transaction := sampleTransaction()
			transaction.Kind = kind

			if got := dto.NewTransactionResponse(transaction, "txn_masked").Kind; got != string(kind) {
				t.Fatalf("expected %q, got %q", kind, got)
			}
		})
	}
}

func TestTransactionResponseKeepsLargeAmountsExact(t *testing.T) {
	transaction := sampleTransaction()
	transaction.Amount = bigmath.MustFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935")
	transaction.Asset = &domain.Asset{Symbol: "WETH", Decimals: 18}

	encoded, err := json.Marshal(dto.NewTransactionResponse(transaction, "txn_masked"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !contains(string(encoded), `"115792089237316195423570985008687907853269984665640564039457584007913129639935"`) {
		t.Fatalf("expected the full value as a string, got %s", encoded)
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
