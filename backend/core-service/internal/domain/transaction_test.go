package domain_test

import (
	"testing"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

func TestParseTransactionKindAcceptsEveryStoredKind(t *testing.T) {
	for _, kind := range domain.AllTransactionKinds {
		parsed, ok := domain.ParseTransactionKind(string(kind))

		if !ok {
			t.Fatalf("expected %q to be accepted", kind)
		}

		if parsed != kind {
			t.Fatalf("expected %q, got %q", kind, parsed)
		}
	}
}

func TestParseTransactionKindRejectsAnythingElse(t *testing.T) {
	cases := []string{
		"",
		" ",
		"teleport",
		"Deposit",
		"DEPOSIT",
		"collateralAdded",
		"collateral-added",
		"deposit ",
	}

	for _, raw := range cases {
		t.Run(raw, func(t *testing.T) {
			parsed, ok := domain.ParseTransactionKind(raw)

			if ok {
				t.Fatalf("expected %q to be refused, got %q", raw, parsed)
			}

			if parsed != "" {
				t.Fatalf("expected an empty kind on refusal, got %q", parsed)
			}
		})
	}
}

func TestAllTransactionKindsCoversTheDatabaseConstraint(t *testing.T) {
	want := []domain.TransactionKind{
		domain.TransactionKindDeposit,
		domain.TransactionKindWithdraw,
		domain.TransactionKindBorrow,
		domain.TransactionKindRepay,
		domain.TransactionKindCollateralAdded,
		domain.TransactionKindCollateralWithdrawn,
		domain.TransactionKindLiquidation,
	}

	if len(domain.AllTransactionKinds) != len(want) {
		t.Fatalf("expected %d kinds to match the check constraint, got %d", len(want), len(domain.AllTransactionKinds))
	}

	for index, kind := range want {
		if domain.AllTransactionKinds[index] != kind {
			t.Fatalf("expected %v, got %v", want, domain.AllTransactionKinds)
		}
	}
}
