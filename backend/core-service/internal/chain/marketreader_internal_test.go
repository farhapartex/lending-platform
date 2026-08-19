package chain

import (
	"errors"
	"math/big"
	"testing"

	"github.com/farhapartex/lending-platform/core-service/internal/chain/bindings"
	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

func sampleMarketData() bindings.MarketData {
	return bindings.MarketData{
		TotalSupplied:           big.NewInt(150_000_000_000),
		TotalBorrowed:           big.NewInt(15_000_000_000),
		AvailableLiquidity:      big.NewInt(135_000_000_000),
		UtilizationBps:          big.NewInt(1000),
		SupplyRatePerSecond:     big.NewInt(51_869_291),
		BorrowRatePerSecond:     big.NewInt(576_325_468),
		SupplyAprBps:            big.NewInt(16),
		BorrowAprBps:            big.NewInt(181),
		SupplyIndex:             big.NewInt(1_000_000_000_000_000_000),
		BorrowIndex:             big.NewInt(1_000_000_000_000_000_000),
		MaxLtvBps:               big.NewInt(7500),
		LiquidationThresholdBps: big.NewInt(8000),
		LiquidationBonusBps:     big.NewInt(500),
		KinkUtilizationBps:      big.NewInt(8000),
		ReserveFactorBps:        big.NewInt(1000),
		MinDeposit:              big.NewInt(1_000_000),
		AccruedReserves:         big.NewInt(0),
		DepositsPaused:          false,
		BorrowPaused:            false,
	}
}

func TestMarketViewFromCopiesEveryField(t *testing.T) {
	view, err := marketViewFrom(sampleMarketData())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if view.TotalSupplied.String() != "150000000000" {
		t.Fatalf("unexpected total supplied %s", view.TotalSupplied)
	}

	if view.UtilizationBps != 1000 || view.SupplyAprBps != 16 || view.BorrowAprBps != 181 {
		t.Fatalf("unexpected rate fields %+v", view)
	}

	if view.MaxLtvBps != 7500 || view.LiquidationThresholdBps != 8000 || view.LiquidationBonusBps != 500 {
		t.Fatalf("unexpected risk fields %+v", view)
	}

	if view.KinkUtilizationBps != 8000 || view.ReserveFactorBps != 1000 {
		t.Fatalf("unexpected curve fields %+v", view)
	}

	if view.MinDeposit.String() != "1000000" || view.AccruedReserves.String() != "0" {
		t.Fatalf("unexpected pool fields %+v", view)
	}

	if view.DepositsPaused || view.BorrowPaused {
		t.Fatal("expected both pause flags to be false")
	}
}

func TestMarketViewFromCarriesPauseFlags(t *testing.T) {
	raw := sampleMarketData()
	raw.DepositsPaused = true
	raw.BorrowPaused = true

	view, err := marketViewFrom(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !view.DepositsPaused || !view.BorrowPaused {
		t.Fatal("expected both pause flags to be true")
	}
}

func TestMarketViewFromRejectsBpsBeyondInt32(t *testing.T) {
	beyond := new(big.Int).Lsh(big.NewInt(1), 40)

	fields := map[string]func(*bindings.MarketData){
		"utilizationBps":          func(m *bindings.MarketData) { m.UtilizationBps = beyond },
		"supplyAprBps":            func(m *bindings.MarketData) { m.SupplyAprBps = beyond },
		"borrowAprBps":            func(m *bindings.MarketData) { m.BorrowAprBps = beyond },
		"maxLtvBps":               func(m *bindings.MarketData) { m.MaxLtvBps = beyond },
		"liquidationThresholdBps": func(m *bindings.MarketData) { m.LiquidationThresholdBps = beyond },
		"liquidationBonusBps":     func(m *bindings.MarketData) { m.LiquidationBonusBps = beyond },
		"kinkUtilizationBps":      func(m *bindings.MarketData) { m.KinkUtilizationBps = beyond },
		"reserveFactorBps":        func(m *bindings.MarketData) { m.ReserveFactorBps = beyond },
	}

	for name, corrupt := range fields {
		t.Run(name, func(t *testing.T) {
			raw := sampleMarketData()
			corrupt(&raw)

			_, err := marketViewFrom(raw)
			if !errors.Is(err, domain.ErrInvalidInput) {
				t.Fatalf("expected ErrInvalidInput, got %v", err)
			}

			if !contains(err.Error(), name) {
				t.Fatalf("expected the error to name %s, got %v", name, err)
			}
		})
	}
}

func TestMarketViewFromRejectsNilBps(t *testing.T) {
	raw := sampleMarketData()
	raw.UtilizationBps = nil

	_, err := marketViewFrom(raw)
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput for a nil bps field, got %v", err)
	}
}

func TestNarrowerReportsOnlyTheFirstFailure(t *testing.T) {
	narrower := &bpsNarrower{}
	beyond := new(big.Int).Lsh(big.NewInt(1), 40)

	narrower.take("first", beyond)
	firstError := narrower.err

	if firstError == nil {
		t.Fatal("expected the first failure to be recorded")
	}

	result := narrower.take("second", big.NewInt(7500))

	if result != 0 {
		t.Fatalf("expected later reads to short circuit to zero, got %d", result)
	}

	if narrower.err.Error() != firstError.Error() {
		t.Fatal("expected the first failure to be preserved")
	}
}

func contains(haystack, needle string) bool {
	if len(needle) == 0 {
		return true
	}

	for index := 0; index+len(needle) <= len(haystack); index++ {
		if haystack[index:index+len(needle)] == needle {
			return true
		}
	}

	return false
}
