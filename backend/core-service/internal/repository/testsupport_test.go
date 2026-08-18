package repository_test

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/platform/database"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

const testChainID int64 = 31337

var (
	sharedDB   *gorm.DB
	openOnce   sync.Once
	openErr    error
	uniqueLock sync.Mutex
	uniqueSeed int64
)

func databaseURL() string {
	if url := os.Getenv("TEST_DATABASE_URL"); url != "" {
		return url
	}

	return "postgres://lending:lending@localhost:5432/lending?sslmode=disable"
}

func openDatabase() (*gorm.DB, error) {
	openOnce.Do(func() {
		sharedDB, openErr = database.Open(database.Options{
			DSN:          databaseURL(),
			MaxOpenConns: 5,
			MaxIdleConns: 2,
			LogQueries:   false,
		})
	})

	return sharedDB, openErr
}

func newTx(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := openDatabase()
	if err != nil {
		t.Skipf("postgres is not reachable at %s: %v", databaseURL(), err)
	}

	pool, err := db.DB()
	if err != nil {
		t.Skipf("postgres pool is unavailable: %v", err)
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := pool.PingContext(pingCtx); err != nil {
		t.Skipf("postgres is not reachable at %s: %v", databaseURL(), err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("could not begin a transaction: %v", tx.Error)
	}

	t.Cleanup(func() {
		tx.Rollback()
	})

	return tx
}

func nextUnique() int64 {
	uniqueLock.Lock()
	defer uniqueLock.Unlock()

	uniqueSeed++

	return time.Now().UnixNano() + uniqueSeed
}

func addressFrom(seed int64) string {
	const hexDigits = "0123456789abcdef"

	body := make([]byte, 40)
	value := seed

	for index := len(body) - 1; index >= 0; index-- {
		body[index] = hexDigits[value&0xf]
		value >>= 4

		if value == 0 && index > 0 {
			value = seed >> 8
		}
	}

	return "0x" + string(body)
}

func newAsset(t *testing.T, tx *gorm.DB, symbol string, decimals int16) domain.Asset {
	t.Helper()

	seed := nextUnique()

	asset := domain.Asset{
		ChainID:      testChainID,
		Address:      addressFrom(seed),
		Symbol:       symbol + "-" + shortSuffix(seed),
		Name:         symbol + " test asset",
		Decimals:     decimals,
		IsCollateral: decimals == 18,
		IsBorrowable: decimals == 6,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	if err := repositoryFor(tx).assets.Upsert(context.Background(), &asset); err != nil {
		t.Fatalf("could not create the test asset: %v", err)
	}

	return asset
}

func newMarket(t *testing.T, tx *gorm.DB) domain.Market {
	t.Helper()

	collateral := newAsset(t, tx, "WETH", 18)
	debt := newAsset(t, tx, "USDC", 6)

	seed := nextUnique()
	now := time.Now().UTC()

	market := domain.Market{
		ChainID:                   testChainID,
		CollateralAssetID:         collateral.ID,
		DebtAssetID:               debt.ID,
		PoolAddress:               addressFrom(seed),
		CollateralVaultAddress:    addressFrom(seed + 1),
		ControllerAddress:         addressFrom(seed + 2),
		LiquidationManagerAddress: addressFrom(seed + 3),
		InterestRateModelAddress:  addressFrom(seed + 4),
		OracleAdapterAddress:      addressFrom(seed + 5),
		MaxLTVBps:                 7500,
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

	if err := repositoryFor(tx).markets.Upsert(context.Background(), &market); err != nil {
		t.Fatalf("could not create the test market: %v", err)
	}

	return market
}

func newSnapshot(marketID int64, capturedAt time.Time, block int64) domain.MarketSnapshot {
	return domain.MarketSnapshot{
		MarketID:           marketID,
		CapturedAt:         capturedAt,
		BlockNumber:        block,
		TotalSupplied:      bigmath.FromInt64(150_000_000_000),
		TotalBorrowed:      bigmath.FromInt64(15_000_000_000),
		AvailableLiquidity: bigmath.FromInt64(135_000_000_000),
		UtilizationBps:     1000,
		SupplyRateBps:      16,
		BorrowRateBps:      181,
		SupplyIndex:        bigmath.MustFromString("1000000000000000000"),
		BorrowIndex:        bigmath.MustFromString("1000000000000000000"),
		PositionsAtRisk:    0,
		CreatedAt:          time.Now().UTC(),
	}
}

func shortSuffix(seed int64) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"

	out := make([]byte, 6)
	value := seed

	for index := range out {
		out[index] = alphabet[value%int64(len(alphabet))]
		value /= int64(len(alphabet))
	}

	return string(out)
}
