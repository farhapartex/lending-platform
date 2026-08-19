package bigmath_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

const uint256Max = "115792089237316195423570985008687907853269984665640564039457584007913129639935"

func TestZeroAndIsZero(t *testing.T) {
	if !bigmath.Zero().IsZero() {
		t.Fatal("expected Zero to be zero")
	}

	if bigmath.FromInt64(1).IsZero() {
		t.Fatal("expected one not to be zero")
	}

	if bigmath.Zero().String() != "0" {
		t.Fatalf("expected \"0\", got %q", bigmath.Zero().String())
	}
}

func TestFromBigCopiesRatherThanAliases(t *testing.T) {
	source := big.NewInt(100)
	wrapped := bigmath.FromBig(source)

	source.SetInt64(999)

	if wrapped.String() != "100" {
		t.Fatalf("expected the wrapped value to be independent, got %s", wrapped.String())
	}
}

func TestFromBigHandlesNil(t *testing.T) {
	if !bigmath.FromBig(nil).IsZero() {
		t.Fatal("expected a nil input to become zero")
	}
}

func TestFromInt64(t *testing.T) {
	cases := []int64{0, 1, -1, 6_900_000_000, -6_900_000_000}

	for _, input := range cases {
		if got := bigmath.FromInt64(input).String(); got != big.NewInt(input).String() {
			t.Fatalf("expected %d, got %s", input, got)
		}
	}
}

func TestFromString(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "zero", input: "0", want: "0"},
		{name: "typical amount", input: "50132158115", want: "50132158115"},
		{name: "uint256 max", input: uint256Max, want: uint256Max},
		{name: "negative", input: "-42", want: "-42"},
		{name: "surrounding whitespace", input: "  1234  ", want: "1234"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			value, err := bigmath.FromString(testCase.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if value.String() != testCase.want {
				t.Fatalf("expected %s, got %s", testCase.want, value.String())
			}
		})
	}
}

func TestFromStringRejectsBadInput(t *testing.T) {
	for _, input := range []string{"", "   ", "abc", "1.5", "0x10", "1e18", "12 34"} {
		if _, err := bigmath.FromString(input); err == nil {
			t.Fatalf("expected an error for %q", input)
		}
	}
}

func TestMustFromString(t *testing.T) {
	if bigmath.MustFromString("42").String() != "42" {
		t.Fatal("expected 42")
	}

	defer func() {
		if recover() == nil {
			t.Fatal("expected MustFromString to panic on bad input")
		}
	}()

	bigmath.MustFromString("not a number")
}

func TestBigReturnsACopy(t *testing.T) {
	value := bigmath.FromInt64(500)

	extracted := value.Big()
	extracted.SetInt64(999)

	if value.String() != "500" {
		t.Fatalf("expected the original to be untouched, got %s", value.String())
	}
}

func TestSignAndCmp(t *testing.T) {
	positive := bigmath.FromInt64(10)
	negative := bigmath.FromInt64(-10)
	zero := bigmath.Zero()

	if positive.Sign() != 1 || negative.Sign() != -1 || zero.Sign() != 0 {
		t.Fatal("unexpected signs")
	}

	if positive.Cmp(negative) != 1 {
		t.Fatal("expected 10 to be greater than -10")
	}

	if negative.Cmp(positive) != -1 {
		t.Fatal("expected -10 to be less than 10")
	}

	if positive.Cmp(bigmath.FromInt64(10)) != 0 {
		t.Fatal("expected equal values to compare equal")
	}
}

func TestArithmetic(t *testing.T) {
	left := bigmath.MustFromString("150000000000")
	right := bigmath.MustFromString("15000000000")

	if got := left.Add(right).String(); got != "165000000000" {
		t.Fatalf("unexpected sum %s", got)
	}

	if got := left.Sub(right).String(); got != "135000000000" {
		t.Fatalf("unexpected difference %s", got)
	}

	if got := bigmath.FromInt64(6).Mul(bigmath.FromInt64(7)).String(); got != "42" {
		t.Fatalf("unexpected product %s", got)
	}
}

func TestArithmeticDoesNotMutateOperands(t *testing.T) {
	left := bigmath.FromInt64(10)
	right := bigmath.FromInt64(3)

	left.Add(right)
	left.Sub(right)
	left.Mul(right)

	if left.String() != "10" || right.String() != "3" {
		t.Fatalf("expected operands to be untouched, got %s and %s", left.String(), right.String())
	}
}

func TestArithmeticSurvivesUint256Scale(t *testing.T) {
	huge := bigmath.MustFromString(uint256Max)

	if got := huge.Sub(huge).String(); got != "0" {
		t.Fatalf("expected zero, got %s", got)
	}

	if got := huge.Add(bigmath.FromInt64(1)).String(); got == uint256Max {
		t.Fatal("expected the sum to exceed uint256 max rather than wrap")
	}
}

