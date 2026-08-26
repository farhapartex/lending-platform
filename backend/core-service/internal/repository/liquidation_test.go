package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/repository"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
)

type liquidationBook struct {
	liquidations domain.LiquidationRepository
	market       domain.Market
	borrower     domain.User
	liquidator   domain.User
	baseTime     time.Time
	tx           *gorm.DB
}

func newLiquidationBook(t *testing.T) liquidationBook {
	t.Helper()

	tx := newTx(t)
	market := newMarket(t, tx)
	users := repository.NewUserRepository(tx)

	borrower, err := users.EnsureByAddress(context.Background(), testChainID, addressFrom(nextUnique()))
	if err != nil {
		t.Fatalf("could not create the borrower: %v", err)
	}

	liquidator, err := users.EnsureByAddress(context.Background(), testChainID, addressFrom(nextUnique()))
	if err != nil {
		t.Fatalf("could not create the liquidator: %v", err)
	}

	return liquidationBook{
		liquidations: repository.NewLiquidationRepository(tx),
		market:       market,
		borrower:     borrower,
		liquidator:   liquidator,
		baseTime:     time.Now().UTC().Truncate(time.Microsecond),
		tx:           tx,
	}
}

func (b liquidationBook) add(t *testing.T, minutesAgo int, shortfall int64) domain.Liquidation {
	t.Helper()

	return b.addForMarket(t, b.market, minutesAgo, shortfall)
}

func (b liquidationBook) addForMarket(
	t *testing.T,
	market domain.Market,
	minutesAgo int,
	shortfall int64,
) domain.Liquidation {
	t.Helper()

	event := newEvent(t, b.tx, market.ID, b.baseTime.Add(-time.Duration(minutesAgo)*time.Minute))
	health := int32(9_098)

	liquidation := domain.Liquidation{
		EventID:               event.ID,
		MarketID:              market.ID,
		BorrowerUserID:        b.borrower.ID,
		LiquidatorUserID:      b.liquidator.ID,
		DebtRepaid:            bigmath.FromInt64(5_100_000_000),
		CollateralSeized:      bigmath.MustFromString("1846551724137931035"),
		BonusAmount:           bigmath.FromInt64(25_500_000_000),
		HealthFactorBeforeBps: &health,
		TriggerPrice:          bigmath.FromInt64(220_000_000_000),
		TriggerPriceDecimals:  8,
		ShortfallAmount:       bigmath.FromInt64(shortfall),
		BlockNumber:           event.BlockNumber,
		BlockTime:             event.BlockTime,
		TxHash:                event.TxHash,
		CreatedAt:             time.Now().UTC(),
	}

	if err := b.liquidations.Insert(context.Background(), &liquidation); err != nil {
		t.Fatalf("could not insert the liquidation: %v", err)
	}

	return liquidation
}

