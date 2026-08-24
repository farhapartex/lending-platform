package dto

import (
	"fmt"
	"net/url"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
	"github.com/farhapartex/lending-platform/core-service/pkg/queryparam"
)

const (
	ParamKind   = "kind"
	ParamFrom   = "from"
	ParamTo     = "to"
	ParamCursor = "cursor"
	ParamLimit  = "limit"
)

func ParseTransactionListRequest(address string, values url.Values) (domain.TransactionListRequest, error) {
	kinds, err := parseKinds(values)
	if err != nil {
		return domain.TransactionListRequest{}, err
	}

	from, err := queryparam.Time(values, ParamFrom)
	if err != nil {
		return domain.TransactionListRequest{}, err
	}

	to, err := queryparam.Time(values, ParamTo)
	if err != nil {
		return domain.TransactionListRequest{}, err
	}

	after, err := parseCursor(values)
	if err != nil {
		return domain.TransactionListRequest{}, err
	}

	limit, err := ParseLimit(values)
	if err != nil {
		return domain.TransactionListRequest{}, err
	}

	return domain.TransactionListRequest{
		Address: address,
		Kinds:   kinds,
		From:    from,
		To:      to,
		After:   after,
		Limit:   limit,
	}, nil
}

func ParseLimit(values url.Values) (int, error) {
	limit, err := queryparam.Int(values, ParamLimit, 0)
	if err != nil {
		return 0, err
	}

	if limit < 0 {
		return 0, &queryparam.ParamError{Param: ParamLimit, Reason: "must not be negative"}
	}

	return limit, nil
}

func parseKinds(values url.Values) ([]domain.TransactionKind, error) {
	raw := queryparam.List(values, ParamKind)

	kinds := make([]domain.TransactionKind, 0, len(raw))

	for _, item := range raw {
		kind, ok := domain.ParseTransactionKind(item)
		if !ok {
			return nil, &queryparam.ParamError{
				Param:  ParamKind,
				Reason: fmt.Sprintf("%q is not a transaction kind", item),
			}
		}

		kinds = append(kinds, kind)
	}

	return kinds, nil
}

func parseCursor(values url.Values) (cursor.Key, error) {
	key, err := cursor.DecodeOptional(queryparam.String(values, ParamCursor))
	if err != nil {
		return cursor.Key{}, &queryparam.ParamError{
			Param:  ParamCursor,
			Reason: "is not a cursor from a previous response",
		}
	}

	return key, nil
}
