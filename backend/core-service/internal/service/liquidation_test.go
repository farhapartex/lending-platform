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

type stubLiquidations struct {
	page         []domain.Liquidation
	byKey        map[int64]domain.Liquidation
	listFailWith error
	byIDFailWith error
	lastQuery    domain.LiquidationQuery
	lastID       int64
	listCount    int
	byIDCount    int
}

func (s *stubLiquidations) List(
	_ context.Context,
	query domain.LiquidationQuery,
) ([]domain.Liquidation, error) {
	s.lastQuery = query
	s.listCount++

	if s.listFailWith != nil {
		return nil, s.listFailWith
	}

	if query.Limit > 0 && len(s.page) > query.Limit {
		return s.page[:query.Limit], nil
	}

	return s.page, nil
}

func (s *stubLiquidations) ByID(_ context.Context, id int64) (domain.Liquidation, error) {
	s.lastID = id
	s.byIDCount++

	if s.byIDFailWith != nil {
		return domain.Liquidation{}, s.byIDFailWith
	}

	found, ok := s.byKey[id]
	if !ok {
		return domain.Liquidation{}, domain.ErrNotFound
	}

	return found, nil
}

func (s *stubLiquidations) Insert(context.Context, *domain.Liquidation) error {
	return errors.New("not expected in these tests")
}

func liquidationRows(count int) []domain.Liquidation {
	rows := make([]domain.Liquidation, 0, count)

	for index := 0; index < count; index++ {
		rows = append(rows, domain.Liquidation{
			ID:               int64(900 - index),
			DebtRepaid:       bigmath.FromInt64(5_100_000_000),
			CollateralSeized: bigmath.FromInt64(1_800_000_000),
			BonusAmount:      bigmath.FromInt64(25_500_000_000),
			ShortfallAmount:  bigmath.FromInt64(0),
			BlockNumber:      int64(500 - index),
			BlockTime:        listClock.Add(-time.Duration(index) * time.Minute),
		})
	}

	return rows
}

func newLiquidationService(
	liquidations *stubLiquidations,
	checkpoints domain.CheckpointRepository,
) domain.LiquidationService {
	return service.NewLiquidationService(service.LiquidationServiceParams{
		Liquidations: liquidations,
		Checkpoints:  checkpoints,
		Now:          func() time.Time { return listClock },
	})
}