func TestLiquidationListReturnsNewestFirst(t *testing.T) {
	book := newLiquidationBook(t)
	oldest := book.add(t, 30, 0)
	newest := book.add(t, 1, 0)
	middle := book.add(t, 10, 0)

	found, err := book.liquidations.List(context.Background(), domain.LiquidationQuery{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 3 {
		t.Fatalf("expected 3 rows, got %d", len(found))
	}

	want := []int64{newest.ID, middle.ID, oldest.ID}

	for index, id := range want {
		if found[index].ID != id {
			t.Fatalf("expected newest first %v, got %d at %d", want, found[index].ID, index)
		}
	}
}

func TestLiquidationListPreloadsBothPartiesAndAssets(t *testing.T) {
	book := newLiquidationBook(t)
	book.add(t, 1, 0)

	found, err := book.liquidations.List(context.Background(), domain.LiquidationQuery{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	row := found[0]

	if row.Borrower == nil || row.Borrower.Address != book.borrower.Address {
		t.Fatalf("expected the borrower to be loaded, got %+v", row.Borrower)
	}

	if row.Liquidator == nil || row.Liquidator.Address != book.liquidator.Address {
		t.Fatalf("expected the liquidator to be loaded, got %+v", row.Liquidator)
	}

	if row.Market == nil {
		t.Fatal("expected the market to be loaded")
	}

	if row.Market.CollateralAsset == nil || row.Market.DebtAsset == nil {
		t.Fatal("expected both assets to be loaded so amounts can be formatted without another query")
	}
}

func TestLiquidationListFiltersByMarket(t *testing.T) {
	book := newLiquidationBook(t)
	other := newMarket(t, book.tx)

	book.add(t, 5, 0)
	book.add(t, 6, 0)
	elsewhere := book.addForMarket(t, other, 7, 0)

	found, err := book.liquidations.List(context.Background(), domain.LiquidationQuery{MarketID: &other.ID})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 1 || found[0].ID != elsewhere.ID {
		t.Fatalf("expected only the other market's row, got %d rows", len(found))
	}
}

func TestLiquidationListRejectsABadMarketID(t *testing.T) {
	book := newLiquidationBook(t)

	for _, marketID := range []int64{0, -1} {
		_, err := book.liquidations.List(context.Background(), domain.LiquidationQuery{MarketID: &marketID})

		if !errors.Is(err, domain.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput for market %d, got %v", marketID, err)
		}
	}
}

func TestLiquidationListRespectsTheLimit(t *testing.T) {
	book := newLiquidationBook(t)
	book.add(t, 1, 0)
	book.add(t, 2, 0)
	book.add(t, 3, 0)

	found, err := book.liquidations.List(context.Background(), domain.LiquidationQuery{Limit: 2})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 2 {
		t.Fatalf("expected the limit to apply, got %d", len(found))
	}
}

func TestLiquidationListCapsAnAbsurdLimit(t *testing.T) {
	book := newLiquidationBook(t)
	book.add(t, 1, 0)

	found, err := book.liquidations.List(context.Background(), domain.LiquidationQuery{Limit: 1_000_000})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 1 {
		t.Fatalf("expected the single row, got %d", len(found))
	}
}

func TestLiquidationListPagesWithoutOverlapOrGaps(t *testing.T) {
	book := newLiquidationBook(t)

	for minutesAgo := 1; minutesAgo <= 5; minutesAgo++ {
		book.add(t, minutesAgo, 0)
	}

	seen := make(map[int64]struct{})
	after := cursor.Key{}

	for page := 0; page < 5; page++ {
		found, err := book.liquidations.List(context.Background(), domain.LiquidationQuery{
			After: after,
			Limit: 2,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(found) == 0 {
			break
		}

		for _, row := range found {
			if _, duplicate := seen[row.ID]; duplicate {
				t.Fatalf("row %d appeared on more than one page", row.ID)
			}

			seen[row.ID] = struct{}{}
		}

		last := found[len(found)-1]
		after = cursor.Key{Time: last.BlockTime, ID: last.ID}
	}

	if len(seen) != 5 {
		t.Fatalf("expected every row exactly once, got %d", len(seen))
	}
}

func TestLiquidationListCursorSkipsEarlierRows(t *testing.T) {
	book := newLiquidationBook(t)
	newest := book.add(t, 1, 0)
	older := book.add(t, 10, 0)

	found, err := book.liquidations.List(context.Background(), domain.LiquidationQuery{
		After: cursor.Key{Time: newest.BlockTime, ID: newest.ID},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 1 || found[0].ID != older.ID {
		t.Fatalf("expected only the older row, got %d rows", len(found))
	}
}

func TestLiquidationListReturnsEmptySliceNotNil(t *testing.T) {
	book := newLiquidationBook(t)

	found, err := book.liquidations.List(context.Background(), domain.LiquidationQuery{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found == nil {
		t.Fatal("expected an empty slice so callers can range without a nil check")
	}

	if len(found) != 0 {
		t.Fatalf("expected no rows, got %d", len(found))
	}
}

func TestLiquidationByIDReturnsTheReceipt(t *testing.T) {
	book := newLiquidationBook(t)
	stored := book.add(t, 5, 0)

	found, err := book.liquidations.ByID(context.Background(), stored.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.ID != stored.ID {
		t.Fatalf("expected row %d, got %d", stored.ID, found.ID)
	}

	if found.TriggerPrice.String() != "220000000000" || found.TriggerPriceDecimals != 8 {
		t.Fatalf("expected the trigger price to survive the round trip, got %s at %d decimals",
			found.TriggerPrice.String(), found.TriggerPriceDecimals)
	}

	if found.CollateralSeized.String() != "1846551724137931035" {
		t.Fatalf("expected an exact 18 decimal amount, got %s", found.CollateralSeized.String())
	}

	if found.Borrower == nil || found.Market == nil || found.Market.DebtAsset == nil {
		t.Fatal("expected the receipt to arrive with its relations loaded")
	}
}

func TestLiquidationByIDReportsAMissingRow(t *testing.T) {
	book := newLiquidationBook(t)

	_, err := book.liquidations.ByID(context.Background(), 987_654)

	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestLiquidationByIDRejectsABadID(t *testing.T) {
	book := newLiquidationBook(t)

	for _, id := range []int64{0, -1} {
		_, err := book.liquidations.ByID(context.Background(), id)

		if !errors.Is(err, domain.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput for id %d, got %v", id, err)
		}
	}
}

func TestLiquidationInsertRejectsBadInput(t *testing.T) {
	book := newLiquidationBook(t)

	if err := book.liquidations.Insert(context.Background(), nil); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput for a nil row, got %v", err)
	}

	cases := map[string]domain.Liquidation{
		"no borrower":   {LiquidatorUserID: book.liquidator.ID},
		"no liquidator": {BorrowerUserID: book.borrower.ID},
	}

	for name, row := range cases {
		t.Run(name, func(t *testing.T) {
			candidate := row

			if err := book.liquidations.Insert(context.Background(), &candidate); !errors.Is(err, domain.ErrInvalidInput) {
				t.Fatalf("expected ErrInvalidInput, got %v", err)
			}
		})
	}
}

func TestLiquidationInsertRefusesTheSamePartyOnBothSides(t *testing.T) {
	book := newLiquidationBook(t)
	event := newEvent(t, book.tx, book.market.ID, book.baseTime)

	liquidation := domain.Liquidation{
		EventID:              event.ID,
		MarketID:             book.market.ID,
		BorrowerUserID:       book.borrower.ID,
		LiquidatorUserID:     book.borrower.ID,
		DebtRepaid:           bigmath.FromInt64(1_000),
		CollateralSeized:     bigmath.FromInt64(1),
		BonusAmount:          bigmath.FromInt64(0),
		TriggerPrice:         bigmath.FromInt64(1),
		TriggerPriceDecimals: 8,
		ShortfallAmount:      bigmath.FromInt64(0),
		BlockNumber:          event.BlockNumber,
		BlockTime:            event.BlockTime,
		TxHash:               event.TxHash,
		CreatedAt:            time.Now().UTC(),
	}

	if err := book.liquidations.Insert(context.Background(), &liquidation); err == nil {
		t.Fatal("a wallet cannot liquidate itself, so the database must refuse this")
	}
}

func TestLiquidationInsertRefusesANonPositiveDebt(t *testing.T) {
	book := newLiquidationBook(t)
	event := newEvent(t, book.tx, book.market.ID, book.baseTime)

	liquidation := domain.Liquidation{
		EventID:              event.ID,
		MarketID:             book.market.ID,
		BorrowerUserID:       book.borrower.ID,
		LiquidatorUserID:     book.liquidator.ID,
		DebtRepaid:           bigmath.FromInt64(0),
		CollateralSeized:     bigmath.FromInt64(1),
		BonusAmount:          bigmath.FromInt64(0),
		TriggerPrice:         bigmath.FromInt64(1),
		TriggerPriceDecimals: 8,
		ShortfallAmount:      bigmath.FromInt64(0),
		BlockNumber:          event.BlockNumber,
		BlockTime:            event.BlockTime,
		TxHash:               event.TxHash,
		CreatedAt:            time.Now().UTC(),
	}

	if err := book.liquidations.Insert(context.Background(), &liquidation); err == nil {
		t.Fatal("a liquidation that repaid nothing is not a liquidation")
	}
}

func TestLiquidationInsertRefusesADuplicateEvent(t *testing.T) {
	book := newLiquidationBook(t)
	first := book.add(t, 1, 0)

	duplicate := domain.Liquidation{
		EventID:              first.EventID,
		MarketID:             book.market.ID,
		BorrowerUserID:       book.borrower.ID,
		LiquidatorUserID:     book.liquidator.ID,
		DebtRepaid:           bigmath.FromInt64(1_000),
		CollateralSeized:     bigmath.FromInt64(1),
		BonusAmount:          bigmath.FromInt64(0),
		TriggerPrice:         bigmath.FromInt64(1),
		TriggerPriceDecimals: 8,
		ShortfallAmount:      bigmath.FromInt64(0),
		BlockNumber:          first.BlockNumber,
		BlockTime:            first.BlockTime,
		TxHash:               first.TxHash,
		CreatedAt:            time.Now().UTC(),
	}

	err := book.liquidations.Insert(context.Background(), &duplicate)

	if !errors.Is(err, domain.ErrAlreadyExists) {
		t.Fatalf("replaying one event must not create a second receipt, got %v", err)
	}
}

func TestLiquidationStoresAShortfall(t *testing.T) {
	book := newLiquidationBook(t)
	stored := book.add(t, 1, 4_200_000_000)

	found, err := book.liquidations.ByID(context.Background(), stored.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.ShortfallAmount.String() != "4200000000" {
		t.Fatalf("expected the shortfall to be kept, got %s", found.ShortfallAmount.String())
	}
}

func TestLiquidationRepositorySurfacesDatabaseFailures(t *testing.T) {
	book := newLiquidationBook(t)

	if err := book.tx.Exec("DROP TABLE liquidations").Error; err != nil {
		t.Fatalf("could not drop the table: %v", err)
	}

	if _, err := book.liquidations.List(context.Background(), domain.LiquidationQuery{}); err == nil {
		t.Fatal("expected a list failure to surface")
	}

	if _, err := book.liquidations.ByID(context.Background(), 1); err == nil {
		t.Fatal("expected a lookup failure to surface")
	}

	row := domain.Liquidation{BorrowerUserID: book.borrower.ID, LiquidatorUserID: book.liquidator.ID}

	if err := book.liquidations.Insert(context.Background(), &row); err == nil {
		t.Fatal("expected an insert failure to surface")
	}
}
