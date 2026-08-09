package bigmath

import (
	"database/sql/driver"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type Int struct {
	value big.Int
}

func Zero() Int {
	return Int{}
}

func FromBig(value *big.Int) Int {
	if value == nil {
		return Int{}
	}

	var result Int
	result.value.Set(value)

	return result
}

func FromInt64(value int64) Int {
	var result Int
	result.value.SetInt64(value)

	return result
}

func FromString(raw string) (Int, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return Int{}, fmt.Errorf("bigmath: empty value")
	}

	var result Int
	if _, ok := result.value.SetString(trimmed, 10); !ok {
		return Int{}, fmt.Errorf("bigmath: %q is not a base-10 integer", raw)
	}

	return result, nil
}

func MustFromString(raw string) Int {
	value, err := FromString(raw)
	if err != nil {
		panic(err)
	}

	return value
}

func (i Int) Big() *big.Int {
	return new(big.Int).Set(&i.value)
}

func (i Int) String() string {
	return i.value.String()
}

func (i Int) IsZero() bool {
	return i.value.Sign() == 0
}

func (i Int) Sign() int {
	return i.value.Sign()
}

func (i Int) Cmp(other Int) int {
	return i.value.Cmp(&other.value)
}

func (i Int) Add(other Int) Int {
	var result Int
	result.value.Add(&i.value, &other.value)

	return result
}

func (i Int) Sub(other Int) Int {
	var result Int
	result.value.Sub(&i.value, &other.value)

	return result
}

func (i Int) Mul(other Int) Int {
	var result Int
	result.value.Mul(&i.value, &other.value)

	return result
}

func (i Int) Div(other Int) (Int, error) {
	if other.IsZero() {
		return Int{}, fmt.Errorf("bigmath: division by zero")
	}

	var result Int
	result.value.Quo(&i.value, &other.value)

	return result, nil
}

func (i *Int) Scan(src any) error {
	if src == nil {
		i.value.SetInt64(0)

		return nil
	}

	switch typed := src.(type) {
	case int64:
		i.value.SetInt64(typed)

		return nil
	case string:
		return i.scanString(typed)
	case []byte:
		return i.scanString(string(typed))
	case float64:
		return fmt.Errorf("bigmath: refusing to scan float64 %v, precision would be lost", typed)
	default:
		return fmt.Errorf("bigmath: cannot scan %T into Int", src)
	}
}

func (i *Int) scanString(raw string) error {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		i.value.SetInt64(0)

		return nil
	}

	if _, ok := i.value.SetString(trimmed, 10); !ok {
		return fmt.Errorf("bigmath: cannot scan %q into Int", raw)
	}

	return nil
}

func (i Int) Value() (driver.Value, error) {
	return i.value.String(), nil
}

func (i Int) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(i.value.String())), nil
}

func (i *Int) UnmarshalJSON(data []byte) error {
	raw := string(data)

	if raw == "null" {
		i.value.SetInt64(0)

		return nil
	}

	if unquoted, err := strconv.Unquote(raw); err == nil {
		raw = unquoted
	}

	return i.scanString(raw)
}

func (Int) GormDataType() string {
	return "numeric(78,0)"
}
