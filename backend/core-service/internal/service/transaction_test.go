package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/service"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

const (
	knownAddress   = "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"
	unknownAddress = "0x70997970c51812dc3a010c7d01b50e0d17dc79c8"
)

type stubUsers struct {
	byAddress    map[string]domain.User
	failWith     error
	lastRequest  string
	requestCount int
}

func (s *stubUsers) ByAddress(_ context.Context, address string) (domain.User, error) {
	s.lastRequest = address
	s.requestCount++

	if s.failWith != nil {
		return domain.User{}, s.failWith
	}

	user, found := s.byAddress[address]
	if !found {
		return domain.User{}, domain.ErrNotFound
	}

	return user, nil
}

func (s *stubUsers) EnsureByAddress(context.Context, int64, string) (domain.User, error) {
	return domain.User{}, errors.New("not expected in these tests")
}

type stubTransactions struct {
	byKey        map[int64]domain.UserTransaction
	failWith     error
	lastUserID   int64
	lastID       int64
	requestCount int
}

func (s *stubTransactions) List(context.Context, domain.TransactionQuery) ([]domain.UserTransaction, error) {
	return nil, errors.New("not expected in these tests")
}

func (s *stubTransactions) ByID(_ context.Context, userID int64, id int64) (domain.UserTransaction, error) {
	s.lastUserID = userID
	s.lastID = id
	s.requestCount++

	if s.failWith != nil {
		return domain.UserTransaction{}, s.failWith
	}

	transaction, found := s.byKey[id]
	if !found || transaction.UserID != userID {
		return domain.UserTransaction{}, domain.ErrNotFound
	}

	return transaction, nil
}

func (s *stubTransactions) Insert(context.Context, *domain.UserTransaction) error {
	return errors.New("not expected in these tests")
}

func newService(users *stubUsers, transactions *stubTransactions) domain.TransactionService {
	return service.NewTransactionService(service.TransactionServiceParams{
		Users:        users,
		Transactions: transactions,
	})
}

func fixtures() (*stubUsers, *stubTransactions) {
	users := &stubUsers{byAddress: map[string]domain.User{
		knownAddress: {ID: 9, Address: knownAddress},
	}}

	transactions := &stubTransactions{byKey: map[int64]domain.UserTransaction{
		77: {ID: 77, UserID: 9, Kind: domain.TransactionKindBorrow, Amount: bigmath.FromInt64(5_100_000_000)},
		88: {ID: 88, UserID: 5, Kind: domain.TransactionKindDeposit, Amount: bigmath.FromInt64(1_000)},
	}}

	return users, transactions
}

func TestByIDReturnsTheTransaction(t *testing.T) {
	users, transactions := fixtures()

	found, err := newService(users, transactions).ByID(context.Background(), knownAddress, 77)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.ID != 77 {
		t.Fatalf("expected transaction 77, got %d", found.ID)
	}

	if transactions.lastUserID != 9 {
		t.Fatalf("expected the resolved user id to be used, got %d", transactions.lastUserID)
	}
}

func TestByIDNormalisesTheAddressBeforeLookup(t *testing.T) {
	users, transactions := fixtures()

	mixedCase := "0xF39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

	if _, err := newService(users, transactions).ByID(context.Background(), "  "+mixedCase+" ", 77); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if users.lastRequest != knownAddress {
		t.Fatalf("expected a normalised address, got %q", users.lastRequest)
	}
}

func TestByIDRejectsABadAddress(t *testing.T) {
	users, transactions := fixtures()

	for _, address := range []string{"", "   ", "nonsense", "0x1234", "f39fd6e51aad88f6f4ce6ab8827279cfffb92266"} {
		_, err := newService(users, transactions).ByID(context.Background(), address, 77)

		if !errors.Is(err, domain.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput for %q, got %v", address, err)
		}
	}

	if users.requestCount != 0 {
		t.Fatal("expected a bad address to be rejected before any lookup")
	}
}

func TestByIDRejectsANonPositiveID(t *testing.T) {
	users, transactions := fixtures()

	for _, id := range []int64{0, -1} {
		_, err := newService(users, transactions).ByID(context.Background(), knownAddress, id)

		if !errors.Is(err, domain.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput for id %d, got %v", id, err)
		}
	}

	if transactions.requestCount != 0 {
		t.Fatal("expected a bad id to be rejected before any lookup")
	}
}

func TestByIDReportsAnUnknownAddressAsNotFound(t *testing.T) {
	users, transactions := fixtures()

	_, err := newService(users, transactions).ByID(context.Background(), unknownAddress, 77)

	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	if transactions.requestCount != 0 {
		t.Fatal("expected no transaction lookup when the address is unknown")
	}
}

func TestByIDDoesNotLeakAnotherUsersTransaction(t *testing.T) {
	users, transactions := fixtures()

	_, err := newService(users, transactions).ByID(context.Background(), knownAddress, 88)

	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected another user's transaction to be not found, got %v", err)
	}
}

func TestByIDReportsAMissingTransaction(t *testing.T) {
	users, transactions := fixtures()

	_, err := newService(users, transactions).ByID(context.Background(), knownAddress, 999)

	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestByIDPropagatesAUserLookupFailure(t *testing.T) {
	users, transactions := fixtures()
	users.failWith = errors.New("database is down")

	_, err := newService(users, transactions).ByID(context.Background(), knownAddress, 77)

	if err == nil {
		t.Fatal("expected the failure to surface")
	}

	if errors.Is(err, domain.ErrNotFound) || errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected the raw failure rather than a domain error, got %v", err)
	}
}

func TestByIDPropagatesATransactionLookupFailure(t *testing.T) {
	users, transactions := fixtures()
	transactions.failWith = errors.New("database is down")

	_, err := newService(users, transactions).ByID(context.Background(), knownAddress, 77)

	if err == nil {
		t.Fatal("expected the failure to surface")
	}

	if errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected the raw failure rather than not found, got %v", err)
	}
}
