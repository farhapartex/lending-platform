package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/service"
)

func TestRecentActivityReturnsTheNewestRows(t *testing.T) {
	users, transactions := listFixtures(3)

	page, err := newListService(users, transactions, nil).
		RecentActivity(context.Background(), knownAddress, 5)
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

func TestRecentActivityNeverPages(t *testing.T) {
	users, transactions := listFixtures(50)

	page, err := newListService(users, transactions, nil).
		RecentActivity(context.Background(), knownAddress, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !page.NextCursor.IsZero() {
		t.Fatal("the dashboard list does not page, so it must not offer a cursor")
	}

	if len(page.Items) != 5 {
		t.Fatalf("expected exactly the requested count, got %d", len(page.Items))
	}
}

func TestRecentActivityAsksForNoProbeRow(t *testing.T) {
	users, transactions := listFixtures(50)

	if _, err := newListService(users, transactions, nil).
		RecentActivity(context.Background(), knownAddress, 5); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if transactions.lastQuery.Limit != 5 {
		t.Fatalf("expected no extra row to be fetched when there is no paging, got limit %d", transactions.lastQuery.Limit)
	}
}

func TestRecentActivityAppliesNoFilters(t *testing.T) {
	users, transactions := listFixtures(3)

	if _, err := newListService(users, transactions, nil).
		RecentActivity(context.Background(), knownAddress, 5); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	query := transactions.lastQuery

	if len(query.Kinds) != 0 {
		t.Fatalf("expected no kind filter, got %v", query.Kinds)
	}

	if query.From != nil || query.To != nil {
		t.Fatalf("expected no time window, got %v to %v", query.From, query.To)
	}

	if !query.After.IsZero() {
		t.Fatalf("expected no cursor, got %+v", query.After)
	}
}

func TestRecentActivityAppliesTheDefaultSize(t *testing.T) {
	for _, limit := range []int{0, -1, -50} {
		users, transactions := listFixtures(50)

		if _, err := newListService(users, transactions, nil).
			RecentActivity(context.Background(), knownAddress, limit); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if transactions.lastQuery.Limit != service.DefaultActivitySize {
			t.Fatalf(
				"expected the default %d for a requested %d, got %d",
				service.DefaultActivitySize, limit, transactions.lastQuery.Limit,
			)
		}
	}
}

func TestRecentActivityCapsTheSize(t *testing.T) {
	users, transactions := listFixtures(50)

	if _, err := newListService(users, transactions, nil).
		RecentActivity(context.Background(), knownAddress, 500); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if transactions.lastQuery.Limit != service.MaxActivitySize {
		t.Fatalf("expected the size to be capped at %d, got %d", service.MaxActivitySize, transactions.lastQuery.Limit)
	}
}

func TestRecentActivityUsesASmallerCapThanTheFullList(t *testing.T) {
	if service.MaxActivitySize >= service.MaxTransactionPageSize {
		t.Fatal("the dashboard list is meant to be short, so its cap must stay below the history cap")
	}

	if service.DefaultActivitySize >= service.DefaultTransactionPageSize {
		t.Fatal("the dashboard default must stay below the history default")
	}
}

func TestRecentActivityNormalisesTheAddress(t *testing.T) {
	users, transactions := listFixtures(1)

	if _, err := newListService(users, transactions, nil).
		RecentActivity(context.Background(), "  0xF39Fd6e51aad88F6F4ce6aB8827279cffFb92266 ", 5); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if users.lastRequest != knownAddress {
		t.Fatalf("expected a normalised address, got %q", users.lastRequest)
	}
}

func TestRecentActivityRejectsABadAddress(t *testing.T) {
	for _, address := range []string{"", "   ", "nonsense", "0x1234"} {
		users, transactions := listFixtures(1)

		_, err := newListService(users, transactions, nil).
			RecentActivity(context.Background(), address, 5)

		if !errors.Is(err, domain.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput for %q, got %v", address, err)
		}

		if transactions.listCount != 0 {
			t.Fatal("expected a bad address to be refused before any lookup")
		}
	}
}

func TestRecentActivityReturnsAnEmptyListForAnUnknownAddress(t *testing.T) {
	users, transactions := listFixtures(3)

	page, err := newListService(users, transactions, nil).
		RecentActivity(context.Background(), unknownAddress, 5)
	if err != nil {
		t.Fatalf("a wallet with no history is not an error: %v", err)
	}

	if page.Items == nil {
		t.Fatal("expected an empty slice rather than nil so it serialises as []")
	}

	if len(page.Items) != 0 {
		t.Fatalf("expected no items, got %d", len(page.Items))
	}

	if transactions.listCount != 0 {
		t.Fatal("expected no transaction lookup for an unknown address")
	}
}

func TestRecentActivityReportsTheIndexedBlock(t *testing.T) {
	users, transactions := listFixtures(2)
	updated := listClock.Add(-2 * time.Minute)
	checkpoints := &stubCheckpoints{checkpoint: domain.IndexerCheckpoint{
		LastProcessedBlock: 4_218,
		UpdatedAt:          updated,
	}}

	page, err := newListService(users, transactions, checkpoints).
		RecentActivity(context.Background(), knownAddress, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if page.AsOf.Block == nil || *page.AsOf.Block != 4_218 {
		t.Fatalf("expected block 4218, got %v", page.AsOf.Block)
	}

	if !page.AsOf.Time.Equal(updated) {
		t.Fatalf("expected the checkpoint time, got %s", page.AsOf.Time)
	}
}

func TestRecentActivityStillReportsAsOfForAnUnknownAddress(t *testing.T) {
	users, transactions := listFixtures(0)
	checkpoints := &stubCheckpoints{checkpoint: domain.IndexerCheckpoint{
		LastProcessedBlock: 4_218,
		UpdatedAt:          listClock,
	}}

	page, err := newListService(users, transactions, checkpoints).
		RecentActivity(context.Background(), unknownAddress, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if page.AsOf.Block == nil || *page.AsOf.Block != 4_218 {
		t.Fatalf("expected the indexed block even with no rows, got %v", page.AsOf.Block)
	}
}

func TestRecentActivityFallsBackToTheClockWhenNothingHasBeenIndexed(t *testing.T) {
	users, transactions := listFixtures(1)
	checkpoints := &stubCheckpoints{failWith: domain.ErrNotFound}

	page, err := newListService(users, transactions, checkpoints).
		RecentActivity(context.Background(), knownAddress, 5)
	if err != nil {
		t.Fatalf("a missing checkpoint must not fail the request: %v", err)
	}

	if page.AsOf.Block != nil {
		t.Fatalf("expected no block rather than a made up one, got %d", *page.AsOf.Block)
	}

	if !page.AsOf.Time.Equal(listClock) {
		t.Fatalf("expected the current time, got %s", page.AsOf.Time)
	}
}

func TestRecentActivityPropagatesACheckpointFailure(t *testing.T) {
	users, transactions := listFixtures(1)
	checkpoints := &stubCheckpoints{failWith: errors.New("database is down")}

	_, err := newListService(users, transactions, checkpoints).
		RecentActivity(context.Background(), knownAddress, 5)

	if err == nil {
		t.Fatal("expected the failure to surface")
	}

	if errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected the raw failure rather than not found, got %v", err)
	}
}

func TestRecentActivityPropagatesAUserLookupFailure(t *testing.T) {
	users, transactions := listFixtures(1)
	users.failWith = errors.New("database is down")

	_, err := newListService(users, transactions, nil).
		RecentActivity(context.Background(), knownAddress, 5)

	if err == nil {
		t.Fatal("expected the failure to surface")
	}

	if errors.Is(err, domain.ErrNotFound) || errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected the raw failure, got %v", err)
	}
}

func TestRecentActivityPropagatesARepositoryFailure(t *testing.T) {
	users, transactions := listFixtures(1)
	transactions.listFailWith = errors.New("database is down")

	_, err := newListService(users, transactions, nil).
		RecentActivity(context.Background(), knownAddress, 5)

	if err == nil {
		t.Fatal("expected the failure to surface")
	}
}
