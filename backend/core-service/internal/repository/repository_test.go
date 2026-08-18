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
)

type repositorySet struct {
	assets    domain.AssetRepository
	markets   domain.MarketRepository
	snapshots domain.MarketSnapshotRepository
}

func repositoryFor(tx *gorm.DB) repositorySet {
	return repositorySet{
		assets:    repository.NewAssetRepository(tx),
		markets:   repository.NewMarketRepository(tx),
		snapshots: repository.NewMarketSnapshotRepository(tx),
	}
}

func TestAssetUpsertThenRead(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	created := newAsset(t, tx, "WETH", 18)

	if created.ID < 1 {
		t.Fatalf("expected an assigned id, got %d", created.ID)
	}

	byID, err := repos.assets.ByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if byID.Symbol != created.Symbol {
		t.Fatalf("expected symbol %q, got %q", created.Symbol, byID.Symbol)
	}

	byAddress, err := repos.assets.ByAddress(ctx, testChainID, created.Address)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if byAddress.ID != created.ID {
		t.Fatalf("expected id %d, got %d", created.ID, byAddress.ID)
	}
}

func TestAssetUpsertStoresChecksumAndLowercase(t *testing.T) {
	tx := newTx(t)
	ctx := context.Background()

	created := newAsset(t, tx, "WETH", 18)

	if created.Address != lower(created.Address) {
		t.Fatalf("expected a lowercase address, got %q", created.Address)
	}

	if created.AddressChecksum == created.Address {
		t.Fatalf("expected a checksummed variant, got %q", created.AddressChecksum)
	}

	if lower(created.AddressChecksum) != created.Address {
		t.Fatalf("checksum %q does not match address %q", created.AddressChecksum, created.Address)
	}

	_ = ctx
}

func TestAssetByAddressIsCaseInsensitive(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	created := newAsset(t, tx, "WETH", 18)

	found, err := repos.assets.ByAddress(ctx, testChainID, upper(created.Address))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.ID != created.ID {
		t.Fatalf("expected id %d, got %d", created.ID, found.ID)
	}
}

func TestAssetUpsertUpdatesInsteadOfDuplicating(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	created := newAsset(t, tx, "WETH", 18)

	updated := created
	updated.ID = 0
	updated.Name = "renamed asset"
	updated.IsBorrowable = true
	updated.UpdatedAt = time.Now().UTC()

	if err := repos.assets.Upsert(ctx, &updated); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	found, err := repos.assets.ByAddress(ctx, testChainID, created.Address)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.Name != "renamed asset" {
		t.Fatalf("expected the update to apply, got %q", found.Name)
	}

	if !found.IsBorrowable {
		t.Fatal("expected is_borrowable to be updated")
	}

	all, err := repos.assets.List(ctx, testChainID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	matches := 0
	for _, asset := range all {
		if asset.Address == created.Address {
			matches++
		}
	}

	if matches != 1 {
		t.Fatalf("expected exactly one row for the address, found %d", matches)
	}
}

func TestAssetNotFound(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	if _, err := repos.assets.ByID(ctx, 999_999_999); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	missing := addressFrom(nextUnique())
	if _, err := repos.assets.ByAddress(ctx, testChainID, missing); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestAssetRejectsBadInput(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	if _, err := repos.assets.ByID(ctx, 0); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput for id 0, got %v", err)
	}

	if _, err := repos.assets.ByID(ctx, -5); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput for a negative id, got %v", err)
	}

	if _, err := repos.assets.ByAddress(ctx, testChainID, "not-an-address"); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	if err := repos.assets.Upsert(ctx, nil); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput for a nil asset, got %v", err)
	}

	bad := domain.Asset{ChainID: testChainID, Address: "0x123", Symbol: "BAD", Name: "bad"}
	if err := repos.assets.Upsert(ctx, &bad); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput for a malformed address, got %v", err)
	}
}

