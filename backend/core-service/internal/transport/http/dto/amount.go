package dto

import (
	"math/big"

	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
)

type Amount struct {
	Amount   string `json:"amount"`
	Decimals int    `json:"decimals"`
	Symbol   string `json:"symbol"`
}

func NewAmount(value bigmath.Int, decimals int16, symbol string) Amount {
	return Amount{
		Amount:   value.String(),
		Decimals: int(decimals),
		Symbol:   symbol,
	}
}

func NewAmountFromBig(value *big.Int, decimals int16, symbol string) Amount {
	if value == nil {
		return Amount{Amount: "0", Decimals: int(decimals), Symbol: symbol}
	}

	return Amount{
		Amount:   value.String(),
		Decimals: int(decimals),
		Symbol:   symbol,
	}
}

func ZeroAmount(decimals int16, symbol string) Amount {
	return Amount{Amount: "0", Decimals: int(decimals), Symbol: symbol}
}

func ScaledValue(value bigmath.Int) string {
	return value.String()
}

func ScaledValueFromBig(value *big.Int) string {
	if value == nil {
		return "0"
	}

	return value.String()
}
