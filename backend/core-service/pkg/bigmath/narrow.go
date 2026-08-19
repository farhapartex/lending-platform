package bigmath

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

var (
	ErrNilValue     = errors.New("value must not be nil")
	ErrOutOfRange   = errors.New("value does not fit the target type")
	ErrNegativeOnly = errors.New("value must not be negative")
)

func ToInt32(value *big.Int) (int32, error) {
	if value == nil {
		return 0, ErrNilValue
	}

	if !value.IsInt64() {
		return 0, fmt.Errorf("%w: %s exceeds int32", ErrOutOfRange, value.String())
	}

	asInt64 := value.Int64()

	if asInt64 > math.MaxInt32 || asInt64 < math.MinInt32 {
		return 0, fmt.Errorf("%w: %s exceeds int32", ErrOutOfRange, value.String())
	}

	return int32(asInt64), nil
}

func ToInt64(value *big.Int) (int64, error) {
	if value == nil {
		return 0, ErrNilValue
	}

	if !value.IsInt64() {
		return 0, fmt.Errorf("%w: %s exceeds int64", ErrOutOfRange, value.String())
	}

	return value.Int64(), nil
}

func ToUint64(value *big.Int) (uint64, error) {
	if value == nil {
		return 0, ErrNilValue
	}

	if value.Sign() < 0 {
		return 0, fmt.Errorf("%w: %s", ErrNegativeOnly, value.String())
	}

	if !value.IsUint64() {
		return 0, fmt.Errorf("%w: %s exceeds uint64", ErrOutOfRange, value.String())
	}

	return value.Uint64(), nil
}

func FromBigOrZero(value *big.Int) Int {
	if value == nil {
		return Zero()
	}

	return FromBig(value)
}