func TestAssetListIsScopedToTheChain(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	newAsset(t, tx, "WETH", 18)

	otherChain := domain.Asset{
		ChainID:   999999,
		Address:   addressFrom(nextUnique()),
		Symbol:    "OTHER-" + shortSuffix(nextUnique()),
		Name:      "other chain asset",
		Decimals:  18,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := repos.assets.Upsert(ctx, &otherChain); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assets, err := repos.assets.List(ctx, 999999)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, asset := range assets {
		if asset.ChainID != 999999 {
			t.Fatalf("expected only chain 999999 assets, got chain %d", asset.ChainID)
		}
	}

	if len(assets) != 1 {
		t.Fatalf("expected exactly one asset on the other chain, got %d", len(assets))
	}
}

func TestMarketUpsertThenRead(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	created := newMarket(t, tx)

	byID, err := repos.markets.ByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if byID.PoolAddress != created.PoolAddress {
		t.Fatalf("expected pool %q, got %q", created.PoolAddress, byID.PoolAddress)
	}

	if byID.MaxLTVBps != 7500 || byID.LiquidationThresholdBps != 8000 {
		t.Fatalf("unexpected risk parameters: %+v", byID)
	}

	if byID.MinDeposit.String() != "1000000" {
		t.Fatalf("expected a min deposit of 1000000, got %s", byID.MinDeposit.String())
	}
}

func TestMarketPreloadsBothAssets(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	created := newMarket(t, tx)

	found, err := repos.markets.ByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.CollateralAsset == nil || found.DebtAsset == nil {
		t.Fatal("expected both assets to be preloaded")
	}

	if found.CollateralAsset.Decimals != 18 {
		t.Fatalf("expected 18 decimal collateral, got %d", found.CollateralAsset.Decimals)
	}

	if found.DebtAsset.Decimals != 6 {
		t.Fatalf("expected 6 decimal debt asset, got %d", found.DebtAsset.Decimals)
	}
}

func TestMarketByPoolAddress(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	created := newMarket(t, tx)

	found, err := repos.markets.ByPoolAddress(ctx, testChainID, upper(created.PoolAddress))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.ID != created.ID {
		t.Fatalf("expected id %d, got %d", created.ID, found.ID)
	}

	if found.CollateralAsset == nil {
		t.Fatal("expected assets to be preloaded on the address lookup too")
	}
}

func TestMarketUpsertUpdatesRiskParameters(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	created := newMarket(t, tx)

	updated := created
	updated.ID = 0
	updated.MaxLTVBps = 7000
	updated.LiquidationThresholdBps = 7600
	updated.Status = domain.MarketStatusDepositOnly
	updated.UpdatedAt = time.Now().UTC()

	if err := repos.markets.Upsert(ctx, &updated); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	found, err := repos.markets.ByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.MaxLTVBps != 7000 || found.LiquidationThresholdBps != 7600 {
		t.Fatalf("expected updated risk parameters, got %+v", found)
	}

	if found.Status != domain.MarketStatusDepositOnly {
		t.Fatalf("expected the status to change, got %q", found.Status)
	}
}

func TestMarketRespectsTheLTVCheckConstraint(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	collateral := newAsset(t, tx, "WETH", 18)
	debt := newAsset(t, tx, "USDC", 6)

	seed := nextUnique()
	now := time.Now().UTC()

	invalid := domain.Market{
		ChainID:                   testChainID,
		CollateralAssetID:         collateral.ID,
		DebtAssetID:               debt.ID,
		PoolAddress:               addressFrom(seed),
		CollateralVaultAddress:    addressFrom(seed + 1),
		ControllerAddress:         addressFrom(seed + 2),
		LiquidationManagerAddress: addressFrom(seed + 3),
		InterestRateModelAddress:  addressFrom(seed + 4),
		OracleAdapterAddress:      addressFrom(seed + 5),
		MaxLTVBps:                 8000,
		LiquidationThresholdBps:   8000,
		LiquidationBonusBps:       500,
		ReserveFactorBps:          1000,
		RecommendedLTVBps:         6500,
		KinkUtilizationBps:        8000,
		MinDeposit:                bigmath.FromInt64(1_000_000),
		Status:                    domain.MarketStatusActive,
		DeployedAtBlock:           1,
		CreatedAt:                 now,
		UpdatedAt:                 now,
	}

	if err := repos.markets.Upsert(ctx, &invalid); err == nil {
		t.Fatal("expected the database to reject max ltv equal to the liquidation threshold")
	}
}

func TestMarketRejectsBadInput(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	if _, err := repos.markets.ByID(ctx, 0); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	if _, err := repos.markets.ByPoolAddress(ctx, testChainID, "0xzz"); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	if err := repos.markets.Upsert(ctx, nil); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	bad := domain.Market{PoolAddress: "nope"}
	if err := repos.markets.Upsert(ctx, &bad); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput for a malformed pool address, got %v", err)
	}
}