func TestDiv(t *testing.T) {
	quotient, err := bigmath.FromInt64(42).Div(bigmath.FromInt64(7))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if quotient.String() != "6" {
		t.Fatalf("expected 6, got %s", quotient.String())
	}

	truncated, err := bigmath.FromInt64(7).Div(bigmath.FromInt64(2))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if truncated.String() != "3" {
		t.Fatalf("expected integer division to truncate to 3, got %s", truncated.String())
	}
}

func TestDivByZero(t *testing.T) {
	if _, err := bigmath.FromInt64(1).Div(bigmath.Zero()); err == nil {
		t.Fatal("expected division by zero to error")
	}
}

func TestScan(t *testing.T) {
	cases := []struct {
		name  string
		input any
		want  string
	}{
		{name: "nil becomes zero", input: nil, want: "0"},
		{name: "int64", input: int64(6_900_000_000), want: "6900000000"},
		{name: "negative int64", input: int64(-5), want: "-5"},
		{name: "string", input: "50132158115", want: "50132158115"},
		{name: "byte slice", input: []byte("135000000000"), want: "135000000000"},
		{name: "uint256 max as string", input: uint256Max, want: uint256Max},
		{name: "uint256 max as bytes", input: []byte(uint256Max), want: uint256Max},
		{name: "empty string becomes zero", input: "", want: "0"},
		{name: "blank bytes become zero", input: []byte("   "), want: "0"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			var value bigmath.Int

			if err := value.Scan(testCase.input); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if value.String() != testCase.want {
				t.Fatalf("expected %s, got %s", testCase.want, value.String())
			}
		})
	}
}

func TestScanRefusesFloat(t *testing.T) {
	var value bigmath.Int

	err := value.Scan(float64(1.5))
	if err == nil {
		t.Fatal("expected scanning a float to be refused")
	}
}

func TestScanRejectsUnsupportedTypes(t *testing.T) {
	for _, input := range []any{true, struct{}{}, 12, []string{"a"}} {
		var value bigmath.Int

		if err := value.Scan(input); err == nil {
			t.Fatalf("expected an error scanning %T", input)
		}
	}
}

func TestScanRejectsNonNumericStrings(t *testing.T) {
	for _, input := range []any{"abc", []byte("1.5"), "0x10"} {
		var value bigmath.Int

		if err := value.Scan(input); err == nil {
			t.Fatalf("expected an error scanning %v", input)
		}
	}
}

func TestScanOverwritesAnyPreviousValue(t *testing.T) {
	value := bigmath.FromInt64(999)

	if err := value.Scan("7"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if value.String() != "7" {
		t.Fatalf("expected the previous value to be replaced, got %s", value.String())
	}
}

func TestValueRoundTripsThroughScan(t *testing.T) {
	original := bigmath.MustFromString(uint256Max)

	driverValue, err := original.Value()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	asString, ok := driverValue.(string)
	if !ok {
		t.Fatalf("expected the driver value to be a string, got %T", driverValue)
	}

	var restored bigmath.Int
	if err := restored.Scan(asString); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if restored.Cmp(original) != 0 {
		t.Fatalf("expected a lossless round trip, got %s", restored.String())
	}
}

func TestMarshalJSONEmitsAString(t *testing.T) {
	encoded, err := json.Marshal(bigmath.MustFromString("50132158115"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(encoded) != `"50132158115"` {
		t.Fatalf("expected a quoted string, got %s", encoded)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "quoted string", input: `"50132158115"`, want: "50132158115"},
		{name: "bare number", input: `12345`, want: "12345"},
		{name: "null becomes zero", input: `null`, want: "0"},
		{name: "uint256 max", input: `"` + uint256Max + `"`, want: uint256Max},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			var value bigmath.Int

			if err := json.Unmarshal([]byte(testCase.input), &value); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if value.String() != testCase.want {
				t.Fatalf("expected %s, got %s", testCase.want, value.String())
			}
		})
	}
}

func TestUnmarshalJSONRejectsGarbage(t *testing.T) {
	for _, input := range []string{`"abc"`, `{}`, `"1.5"`} {
		var value bigmath.Int

		if err := json.Unmarshal([]byte(input), &value); err == nil {
			t.Fatalf("expected an error unmarshalling %s", input)
		}
	}
}

func TestJSONRoundTripInsideAStruct(t *testing.T) {
	type payload struct {
		Supplied bigmath.Int `json:"supplied"`
	}

	original := payload{Supplied: bigmath.MustFromString(uint256Max)}

	encoded, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var restored payload
	if err := json.Unmarshal(encoded, &restored); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if restored.Supplied.Cmp(original.Supplied) != 0 {
		t.Fatalf("expected a lossless round trip, got %s", restored.Supplied.String())
	}
}

func TestGormDataType(t *testing.T) {
	if got := (bigmath.Int{}).GormDataType(); got != "numeric(78,0)" {
		t.Fatalf("expected numeric(78,0), got %q", got)
	}
}