func TestLiquidationListReturnsThePage(t *testing.T) {
	stub := &stubLiquidations{page: liquidationRows(3)}

	page, err := newLiquidationService(stub, nil).
		List(context.Background(), domain.LiquidationListRequest{Limit: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(page.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(page.Items))
	}
}

func TestLiquidationListAsksForOneMoreRowThanThePageSize(t *testing.T) {
	stub := &stubLiquidations{page: liquidationRows(3)}

	if _, err := newLiquidationService(stub, nil).
		List(context.Background(), domain.LiquidationListRequest{Limit: 10}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stub.lastQuery.Limit != 11 {
		t.Fatalf("expected a probe row, got limit %d", stub.lastQuery.Limit)
	}
}

func TestLiquidationListHidesTheProbeRowAndReturnsACursor(t *testing.T) {
	stub := &stubLiquidations{page: liquidationRows(6)}

	page, err := newLiquidationService(stub, nil).
		List(context.Background(), domain.LiquidationListRequest{Limit: 5})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(page.Items) != 5 {
		t.Fatalf("expected the probe row to be trimmed, got %d", len(page.Items))
	}

	last := page.Items[len(page.Items)-1]

	if page.NextCursor.ID != last.ID || !page.NextCursor.Time.Equal(last.BlockTime) {
		t.Fatalf("expected the cursor to point at the last returned row, got %+v", page.NextCursor)
	}
}

func TestLiquidationListReturnsNoCursorOnTheLastPage(t *testing.T) {
	for _, count := range []int{0, 1, 4, 5} {
		stub := &stubLiquidations{page: liquidationRows(count)}

		page, err := newLiquidationService(stub, nil).
			List(context.Background(), domain.LiquidationListRequest{Limit: 5})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !page.NextCursor.IsZero() {
			t.Fatalf("expected no cursor with %d rows, got %+v", count, page.NextCursor)
		}
	}
}

func TestLiquidationListAppliesTheDefaultPageSize(t *testing.T) {
	for _, limit := range []int{0, -1} {
		stub := &stubLiquidations{page: liquidationRows(1)}

		if _, err := newLiquidationService(stub, nil).
			List(context.Background(), domain.LiquidationListRequest{Limit: limit}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		want := service.DefaultLiquidationPageSize + 1

		if stub.lastQuery.Limit != want {
			t.Fatalf("expected %d for a requested %d, got %d", want, limit, stub.lastQuery.Limit)
		}
	}
}

func TestLiquidationListCapsThePageSize(t *testing.T) {
	stub := &stubLiquidations{page: liquidationRows(1)}

	if _, err := newLiquidationService(stub, nil).
		List(context.Background(), domain.LiquidationListRequest{Limit: 5_000}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stub.lastQuery.Limit != service.MaxLiquidationPageSize+1 {
		t.Fatalf("expected the page size to be capped, got %d", stub.lastQuery.Limit)
	}
}

func TestLiquidationListPassesTheMarketFilterAndCursorThrough(t *testing.T) {
	stub := &stubLiquidations{page: liquidationRows(1)}
	marketID := int64(7)
	after := cursor.Key{Time: listClock.Add(-time.Hour), ID: 500}

	if _, err := newLiquidationService(stub, nil).List(context.Background(), domain.LiquidationListRequest{
		MarketID: &marketID,
		After:    after,
		Limit:    10,
	}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stub.lastQuery.MarketID == nil || *stub.lastQuery.MarketID != 7 {
		t.Fatalf("expected the market filter to reach the repository, got %v", stub.lastQuery.MarketID)
	}

	if stub.lastQuery.After.ID != after.ID {
		t.Fatalf("expected the cursor to reach the repository, got %+v", stub.lastQuery.After)
	}
}

func TestLiquidationListRejectsABadMarketID(t *testing.T) {
	for _, marketID := range []int64{0, -1} {
		stub := &stubLiquidations{page: liquidationRows(1)}
		candidate := marketID

		_, err := newLiquidationService(stub, nil).
			List(context.Background(), domain.LiquidationListRequest{MarketID: &candidate})

		if !errors.Is(err, domain.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput for market %d, got %v", marketID, err)
		}

		if stub.listCount != 0 {
			t.Fatal("expected a bad market id to be refused before any lookup")
		}
	}
}

func TestLiquidationListAcceptsNoMarketFilter(t *testing.T) {
	stub := &stubLiquidations{page: liquidationRows(1)}

	if _, err := newLiquidationService(stub, nil).
		List(context.Background(), domain.LiquidationListRequest{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stub.lastQuery.MarketID != nil {
		t.Fatalf("expected no market filter, got %v", stub.lastQuery.MarketID)
	}
}

func TestLiquidationListReportsTheIndexedBlock(t *testing.T) {
	stub := &stubLiquidations{page: liquidationRows(2)}
	updated := listClock.Add(-90 * time.Second)
	checkpoints := &stubCheckpoints{checkpoint: domain.IndexerCheckpoint{
		LastProcessedBlock: 4_218,
		UpdatedAt:          updated,
	}}

	page, err := newLiquidationService(stub, checkpoints).
		List(context.Background(), domain.LiquidationListRequest{})
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

func TestLiquidationListFallsBackToTheClockWhenNothingHasBeenIndexed(t *testing.T) {
	stub := &stubLiquidations{page: liquidationRows(1)}
	checkpoints := &stubCheckpoints{failWith: domain.ErrNotFound}

	page, err := newLiquidationService(stub, checkpoints).
		List(context.Background(), domain.LiquidationListRequest{})
	if err != nil {
		t.Fatalf("a missing checkpoint must not fail the request: %v", err)
	}

	if page.AsOf.Block != nil {
		t.Fatalf("expected no block, got %d", *page.AsOf.Block)
	}

	if !page.AsOf.Time.Equal(listClock) {
		t.Fatalf("expected the current time, got %s", page.AsOf.Time)
	}
}

func TestLiquidationListWorksWithoutACheckpointRepository(t *testing.T) {
	stub := &stubLiquidations{page: liquidationRows(1)}

	page, err := newLiquidationService(stub, nil).
		List(context.Background(), domain.LiquidationListRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if page.AsOf.Block != nil {
		t.Fatalf("expected no block, got %d", *page.AsOf.Block)
	}
}

func TestLiquidationListPropagatesACheckpointFailure(t *testing.T) {
	stub := &stubLiquidations{page: liquidationRows(1)}
	checkpoints := &stubCheckpoints{failWith: errors.New("database is down")}

	_, err := newLiquidationService(stub, checkpoints).
		List(context.Background(), domain.LiquidationListRequest{})

	if err == nil {
		t.Fatal("expected the failure to surface")
	}

	if errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected the raw failure, got %v", err)
	}
}

func TestLiquidationListPropagatesARepositoryFailure(t *testing.T) {
	stub := &stubLiquidations{listFailWith: errors.New("database is down")}

	if _, err := newLiquidationService(stub, nil).
		List(context.Background(), domain.LiquidationListRequest{}); err == nil {
		t.Fatal("expected the failure to surface")
	}
}

func TestLiquidationByIDReturnsTheReceipt(t *testing.T) {
	stub := &stubLiquidations{byKey: map[int64]domain.Liquidation{
		42: {ID: 42, DebtRepaid: bigmath.FromInt64(5_100_000_000)},
	}}

	found, err := newLiquidationService(stub, nil).ByID(context.Background(), 42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.ID != 42 {
		t.Fatalf("expected receipt 42, got %d", found.ID)
	}

	if stub.lastID != 42 {
		t.Fatalf("expected the id to reach the repository, got %d", stub.lastID)
	}
}

func TestLiquidationByIDRejectsANonPositiveID(t *testing.T) {
	stub := &stubLiquidations{byKey: map[int64]domain.Liquidation{}}

	for _, id := range []int64{0, -1} {
		_, err := newLiquidationService(stub, nil).ByID(context.Background(), id)

		if !errors.Is(err, domain.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput for id %d, got %v", id, err)
		}
	}

	if stub.byIDCount != 0 {
		t.Fatal("expected a bad id to be refused before any lookup")
	}
}

func TestLiquidationByIDReportsAMissingReceipt(t *testing.T) {
	stub := &stubLiquidations{byKey: map[int64]domain.Liquidation{}}

	_, err := newLiquidationService(stub, nil).ByID(context.Background(), 42)

	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestLiquidationByIDPropagatesAFailure(t *testing.T) {
	stub := &stubLiquidations{byIDFailWith: errors.New("database is down")}

	_, err := newLiquidationService(stub, nil).ByID(context.Background(), 42)

	if err == nil {
		t.Fatal("expected the failure to surface")
	}

	if errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected the raw failure rather than not found, got %v", err)
	}
}

func TestLiquidationServiceDefaultsTheClock(t *testing.T) {
	stub := &stubLiquidations{page: liquidationRows(1)}

	page, err := service.NewLiquidationService(service.LiquidationServiceParams{Liquidations: stub}).
		List(context.Background(), domain.LiquidationListRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if page.AsOf.Time.IsZero() {
		t.Fatal("expected a real timestamp from the default clock")
	}
}