func TestMarketNotFound(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	if _, err := repos.markets.ByID(ctx, 999_999_999); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	missing := addressFrom(nextUnique())
	if _, err := repos.markets.ByPoolAddress(ctx, testChainID, missing); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestMarketListReturnsEmptySliceNotNil(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	markets, err := repos.markets.List(ctx, 888888)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if markets == nil {
		t.Fatal("expected an empty slice rather than nil")
	}

	if len(markets) != 0 {
		t.Fatalf("expected no markets on an unused chain, got %d", len(markets))
	}
}

func TestSnapshotInsertAndLatest(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	market := newMarket(t, tx)
	base := time.Now().UTC().Truncate(time.Second)

	older := newSnapshot(market.ID, base.Add(-2*time.Minute), 10)
	newer := newSnapshot(market.ID, base.Add(-1*time.Minute), 20)

	if err := repos.snapshots.Insert(ctx, &older); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := repos.snapshots.Insert(ctx, &newer); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if older.ID < 1 || newer.ID < 1 {
		t.Fatal("expected inserted snapshots to receive ids")
	}

	latest, err := repos.snapshots.Latest(ctx, market.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if latest.BlockNumber != 20 {
		t.Fatalf("expected the newest snapshot, got block %d", latest.BlockNumber)
	}

	if latest.TotalSupplied.String() != "150000000000" {
		t.Fatalf("expected the numeric column to round trip, got %s", latest.TotalSupplied.String())
	}
}

func TestSnapshotLatestNotFound(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	market := newMarket(t, tx)

	if _, err := repos.snapshots.Latest(ctx, market.ID); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected ErrNotFound for a market with no snapshots, got %v", err)
	}
}

func TestSnapshotSinceFiltersAndOrders(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	market := newMarket(t, tx)
	base := time.Now().UTC().Truncate(time.Second)

	for index := range 5 {
		snapshot := newSnapshot(market.ID, base.Add(time.Duration(index-4)*time.Minute), int64(index+1))

		if err := repos.snapshots.Insert(ctx, &snapshot); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	recent, err := repos.snapshots.Since(ctx, market.ID, base.Add(-2*time.Minute), 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(recent) != 3 {
		t.Fatalf("expected three snapshots within the window, got %d", len(recent))
	}

	for index := 1; index < len(recent); index++ {
		if recent[index].CapturedAt.Before(recent[index-1].CapturedAt) {
			t.Fatal("expected snapshots to be ordered oldest first")
		}
	}
}

func TestSnapshotSinceRespectsTheLimit(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	market := newMarket(t, tx)
	base := time.Now().UTC().Truncate(time.Second)

	for index := range 6 {
		snapshot := newSnapshot(market.ID, base.Add(time.Duration(index-6)*time.Minute), int64(index+1))

		if err := repos.snapshots.Insert(ctx, &snapshot); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	limited, err := repos.snapshots.Since(ctx, market.ID, base.Add(-24*time.Hour), 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(limited) != 2 {
		t.Fatalf("expected the limit to apply, got %d", len(limited))
	}
}

func TestSnapshotSinceReturnsEmptySliceNotNil(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	market := newMarket(t, tx)

	snapshots, err := repos.snapshots.Since(ctx, market.ID, time.Now().UTC().Add(time.Hour), 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if snapshots == nil {
		t.Fatal("expected an empty slice rather than nil")
	}

	if len(snapshots) != 0 {
		t.Fatalf("expected no snapshots in the future window, got %d", len(snapshots))
	}
}

func TestSnapshotRejectsBadInput(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	if err := repos.snapshots.Insert(ctx, nil); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	orphan := newSnapshot(0, time.Now().UTC(), 1)
	if err := repos.snapshots.Insert(ctx, &orphan); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput for a missing market id, got %v", err)
	}

	if _, err := repos.snapshots.Latest(ctx, 0); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}

	if _, err := repos.snapshots.Since(ctx, -1, time.Now().UTC(), 10); !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestSnapshotRejectsAnUnknownMarket(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	orphan := newSnapshot(999_999_999, time.Now().UTC(), 1)

	if err := repos.snapshots.Insert(ctx, &orphan); err == nil {
		t.Fatal("expected the foreign key to reject an unknown market")
	}
}

func TestSnapshotSinceCapsAnAbsurdLimit(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	market := newMarket(t, tx)
	snapshot := newSnapshot(market.ID, time.Now().UTC().Add(-time.Minute), 1)

	if err := repos.snapshots.Insert(ctx, &snapshot); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	found, err := repos.snapshots.Since(ctx, market.ID, time.Now().UTC().Add(-time.Hour), 1_000_000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(found) != 1 {
		t.Fatalf("expected the single snapshot, got %d", len(found))
	}
}

func TestTransactionIsolationBetweenTests(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	before, err := repos.markets.List(ctx, 777777)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(before) != 0 {
		t.Fatalf("expected an unused chain to stay empty across tests, found %d markets", len(before))
	}
}

func lower(value string) string {
	out := []byte(value)

	for index, char := range out {
		if char >= 'A' && char <= 'Z' {
			out[index] = char + 32
		}
	}

	return string(out)
}

func upper(value string) string {
	if len(value) < 2 {
		return value
	}

	body := []byte(value[2:])

	for index, char := range body {
		if char >= 'a' && char <= 'z' {
			body[index] = char - 32
		}
	}

	return value[:2] + string(body)
}

func TestAssetUpsertReportsADuplicateSymbol(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	first := newAsset(t, tx, "WETH", 18)

	clash := domain.Asset{
		ChainID:   testChainID,
		Address:   addressFrom(nextUnique()),
		Symbol:    first.Symbol,
		Name:      "same symbol, different address",
		Decimals:  18,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := repos.assets.Upsert(ctx, &clash)
	if !errors.Is(err, domain.ErrAlreadyExists) {
		t.Fatalf("expected ErrAlreadyExists for a duplicate symbol, got %v", err)
	}
}

func TestRepositoriesSurfaceDatabaseFailures(t *testing.T) {
	tx := newTx(t)
	repos := repositoryFor(tx)
	ctx := context.Background()

	market := newMarket(t, tx)
	asset := newAsset(t, tx, "WETH", 18)
	snapshot := newSnapshot(market.ID, time.Now().UTC(), 1)

	tx.Rollback()

	cases := []struct {
		name string
		call func() error
	}{
		{name: "asset by id", call: func() error {
			_, err := repos.assets.ByID(ctx, asset.ID)

			return err
		}},
		{name: "asset by address", call: func() error {
			_, err := repos.assets.ByAddress(ctx, testChainID, asset.Address)

			return err
		}},
		{name: "asset list", call: func() error {
			_, err := repos.assets.List(ctx, testChainID)

			return err
		}},
		{name: "asset upsert", call: func() error {
			clone := asset
			clone.ID = 0

			return repos.assets.Upsert(ctx, &clone)
		}},
		{name: "market by id", call: func() error {
			_, err := repos.markets.ByID(ctx, market.ID)

			return err
		}},
		{name: "market by pool address", call: func() error {
			_, err := repos.markets.ByPoolAddress(ctx, testChainID, market.PoolAddress)

			return err
		}},
		{name: "market list", call: func() error {
			_, err := repos.markets.List(ctx, testChainID)

			return err
		}},
		{name: "market upsert", call: func() error {
			clone := market
			clone.ID = 0

			return repos.markets.Upsert(ctx, &clone)
		}},
		{name: "snapshot insert", call: func() error {
			return repos.snapshots.Insert(ctx, &snapshot)
		}},
		{name: "snapshot latest", call: func() error {
			_, err := repos.snapshots.Latest(ctx, market.ID)

			return err
		}},
		{name: "snapshot since", call: func() error {
			_, err := repos.snapshots.Since(ctx, market.ID, time.Now().UTC().Add(-time.Hour), 10)

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
