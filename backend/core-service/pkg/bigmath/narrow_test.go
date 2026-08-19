package bigmath_test

import (
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

func TestToInt32(t *testing.T) {
	cases := []struct {
		name  string
		input *big.Int
		want  int32
	}{
		{name: "zero", input: big.NewInt(0), want: 0},
		{name: "typical bps", input: big.NewInt(7500), want: 7500},
		{name: "apr above kink", input: big.NewInt(20753), want: 20753},
		{name: "maximum", input: big.NewInt(math.MaxInt32), want: math.MaxInt32},
		{name: "minimum", input: big.NewInt(math.MinInt32), want: math.MinInt32},
		{name: "negative", input: big.NewInt(-1), want: -1},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := bigmath.ToInt32(testCase.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != testCase.want {
				t.Fatalf("expected %d, got %d", testCase.want, got)
			}
		})
	}
}

func TestToInt32RejectsOutOfRange(t *testing.T) {
	cases := []struct {
		name  string
		input *big.Int
	}{
		{name: "just above max", input: big.NewInt(int64(math.MaxInt32) + 1)},
		{name: "just below min", input: big.NewInt(int64(math.MinInt32) - 1)},
		{name: "beyond int64", input: new(big.Int).Lsh(big.NewInt(1), 100)},
		{name: "uint256 max", input: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := bigmath.ToInt32(testCase.input)

			if !errors.Is(err, bigmath.ErrOutOfRange) {
				t.Fatalf("expected ErrOutOfRange, got %v", err)
			}
		})
	}
}

func TestToInt32RejectsNil(t *testing.T) {
	if _, err := bigmath.ToInt32(nil); !errors.Is(err, bigmath.ErrNilValue) {
		t.Fatalf("expected ErrNilValue, got %v", err)
	}
}

func TestToInt64(t *testing.T) {
	value, err := bigmath.ToInt64(big.NewInt(math.MaxInt64))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if value != math.MaxInt64 {
		t.Fatalf("expected the maximum int64, got %d", value)
	}

	if _, err := bigmath.ToInt64(nil); !errors.Is(err, bigmath.ErrNilValue) {
		t.Fatalf("expected ErrNilValue, got %v", err)
	}

	beyond := new(big.Int).Lsh(big.NewInt(1), 200)
	if _, err := bigmath.ToInt64(beyond); !errors.Is(err, bigmath.ErrOutOfRange) {
		t.Fatalf("expected ErrOutOfRange, got %v", err)
	}
}

func TestToUint64(t *testing.T) {
	value, err := bigmath.ToUint64(new(big.Int).SetUint64(math.MaxUint64))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if value != math.MaxUint64 {
		t.Fatalf("expected the maximum uint64, got %d", value)
	}

	if _, err := bigmath.ToUint64(nil); !errors.Is(err, bigmath.ErrNilValue) {
		t.Fatalf("expected ErrNilValue, got %v", err)
	}

	if _, err := bigmath.ToUint64(big.NewInt(-1)); !errors.Is(err, bigmath.ErrNegativeOnly) {
		t.Fatalf("expected ErrNegativeOnly, got %v", err)
	}

	beyond := new(big.Int).Lsh(big.NewInt(1), 200)
	if _, err := bigmath.ToUint64(beyond); !errors.Is(err, bigmath.ErrOutOfRange) {
		t.Fatalf("expected ErrOutOfRange, got %v", err)
	}
}

func TestFromBigOrZero(t *testing.T) {
	if got := bigmath.FromBigOrZero(nil); !got.IsZero() {
		t.Fatalf("expected a nil value to become zero, got %s", got.String())
	}

	huge := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))

	if got := bigmath.FromBigOrZero(huge); got.String() != huge.String() {
		t.Fatalf("expected the value to survive, got %s", got.String())
	}
}
