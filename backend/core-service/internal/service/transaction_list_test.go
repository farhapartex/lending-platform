package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/service"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
)

var listClock = time.Date(2026, 8, 22, 9, 30, 0, 0, time.UTC)

func newListService(
	users *stubUsers,
	transactions *stubTransactions,
	checkpoints domain.CheckpointRepository,
) domain.TransactionService {
	return service.NewTransactionService(service.TransactionServiceParams{
		Users:        users,
		Transactions: transactions,
		Checkpoints:  checkpoints,
		Now:          func() time.Time { return listClock },
	})
}

func rows(count int) []domain.UserTransaction {
	items := make([]domain.UserTransaction, 0, count)

	for index := 0; index < count; index++ {
		items = append(items, domain.UserTransaction{
			ID:          int64(1_000 - index),
			UserID:      9,
			Kind:        domain.TransactionKindBorrow,
			Amount:      bigmath.FromInt64(int64(index+1) * 1_000),
			BlockNumber: int64(500 - index),
			BlockTime:   listClock.Add(-time.Duration(index) * time.Minute),
		})
	}

	return items
}

func listFixtures(count int) (*stubUsers, *stubTransactions) {
	users, transactions := fixtures()
	transactions.page = rows(count)

	return users, transactions
}

func TestListReturnsTheRequestedPage(t *testing.T) {
	users, transactions := listFixtures(3)

	page, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
		Address: knownAddress,
		Limit:   10,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(page.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(page.Items))
	}

	if transactions.lastQuery.UserID != 9 {
		t.Fatalf("expected the resolved user id, got %d", transactions.lastQuery.UserID)
	}
}

func TestListAsksForOneMoreRowThanThePageSize(t *testing.T) {
	users, transactions := listFixtures(3)

	if _, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
		Address: knownAddress,
		Limit:   10,
	}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if transactions.lastQuery.Limit != 11 {
		t.Fatalf("expected a probe row so paging knows when to stop, got limit %d", transactions.lastQuery.Limit)
	}
}

func TestListHidesTheProbeRowAndReturnsACursor(t *testing.T) {
	users, transactions := listFixtures(6)

	page, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
		Address: knownAddress,
		Limit:   5,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(page.Items) != 5 {
		t.Fatalf("expected the probe row to be trimmed, got %d items", len(page.Items))
	}

	if page.NextCursor.IsZero() {
		t.Fatal("expected a next cursor when more rows exist")
	}

	last := page.Items[len(page.Items)-1]

	if page.NextCursor.ID != last.ID || !page.NextCursor.Time.Equal(last.BlockTime) {
		t.Fatalf("expected the cursor to point at the last returned row, got %+v", page.NextCursor)
	}
}

func TestListReturnsNoCursorOnTheLastPage(t *testing.T) {
	for _, count := range []int{0, 1, 4, 5} {
		users, transactions := listFixtures(count)

		page, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
			Address: knownAddress,
			Limit:   5,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !page.NextCursor.IsZero() {
			t.Fatalf("expected no cursor with %d rows available, got %+v", count, page.NextCursor)
		}
	}
}

func TestListDoesNotOfferACursorForAnExactlyFullFinalPage(t *testing.T) {
	users, transactions := listFixtures(5)

	page, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
		Address: knownAddress,
		Limit:   5,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(page.Items) != 5 {
		t.Fatalf("expected a full page, got %d", len(page.Items))
	}

	if !page.NextCursor.IsZero() {
		t.Fatal("a full final page must not offer a cursor that leads to an empty page")
	}
}

func TestListAppliesTheDefaultPageSize(t *testing.T) {
	for _, limit := range []int{0, -1, -100} {
		users, transactions := listFixtures(1)

		if _, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
			Address: knownAddress,
			Limit:   limit,
		}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := service.DefaultTransactionPageSize + 1

		if transactions.lastQuery.Limit != want {
			t.Fatalf("expected limit %d for a requested %d, got %d", want, limit, transactions.lastQuery.Limit)
		}
	}
}

func TestListCapsThePageSize(t *testing.T) {
	users, transactions := listFixtures(1)

	if _, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
		Address: knownAddress,
		Limit:   5_000,
	}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := service.MaxTransactionPageSize + 1

	if transactions.lastQuery.Limit != want {
		t.Fatalf("expected the page size to be capped at %d, got %d", service.MaxTransactionPageSize, transactions.lastQuery.Limit)
	}
}

