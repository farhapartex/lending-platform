package chain_test

import (
	"context"
	"errors"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/farhapartex/lending-platform/core-service/internal/chain"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

const anvilLensAddress = "0x610178dA211FEF7D417bC0e6FeD39F05609AD788"

func chainRPCURL() string {
	if url := os.Getenv("TEST_CHAIN_RPC_URL"); url != "" {
		return url
	}

	return "http://localhost:8545"
}

func lensAddress() string {
	if address := os.Getenv("TEST_LENS_ADDRESS"); address != "" {
		return address
	}

	return anvilLensAddress
}

func liveReader(t *testing.T) *chain.MarketReader {
	t.Helper()

	client, err := ethclient.Dial(chainRPCURL())
	if err != nil {
		t.Skipf("chain node is not reachable at %s: %v", chainRPCURL(), err)
	}

	t.Cleanup(client.Close)

	pingCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if _, err := client.ChainID(pingCtx); err != nil {
		t.Skipf("chain node is not reachable at %s: %v", chainRPCURL(), err)
	}

	reader, err := chain.NewMarketReader(client, chain.MarketReaderParams{
		LensAddress:    lensAddress(),
		RequestTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	return reader
}

func TestNewMarketReaderRejectsMissingClient(t *testing.T) {
	_, err := chain.NewMarketReader(nil, chain.MarketReaderParams{LensAddress: anvilLensAddress})

	if !errors.Is(err, domain.ErrChainUnreachable) {
		t.Fatalf("expected ErrChainUnreachable, got %v", err)
	}
}

func TestNewMarketReaderRejectsBadAddress(t *testing.T) {
	client, err := ethclient.Dial(chainRPCURL())
	if err != nil {
		t.Skipf("chain node is not reachable: %v", err)
	}

	t.Cleanup(client.Close)

	for _, address := range []string{"", "not-an-address", "0x1234"} {
		_, err := chain.NewMarketReader(client, chain.MarketReaderParams{LensAddress: address})

		if !errors.Is(err, domain.ErrInvalidInput) {
			t.Fatalf("expected ErrInvalidInput for %q, got %v", address, err)
		}
	}
}

func TestNewMarketReaderUsesADefaultTimeout(t *testing.T) {
	client, err := ethclient.Dial(chainRPCURL())
	if err != nil {
		t.Skipf("chain node is not reachable: %v", err)
	}

	t.Cleanup(client.Close)

	reader, err := chain.NewMarketReader(client, chain.MarketReaderParams{LensAddress: lensAddress()})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if reader == nil {
		t.Fatal("expected a reader")
	}
}

func TestReadMarketViewAgainstALiveChain(t *testing.T) {
	reader := liveReader(t)

	view, err := reader.ReadMarketView(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if view.TotalSupplied == nil || view.TotalBorrowed == nil || view.AvailableLiquidity == nil {
		t.Fatal("expected the amount fields to be populated")
	}

	if view.TotalSupplied.Sign() < 0 || view.TotalBorrowed.Sign() < 0 {
		t.Fatalf("expected non negative totals, got %s and %s", view.TotalSupplied, view.TotalBorrowed)
	}

	expectedLiquidity := new(big.Int).Sub(view.TotalSupplied, view.TotalBorrowed)
	if expectedLiquidity.Sign() < 0 {
		expectedLiquidity.SetInt64(0)
	}

	if view.AvailableLiquidity.Cmp(expectedLiquidity) != 0 {
		t.Fatalf(
			"expected liquidity %s to equal supplied minus borrowed %s",
			view.AvailableLiquidity, expectedLiquidity,
		)
	}
}

func TestReadMarketViewReturnsSeededRiskParameters(t *testing.T) {
	reader := liveReader(t)

	view, err := reader.ReadMarketView(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if view.MaxLtvBps != 7500 {
		t.Fatalf("expected a max ltv of 7500 bps, got %d", view.MaxLtvBps)
	}

	if view.LiquidationThresholdBps != 8000 {
		t.Fatalf("expected a liquidation threshold of 8000 bps, got %d", view.LiquidationThresholdBps)
	}

	if view.LiquidationBonusBps != 500 {
		t.Fatalf("expected a liquidation bonus of 500 bps, got %d", view.LiquidationBonusBps)
	}

	if view.KinkUtilizationBps != 8000 {
		t.Fatalf("expected a kink of 8000 bps, got %d", view.KinkUtilizationBps)
	}

	if view.ReserveFactorBps != 1000 {
		t.Fatalf("expected a reserve factor of 1000 bps, got %d", view.ReserveFactorBps)
	}
}

func TestReadMarketViewKeepsTheContractInvariants(t *testing.T) {
	reader := liveReader(t)

	view, err := reader.ReadMarketView(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if view.MaxLtvBps >= view.LiquidationThresholdBps {
		t.Fatalf(
			"expected max ltv %d to stay below the liquidation threshold %d",
			view.MaxLtvBps, view.LiquidationThresholdBps,
		)
	}

	if view.UtilizationBps < 0 || view.UtilizationBps > 10_000 {
		t.Fatalf("expected utilization within 0 to 10000 bps, got %d", view.UtilizationBps)
	}

	if view.BorrowAprBps < view.SupplyAprBps {
		t.Fatalf(
			"expected the borrow apr %d to be at least the supply apr %d",
			view.BorrowAprBps, view.SupplyAprBps,
		)
	}

	if view.BorrowIndex.Cmp(view.SupplyIndex) < 0 {
		t.Fatalf("expected the borrow index to lead the supply index")
	}
}

func TestReadMarketViewFailsAgainstAnEmptyAddress(t *testing.T) {
	client, err := ethclient.Dial(chainRPCURL())
	if err != nil {
		t.Skipf("chain node is not reachable: %v", err)
	}

	t.Cleanup(client.Close)

	reader, err := chain.NewMarketReader(client, chain.MarketReaderParams{
		LensAddress:    "0x000000000000000000000000000000000000dEaD",
		RequestTimeout: 3 * time.Second,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := reader.ReadMarketView(context.Background()); !errors.Is(err, domain.ErrChainUnreachable) {
		t.Fatalf("expected ErrChainUnreachable when no contract is deployed, got %v", err)
	}
}

func TestReadMarketViewRespectsACancelledContext(t *testing.T) {
	reader := liveReader(t)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := reader.ReadMarketView(ctx); err == nil {
		t.Fatal("expected an error for a cancelled context")
	}
}
