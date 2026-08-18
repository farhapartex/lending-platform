package dto_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

func TestNewAmount(t *testing.T) {
	amount := dto.NewAmount(bigmath.FromInt64(50_132_158_115), 6, "USDC")

	if amount.Amount != "50132158115" {
		t.Fatalf("expected the raw base units as a string, got %q", amount.Amount)
	}

	if amount.Decimals != 6 {
		t.Fatalf("expected 6 decimals, got %d", amount.Decimals)
	}

	if amount.Symbol != "USDC" {
		t.Fatalf("expected USDC, got %q", amount.Symbol)
	}
}

func TestNewAmountKeepsFullPrecisionBeyondFloat64(t *testing.T) {
	huge := bigmath.MustFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935")

	amount := dto.NewAmount(huge, 18, "WETH")

	if amount.Amount != huge.String() {
		t.Fatalf("expected the value to survive intact, got %q", amount.Amount)
	}
}

func TestNewAmountZero(t *testing.T) {
	amount := dto.NewAmount(bigmath.Zero(), 18, "WETH")

	if amount.Amount != "0" {
		t.Fatalf("expected \"0\", got %q", amount.Amount)
	}
}

func TestNewAmountFromBig(t *testing.T) {
	amount := dto.NewAmountFromBig(big.NewInt(2_000_000_000_000_000_000), 18, "WETH")

	if amount.Amount != "2000000000000000000" {
		t.Fatalf("expected the raw value, got %q", amount.Amount)
	}

	if amount.Decimals != 18 {
		t.Fatalf("expected 18 decimals, got %d", amount.Decimals)
	}
}

func TestNewAmountFromBigHandlesNil(t *testing.T) {
	amount := dto.NewAmountFromBig(nil, 6, "USDC")

	if amount.Amount != "0" {
		t.Fatalf("expected a nil value to become \"0\", got %q", amount.Amount)
	}

	if amount.Decimals != 6 || amount.Symbol != "USDC" {
		t.Fatalf("expected decimals and symbol to survive, got %+v", amount)
	}
}

func TestZeroAmount(t *testing.T) {
	amount := dto.ZeroAmount(6, "USDC")

	if amount.Amount != "0" || amount.Decimals != 6 || amount.Symbol != "USDC" {
		t.Fatalf("unexpected zero amount %+v", amount)
	}
}

func TestAmountSerialisesAsAStringNotANumber(t *testing.T) {
	encoded, err := json.Marshal(dto.NewAmount(bigmath.FromInt64(50_132_158_115), 6, "USDC"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `{"amount":"50132158115","decimals":6,"symbol":"USDC"}`

	if string(encoded) != expected {
		t.Fatalf("expected %s, got %s", expected, encoded)
	}
}

func TestAmountSurvivesAJavaScriptStyleRoundTrip(t *testing.T) {
	huge := bigmath.MustFromString("9007199254740993")

	encoded, err := json.Marshal(dto.NewAmount(huge, 6, "USDC"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var decoded dto.Amount
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if decoded.Amount != "9007199254740993" {
		t.Fatalf("expected the value above 2^53 to survive, got %q", decoded.Amount)
	}
}

func TestScaledValue(t *testing.T) {
	if got := dto.ScaledValue(bigmath.FromInt64(580_000_000_000)); got != "580000000000" {
		t.Fatalf("expected 580000000000, got %q", got)
	}

	if got := dto.ScaledValue(bigmath.Zero()); got != "0" {
		t.Fatalf("expected 0, got %q", got)
	}
}

func TestScaledValueFromBig(t *testing.T) {
	if got := dto.ScaledValueFromBig(big.NewInt(1_092_025_600_000)); got != "1092025600000" {
		t.Fatalf("expected 1092025600000, got %q", got)
	}

	if got := dto.ScaledValueFromBig(nil); got != "0" {
		t.Fatalf("expected a nil value to become 0, got %q", got)
	}
}
