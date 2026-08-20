package dto

import (
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
)

const TransactionStatusConfirmed = "confirmed"

type TransactionResponse struct {
	ID                   string `json:"id"`
	Kind                 string `json:"kind"`
	Amount               Amount `json:"amount"`
	HealthFactorAfterBps *int32 `json:"health_factor_after_bps"`
	TxHash               string `json:"tx_hash"`
	Block                int64  `json:"block"`
	BlockTime            string `json:"block_time"`
	LogIndex             int32  `json:"log_index"`
	Status               string `json:"status"`
}

func NewTransactionResponse(transaction domain.UserTransaction, publicID string) TransactionResponse {
	decimals, symbol := assetUnits(transaction.Asset)

	return TransactionResponse{
		ID:                   publicID,
		Kind:                 string(transaction.Kind),
		Amount:               NewAmount(transaction.Amount, decimals, symbol),
		HealthFactorAfterBps: transaction.HealthFactorAfterBps,
		TxHash:               transaction.TxHash,
		Block:                transaction.BlockNumber,
		BlockTime:            transaction.BlockTime.UTC().Format(time.RFC3339),
		LogIndex:             transaction.LogIndex,
		Status:               TransactionStatusConfirmed,
	}
}

func assetUnits(asset *domain.Asset) (int16, string) {
	if asset == nil {
		return 0, ""
	}

	return asset.Decimals, asset.Symbol
}