func TestListPassesFiltersThrough(t *testing.T) {
	users, transactions := listFixtures(1)

	from := listClock.Add(-24 * time.Hour)
	to := listClock
	after := cursor.Key{Time: listClock.Add(-time.Hour), ID: 500}

	if _, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
		Address: knownAddress,
		Kinds:   []domain.TransactionKind{domain.TransactionKindBorrow, domain.TransactionKindRepay},
		From:    &from,
		To:      &to,
		After:   after,
		Limit:   10,
	}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	query := transactions.lastQuery

	if len(query.Kinds) != 2 {
		t.Fatalf("expected both kinds to reach the repository, got %v", query.Kinds)
	}

	if query.From == nil || !query.From.Equal(from) || query.To == nil || !query.To.Equal(to) {
		t.Fatalf("expected the time window to reach the repository, got %v to %v", query.From, query.To)
	}

	if query.After.ID != after.ID || !query.After.Time.Equal(after.Time) {
		t.Fatalf("expected the cursor to reach the repository, got %+v", query.After)
	}
}

func TestListNormalisesTheAddress(t *testing.T) {
	users, transactions := listFixtures(1)

	mixedCase := "  0xF39Fd6e51aad88F6F4ce6aB8827279cffFb92266 "

	if _, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
		Address: mixedCase,
		Limit:   10,
	}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if users.lastRequest != knownAddress {
		t.Fatalf("expected a normalised address, got %q", users.lastRequest)
	}
}

func TestListRejectsABadAddress(t *testing.T) {
	for _, address := range []string{"", "   ", "nonsense", "0x1234"} {
		users, transactions := listFixtures(1)

		_, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
			Address: address,
		})

		if !errors.Is(err, domain.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput for %q, got %v", address, err)
		}

		if transactions.listCount != 0 {
			t.Fatal("expected a bad address to be refused before any lookup")
		}
	}
}

func TestListRejectsAnUnknownKind(t *testing.T) {
	users, transactions := listFixtures(1)

	_, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
		Address: knownAddress,
		Kinds:   []domain.TransactionKind{domain.TransactionKindBorrow, "teleport"},
	})

	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	if transactions.listCount != 0 {
		t.Fatal("expected an unknown kind to be refused before any lookup")
	}
}

