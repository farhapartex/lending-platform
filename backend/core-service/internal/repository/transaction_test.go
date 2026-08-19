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
	"github.com/farhapartex/lending-platform/core-service/pkg/jsonb"
)

type ledger struct {
	users        domain.UserRepository
	transactions domain.TransactionRepository
	market       domain.Market
	user         domain.User
	baseTime     time.Time
	tx           *gorm.DB
}

func newLedger(t *testing.T) ledger {
	t.Helper()

	tx := newTx(t)
	market := newMarket(t, tx)

	users := repository.NewUserRepository(tx)

	user, err := users.EnsureByAddress(context.Background(), testChainID, addressFrom(nextUnique()))
	if err != nil {
		t.Fatalf("could not create the test user: %v", err)
	}

	return ledger{
		users:        users,
		transactions: repository.NewTransactionRepository(tx),
		market:       market,
		user:         user,
		baseTime:     time.Now().UTC().Truncate(time.Microsecond),
		tx:           tx,
	}
}

func (l ledger) add(t *testing.T, kind domain.TransactionKind, minutesAgo int, amount int64) domain.UserTransaction {
	t.Helper()

	event := newEvent(t, l.tx, l.market.ID, l.baseTime.Add(-time.Duration(minutesAgo)*time.Minute))

	transaction := domain.UserTransaction{
		EventID:     event.ID,
		UserID:      l.user.ID,
		MarketID:    l.market.ID,
		AssetID:     l.market.DebtAssetID,
		Kind:        kind,
		Amount:      bigmath.FromInt64(amount),
		BlockNumber: event.BlockNumber,
		BlockTime:   event.BlockTime,
		TxHash:      event.TxHash,
		LogIndex:    event.LogIndex,
		CreatedAt:   time.Now().UTC(),
	}

	if err := l.transactions.Insert(context.Background(), &transaction); err != nil {
		t.Fatalf("could not insert the transaction: %v", err)
	}

	return transaction
}

func newEvent(t *testing.T, tx *gorm.DB, marketID int64, blockTime time.Time) domain.ProtocolEvent {
	t.Helper()

	seed := nextUnique()

	event := domain.ProtocolEvent{
		ChainID:         testChainID,
		MarketID:        &marketID,
		EventType:       domain.EventTypeBorrow,
		ContractAddress: addressFrom(seed),
		BlockNumber:     seed % 1_000_000,
		BlockHash:       hashFrom(seed),
		BlockTime:       blockTime,
		TxHash:          hashFrom(seed + 1),
		LogIndex:        int32(seed % 100),
		Payload:         jsonb.FromBytes([]byte(`{"source":"test"}`)),
		CreatedAt:       time.Now().UTC(),
	}

	if err := tx.Omit("Market").Create(&event).Error; err != nil {
		t.Fatalf("could not insert the protocol event: %v", err)
	}

	return event
}

func hashFrom(seed int64) string {
	const hexDigits = "0123456789abcdef"

	body := make([]byte, 64)
	value := seed

	for index := len(body) - 1; index >= 0; index-- {
		body[index] = hexDigits[value&0xf]
		value >>= 4

		if value == 0 && index > 0 {
			value = seed >> 4
		}
	}

	return "0x" + string(body)
}

func TestUserEnsureByAddressCreatesThenReuses(t *testing.T) {
	tx := newTx(t)
	users := repository.NewUserRepository(tx)
	ctx := context.Background()

	address := addressFrom(nextUnique())

	first, err := users.EnsureByAddress(ctx, testChainID, address)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if first.ID < 1 {
		t.Fatalf("expected an assigned id, got %d", first.ID)
	}

	second, err := users.EnsureByAddress(ctx, testChainID, upper(address))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if second.ID != first.ID {
		t.Fatalf("expected the same user, got %d then %d", first.ID, second.ID)
	}
}

func TestUserEnsureStoresLowercaseAndChecksum(t *testing.T) {
	tx := newTx(t)
	users := repository.NewUserRepository(tx)

	address := addressFrom(nextUnique())

	created, err := users.EnsureByAddress(context.Background(), testChainID, upper(address))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if created.Address != lower(address) {
		t.Fatalf("expected a lowercase address, got %q", created.Address)
	}

	if lower(created.AddressChecksum) != created.Address {
		t.Fatalf("checksum %q does not match %q", created.AddressChecksum, created.Address)
	}
}

