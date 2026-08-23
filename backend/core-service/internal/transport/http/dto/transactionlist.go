package dto

import (
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
)

type AsOfResponse struct {
	Block *int64 `json:"block"`
	Time  string `json:"time"`
}

type TransactionListResponse struct {
	Items      []TransactionResponse `json:"items"`
	NextCursor *string               `json:"next_cursor"`
	AsOf       AsOfResponse          `json:"as_of"`
}

type PublicIDFunc func(id int64) (string, error)

func NewTransactionListResponse(page domain.TransactionPage, publicID PublicIDFunc) (TransactionListResponse, error) {
	items := make([]TransactionResponse, 0, len(page.Items))

	for _, transaction := range page.Items {
		masked, err := publicID(transaction.ID)
		if err != nil {
			return TransactionListResponse{}, err
		}

		items = append(items, NewTransactionResponse(transaction, masked))
	}

	return TransactionListResponse{
		Items:      items,
		NextCursor: encodeNextCursor(page.NextCursor),
		AsOf:       newAsOfResponse(page.AsOf),
	}, nil
}

func encodeNextCursor(key cursor.Key) *string {
	if key.IsZero() {
		return nil
	}

	encoded := cursor.Encode(key)

	return &encoded
}

func newAsOfResponse(asOf domain.IndexedAt) AsOfResponse {
	return AsOfResponse{
		Block: asOf.Block,
		Time:  asOf.Time.UTC().Format(time.RFC3339),
	}
}