func TestListAcceptsEveryKnownKind(t *testing.T) {
	users, transactions := listFixtures(1)

	_, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
		Address: knownAddress,
		Kinds:   domain.AllTransactionKinds,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestListRejectsABackwardsTimeWindow(t *testing.T) {
	users, transactions := listFixtures(1)

	from := listClock
	to := listClock.Add(-time.Hour)

	_, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
		Address: knownAddress,
		From:    &from,
		To:      &to,
	})

	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestListAcceptsAWindowWithOneOpenEnd(t *testing.T) {
	moment := listClock

	cases := map[string]domain.TransactionListRequest{
		"only from": {Address: knownAddress, From: &moment},
		"only to":   {Address: knownAddress, To: &moment},
		"neither":   {Address: knownAddress},
		"equal":     {Address: knownAddress, From: &moment, To: &moment},
	}

	for name, request := range cases {
		t.Run(name, func(t *testing.T) {
			users, transactions := listFixtures(1)

			if _, err := newListService(users, transactions, nil).List(context.Background(), request); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestListReturnsAnEmptyPageForAnUnknownAddress(t *testing.T) {
	users, transactions := listFixtures(3)

	page, err := newListService(users, transactions, nil).List(context.Background(), domain.TransactionListRequest{
		Address: unknownAddress,
		Limit:   10,
	})
	if err != nil {
		t.Fatalf("every address is valid on a blockchain, so this must not be an error: %v", err)
	}

	if page.Items == nil {
		t.Fatal("expected an empty slice rather than nil so it serialises as []")
	}

	if len(page.Items) != 0 {
		t.Fatalf("expected no items, got %d", len(page.Items))
	}

	if !page.NextCursor.IsZero() {
		t.Fatal("expected no cursor for an unknown address")
	}

	if transactions.listCount != 0 {
		t.Fatal("expected no transaction lookup for an unknown address")
	}
}

func TestListStillReportsAsOfForAnUnknownAddress(t *testing.T) {
	users, transactions := listFixtures(0)
	checkpoints := &stubCheckpoints{checkpoint: domain.IndexerCheckpoint{
		LastProcessedBlock: 4_218,
		UpdatedAt:          listClock.Add(-time.Minute),
	}}

	page, err := newListService(users, transactions, checkpoints).
		List(context.Background(), domain.TransactionListRequest{Address: unknownAddress})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if page.AsOf.Block == nil || *page.AsOf.Block != 4_218 {
		t.Fatalf("expected the indexed block to be reported even with no rows, got %v", page.AsOf.Block)
	}
}

func TestListReportsTheIndexedBlock(t *testing.T) {
	users, transactions := listFixtures(2)
	updated := listClock.Add(-90 * time.Second)
	checkpoints := &stubCheckpoints{checkpoint: domain.IndexerCheckpoint{
		LastProcessedBlock: 4_218,
		UpdatedAt:          updated,
	}}

	page, err := newListService(users, transactions, checkpoints).
		List(context.Background(), domain.TransactionListRequest{Address: knownAddress})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if page.AsOf.Block == nil || *page.AsOf.Block != 4_218 {
		t.Fatalf("expected block 4218, got %v", page.AsOf.Block)
	}

	if !page.AsOf.Time.Equal(updated) {
		t.Fatalf("expected the checkpoint time so a stalled indexer is visible, got %s", page.AsOf.Time)
	}

	if checkpoints.lastStream != domain.IndexerStreamProtocolEvents {
		t.Fatalf("expected the events stream, got %q", checkpoints.lastStream)
	}
}

func TestListFallsBackToTheClockWhenNothingHasBeenIndexed(t *testing.T) {
	users, transactions := listFixtures(1)
	checkpoints := &stubCheckpoints{failWith: domain.ErrNotFound}

	page, err := newListService(users, transactions, checkpoints).
		List(context.Background(), domain.TransactionListRequest{Address: knownAddress})
	if err != nil {
		t.Fatalf("a missing checkpoint must not fail the request: %v", err)
	}

	if page.AsOf.Block != nil {
		t.Fatalf("expected no block rather than a made up one, got %v", *page.AsOf.Block)
	}

	if !page.AsOf.Time.Equal(listClock) {
		t.Fatalf("expected the current time, got %s", page.AsOf.Time)
	}
}

func TestListWorksWithoutACheckpointRepository(t *testing.T) {
	users, transactions := listFixtures(1)

	page, err := newListService(users, transactions, nil).
		List(context.Background(), domain.TransactionListRequest{Address: knownAddress})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if page.AsOf.Block != nil {
		t.Fatalf("expected no block, got %v", *page.AsOf.Block)
	}

	if !page.AsOf.Time.Equal(listClock) {
		t.Fatalf("expected the current time, got %s", page.AsOf.Time)
	}
}

func TestListPropagatesACheckpointFailure(t *testing.T) {
	users, transactions := listFixtures(1)
	checkpoints := &stubCheckpoints{failWith: errors.New("database is down")}

	_, err := newListService(users, transactions, checkpoints).
		List(context.Background(), domain.TransactionListRequest{Address: knownAddress})

	if err == nil {
		t.Fatal("expected the failure to surface")
	}

	if errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected the raw failure rather than not found, got %v", err)
	}
}

func TestListPropagatesAUserLookupFailure(t *testing.T) {
	users, transactions := listFixtures(1)
	users.failWith = errors.New("database is down")

	_, err := newListService(users, transactions, nil).
		List(context.Background(), domain.TransactionListRequest{Address: knownAddress})

	if err == nil {
		t.Fatal("expected the failure to surface")
	}

	if errors.Is(err, domain.ErrNotFound) || errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected the raw failure, got %v", err)
	}
}

func TestListPropagatesARepositoryFailure(t *testing.T) {
	users, transactions := listFixtures(1)
	transactions.listFailWith = errors.New("database is down")

	_, err := newListService(users, transactions, nil).
		List(context.Background(), domain.TransactionListRequest{Address: knownAddress})

	if err == nil {
		t.Fatal("expected the failure to surface")
	}
}

func TestListDefaultsTheClockWhenNoneIsSupplied(t *testing.T) {
	users, transactions := listFixtures(1)

	page, err := service.NewTransactionService(service.TransactionServiceParams{
		Users:        users,
		Transactions: transactions,
	}).List(context.Background(), domain.TransactionListRequest{Address: knownAddress})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if page.AsOf.Time.IsZero() {
		t.Fatal("expected a real timestamp from the default clock")
	}
}