func TestUserByAddressNotFound(t *testing.T) {
	tx := newTx(t)
	users := repository.NewUserRepository(tx)

	_, err := users.ByAddress(context.Background(), addressFrom(nextUnique()))
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestUserRejectsBadAddress(t *testing.T) {
	tx := newTx(t)
	users := repository.NewUserRepository(tx)
	ctx := context.Background()

	if _, err := users.ByAddress(ctx, "nonsense"); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	if _, err := users.EnsureByAddress(ctx, testChainID, "0x123"); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestTransactionListReturnsNewestFirst(t *testing.T) {
	book := newLedger(t)

	book.add(t, domain.TransactionKindDeposit, 30, 1_000)
	middle := book.add(t, domain.TransactionKindBorrow, 20, 2_000)
	newest := book.add(t, domain.TransactionKindRepay, 10, 3_000)

	found, err := book.transactions.List(context.Background(), domain.TransactionQuery{UserID: book.user.ID})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 3 {
		t.Fatalf("expected three transactions, got %d", len(found))
	}

	if found[0].ID != newest.ID {
		t.Fatalf("expected the newest first, got id %d", found[0].ID)
	}

	if found[1].ID != middle.ID {
		t.Fatalf("expected the middle transaction second, got id %d", found[1].ID)
	}
}

func TestTransactionListPreloadsTheAsset(t *testing.T) {
	book := newLedger(t)
	book.add(t, domain.TransactionKindBorrow, 5, 5_100_000_000)

	found, err := book.transactions.List(context.Background(), domain.TransactionQuery{UserID: book.user.ID})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found[0].Asset == nil {
		t.Fatal("expected the asset to be preloaded")
	}

	if found[0].Asset.Decimals != 6 {
		t.Fatalf("expected a 6 decimal debt asset, got %d", found[0].Asset.Decimals)
	}
}

func TestTransactionListFiltersByKind(t *testing.T) {
	book := newLedger(t)

	book.add(t, domain.TransactionKindDeposit, 40, 1_000)
	borrow := book.add(t, domain.TransactionKindBorrow, 30, 2_000)
	book.add(t, domain.TransactionKindRepay, 20, 3_000)
	secondBorrow := book.add(t, domain.TransactionKindBorrow, 10, 4_000)

	found, err := book.transactions.List(context.Background(), domain.TransactionQuery{
		UserID: book.user.ID,
		Kinds:  []domain.TransactionKind{domain.TransactionKindBorrow},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 2 {
		t.Fatalf("expected two borrows, got %d", len(found))
	}

	if found[0].ID != secondBorrow.ID || found[1].ID != borrow.ID {
		t.Fatal("expected only borrows, newest first")
	}
}

func TestTransactionListFiltersBySeveralKinds(t *testing.T) {
	book := newLedger(t)

	book.add(t, domain.TransactionKindDeposit, 40, 1_000)
	book.add(t, domain.TransactionKindBorrow, 30, 2_000)
	book.add(t, domain.TransactionKindLiquidation, 20, 3_000)

	found, err := book.transactions.List(context.Background(), domain.TransactionQuery{
		UserID: book.user.ID,
		Kinds: []domain.TransactionKind{
			domain.TransactionKindBorrow,
			domain.TransactionKindLiquidation,
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 2 {
		t.Fatalf("expected two matches, got %d", len(found))
	}
}

func TestTransactionListFiltersByTimeWindow(t *testing.T) {
	book := newLedger(t)

	book.add(t, domain.TransactionKindDeposit, 120, 1_000)
	inWindow := book.add(t, domain.TransactionKindBorrow, 60, 2_000)
	book.add(t, domain.TransactionKindRepay, 5, 3_000)

	from := book.baseTime.Add(-90 * time.Minute)
	to := book.baseTime.Add(-30 * time.Minute)

	found, err := book.transactions.List(context.Background(), domain.TransactionQuery{
		UserID: book.user.ID,
		From:   &from,
		To:     &to,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 1 {
		t.Fatalf("expected one transaction inside the window, got %d", len(found))
	}

	if found[0].ID != inWindow.ID {
		t.Fatalf("expected the windowed transaction, got id %d", found[0].ID)
	}
}

func TestTransactionListAppliesOnlyFromWhenToIsAbsent(t *testing.T) {
	book := newLedger(t)

	book.add(t, domain.TransactionKindDeposit, 120, 1_000)
	book.add(t, domain.TransactionKindBorrow, 10, 2_000)

	from := book.baseTime.Add(-60 * time.Minute)

	found, err := book.transactions.List(context.Background(), domain.TransactionQuery{
		UserID: book.user.ID,
		From:   &from,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 1 {
		t.Fatalf("expected one recent transaction, got %d", len(found))
	}
}

func TestTransactionListRespectsTheLimit(t *testing.T) {
	book := newLedger(t)

	for index := range 5 {
		book.add(t, domain.TransactionKindDeposit, 50-index*5, int64(index+1)*1_000)
	}

	found, err := book.transactions.List(context.Background(), domain.TransactionQuery{
		UserID: book.user.ID,
		Limit:  2,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 2 {
		t.Fatalf("expected the limit to apply, got %d", len(found))
	}
}

func TestTransactionListCapsAnAbsurdLimit(t *testing.T) {
	book := newLedger(t)
	book.add(t, domain.TransactionKindDeposit, 5, 1_000)

	found, err := book.transactions.List(context.Background(), domain.TransactionQuery{
		UserID: book.user.ID,
		Limit:  1_000_000,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 1 {
		t.Fatalf("expected the single transaction, got %d", len(found))
	}
}

func TestTransactionListPagesWithoutOverlapOrGaps(t *testing.T) {
	book := newLedger(t)

	created := make([]int64, 0, 6)
	for index := range 6 {
		created = append(created, book.add(t, domain.TransactionKindDeposit, 60-index*5, int64(index+1)*1_000).ID)
	}

	seen := make([]int64, 0, len(created))
	after := cursor.Key{}

	for range 10 {
		page, err := book.transactions.List(context.Background(), domain.TransactionQuery{
			UserID: book.user.ID,
			Limit:  2,
			After:  after,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(page) == 0 {
			break
		}

		for _, transaction := range page {
			seen = append(seen, transaction.ID)
		}

		last := page[len(page)-1]
		after = cursor.Key{Time: last.BlockTime, ID: last.ID}
	}

	if len(seen) != len(created) {
		t.Fatalf("expected to page through %d transactions, saw %d", len(created), len(seen))
	}

	unique := make(map[int64]struct{}, len(seen))
	for _, id := range seen {
		if _, duplicated := unique[id]; duplicated {
			t.Fatalf("transaction %d appeared on more than one page", id)
		}

		unique[id] = struct{}{}
	}
}

func TestTransactionListCursorSkipsEarlierRows(t *testing.T) {
	book := newLedger(t)

	book.add(t, domain.TransactionKindDeposit, 30, 1_000)
	middle := book.add(t, domain.TransactionKindBorrow, 20, 2_000)
	book.add(t, domain.TransactionKindRepay, 10, 3_000)

	found, err := book.transactions.List(context.Background(), domain.TransactionQuery{
		UserID: book.user.ID,
		After:  cursor.Key{Time: middle.BlockTime, ID: middle.ID},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 1 {
		t.Fatalf("expected only rows older than the cursor, got %d", len(found))
	}

	if found[0].BlockTime.After(middle.BlockTime) {
		t.Fatal("expected the remaining row to be older than the cursor")
	}
}

func TestTransactionListIsScopedToTheUser(t *testing.T) {
	book := newLedger(t)
	book.add(t, domain.TransactionKindDeposit, 10, 1_000)

	other, err := book.users.EnsureByAddress(context.Background(), testChainID, addressFrom(nextUnique()))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	found, err := book.transactions.List(context.Background(), domain.TransactionQuery{UserID: other.ID})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 0 {
		t.Fatalf("expected another user to see nothing, got %d", len(found))
	}
}

func TestTransactionListReturnsEmptySliceNotNil(t *testing.T) {
	book := newLedger(t)

	found, err := book.transactions.List(context.Background(), domain.TransactionQuery{UserID: book.user.ID})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found == nil {
		t.Fatal("expected an empty slice rather than nil")
	}
}

func TestTransactionByID(t *testing.T) {
	book := newLedger(t)
	created := book.add(t, domain.TransactionKindBorrow, 10, 5_100_000_000)

	found, err := book.transactions.ByID(context.Background(), book.user.ID, created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.Amount.String() != "5100000000" {
		t.Fatalf("expected the amount to round trip, got %s", found.Amount.String())
	}

	if found.Asset == nil {
		t.Fatal("expected the asset to be preloaded")
	}
}

func TestTransactionByIDIsScopedToTheUser(t *testing.T) {
	book := newLedger(t)
	created := book.add(t, domain.TransactionKindBorrow, 10, 1_000)

	other, err := book.users.EnsureByAddress(context.Background(), testChainID, addressFrom(nextUnique()))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = book.transactions.ByID(context.Background(), other.ID, created.ID)
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected another user's lookup to be not found, got %v", err)
	}
}

func TestTransactionByIDNotFound(t *testing.T) {
	book := newLedger(t)

	_, err := book.transactions.ByID(context.Background(), book.user.ID, 999_999_999)
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestTransactionRejectsBadInput(t *testing.T) {
	book := newLedger(t)
	ctx := context.Background()

	if _, err := book.transactions.List(ctx, domain.TransactionQuery{UserID: 0}); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	if _, err := book.transactions.ByID(ctx, 0, 1); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	if _, err := book.transactions.ByID(ctx, 1, 0); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	if err := book.transactions.Insert(ctx, nil); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	orphan := domain.UserTransaction{UserID: 0}
	if err := book.transactions.Insert(ctx, &orphan); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestTransactionInsertRejectsANegativeAmount(t *testing.T) {
	book := newLedger(t)
	event := newEvent(t, book.tx, book.market.ID, book.baseTime)

	invalid := domain.UserTransaction{
		EventID:     event.ID,
		UserID:      book.user.ID,
		MarketID:    book.market.ID,
		AssetID:     book.market.DebtAssetID,
		Kind:        domain.TransactionKindBorrow,
		Amount:      bigmath.FromInt64(-1),
		BlockNumber: event.BlockNumber,
		BlockTime:   event.BlockTime,
		TxHash:      event.TxHash,
		LogIndex:    event.LogIndex,
		CreatedAt:   time.Now().UTC(),
	}

	if err := book.transactions.Insert(context.Background(), &invalid); err == nil {
		t.Fatal("expected the database to reject a non positive amount")
	}
}

func TestTransactionInsertRejectsADuplicateEvent(t *testing.T) {
	book := newLedger(t)
	first := book.add(t, domain.TransactionKindBorrow, 10, 1_000)

	duplicate := domain.UserTransaction{
		EventID:     first.EventID,
		UserID:      book.user.ID,
		MarketID:    book.market.ID,
		AssetID:     book.market.DebtAssetID,
		Kind:        domain.TransactionKindBorrow,
		Amount:      bigmath.FromInt64(2_000),
		BlockNumber: first.BlockNumber,
		BlockTime:   first.BlockTime,
		TxHash:      first.TxHash,
		LogIndex:    first.LogIndex,
		CreatedAt:   time.Now().UTC(),
	}

	err := book.transactions.Insert(context.Background(), &duplicate)
	if !errors.Is(err, domain.ErrAlreadyExists) {
		t.Fatalf("expected ErrAlreadyExists for a repeated event, got %v", err)
	}
}

func TestUserEnsureRefreshesLastSeen(t *testing.T) {
	tx := newTx(t)
	users := repository.NewUserRepository(tx)
	ctx := context.Background()

	address := addressFrom(nextUnique())

	first, err := users.EnsureByAddress(ctx, testChainID, address)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	time.Sleep(2 * time.Millisecond)

	second, err := users.EnsureByAddress(ctx, testChainID, address)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if second.ID != first.ID {
		t.Fatalf("expected the same row, got %d then %d", first.ID, second.ID)
	}

	if second.LastSeenAt == nil || first.LastSeenAt == nil {
		t.Fatal("expected last seen to be recorded")
	}

	if !second.LastSeenAt.After(*first.LastSeenAt) {
		t.Fatal("expected last seen to move forward on a repeat visit")
	}

	stored, err := users.ByAddress(ctx, address)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stored.FirstSeenAt.After(*first.LastSeenAt) {
		t.Fatal("expected first seen to stay at the original visit")
	}
}

func TestTransactionRepositoriesSurfaceDatabaseFailures(t *testing.T) {
	book := newLedger(t)
	ctx := context.Background()

	existing := book.add(t, domain.TransactionKindBorrow, 10, 1_000)
	address := addressFrom(nextUnique())

	book.tx.Rollback()

	cases := []struct {
		name string
		call func() error
	}{
		{name: "transaction list", call: func() error {
			_, err := book.transactions.List(ctx, domain.TransactionQuery{UserID: book.user.ID})

			return err
		}},
		{name: "transaction by id", call: func() error {
			_, err := book.transactions.ByID(ctx, book.user.ID, existing.ID)

			return err
		}},
		{name: "transaction insert", call: func() error {
			clone := existing
			clone.ID = 0

			return book.transactions.Insert(ctx, &clone)
		}},
		{name: "user by address", call: func() error {
			_, err := book.users.ByAddress(ctx, address)

			return err
		}},
		{name: "user ensure", call: func() error {
			_, err := book.users.EnsureByAddress(ctx, testChainID, address)

			return err
		}},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.call()

			if err == nil {
				t.Fatal("expected an error once the transaction is gone")
			}

			if errors.Is(err, domain.ErrInvalidInput) {
				t.Fatalf("expected a database failure rather than a validation error, got %v", err)
			}
		})
	}
}
