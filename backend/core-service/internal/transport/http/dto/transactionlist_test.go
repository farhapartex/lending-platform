package dto_test

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/domain"
	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
	"github.com/farhapartex/lending-platform/core-service/pkg/bigmath"
	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
	"github.com/farhapartex/lending-platform/core-service/pkg/queryparam"
)

var listMoment = time.Date(2026, 8, 22, 9, 30, 0, 0, time.UTC)

func maskAsHex(id int64) (string, error) {
	return "txn_" + strings.Repeat("a", 4) + string(rune('0'+id%10)), nil
}

func failingMask(int64) (string, error) {
	return "", errors.New("masking failed")
}

func pageWith(count int, next cursor.Key, block *int64) domain.TransactionPage {
	items := make([]domain.UserTransaction, 0, count)

	for index := 0; index < count; index++ {
		items = append(items, domain.UserTransaction{
			ID:          int64(index + 1),
			Kind:        domain.TransactionKindBorrow,
			Amount:      bigmath.FromInt64(1_000),
			BlockNumber: int64(500 + index),
			BlockTime:   listMoment,
			TxHash:      "0xabc",
			Asset:       &domain.Asset{Symbol: "USDC", Decimals: 6},
		})
	}

	return domain.TransactionPage{
		Items:      items,
		NextCursor: next,
		AsOf:       domain.IndexedAt{Block: block, Time: listMoment},
	}
}

func TestTransactionListResponseMasksEveryItem(t *testing.T) {
	response, err := dto.NewTransactionListResponse(pageWith(3, cursor.Key{}, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(response.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(response.Items))
	}

	for _, item := range response.Items {
		if !strings.HasPrefix(item.ID, "txn_") {
			t.Fatalf("expected a masked id, got %q", item.ID)
		}
	}
}

func TestTransactionListResponseCarriesTheItemShape(t *testing.T) {
	response, err := dto.NewTransactionListResponse(pageWith(1, cursor.Key{}, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	item := response.Items[0]

	if item.Kind != "borrow" || item.Amount.Amount != "1000" || item.Amount.Symbol != "USDC" {
		t.Fatalf("unexpected item %+v", item)
	}

	if item.Status != dto.TransactionStatusConfirmed {
		t.Fatalf("expected a confirmed status, got %q", item.Status)
	}
}

func TestTransactionListResponseEncodesTheCursorOnlyWhenPresent(t *testing.T) {
	withoutCursor, err := dto.NewTransactionListResponse(pageWith(1, cursor.Key{}, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if withoutCursor.NextCursor != nil {
		t.Fatalf("expected a null cursor, got %q", *withoutCursor.NextCursor)
	}

	key := cursor.Key{Time: listMoment, ID: 42}

	withCursor, err := dto.NewTransactionListResponse(pageWith(1, key, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if withCursor.NextCursor == nil {
		t.Fatal("expected a cursor")
	}

	decoded, err := cursor.Decode(*withCursor.NextCursor)
	if err != nil {
		t.Fatalf("expected the cursor to round trip, got %v", err)
	}

	if decoded.ID != 42 || !decoded.Time.Equal(listMoment) {
		t.Fatalf("expected the cursor to round trip, got %+v", decoded)
	}
}

func TestTransactionListResponseReportsAsOf(t *testing.T) {
	block := int64(4_218)

	response, err := dto.NewTransactionListResponse(pageWith(1, cursor.Key{}, &block), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response.AsOf.Block == nil || *response.AsOf.Block != 4_218 {
		t.Fatalf("expected block 4218, got %v", response.AsOf.Block)
	}

	if response.AsOf.Time != "2026-08-22T09:30:00Z" {
		t.Fatalf("unexpected time %q", response.AsOf.Time)
	}
}

func TestTransactionListResponseLeavesTheBlockNullWhenUnknown(t *testing.T) {
	response, err := dto.NewTransactionListResponse(pageWith(1, cursor.Key{}, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response.AsOf.Block != nil {
		t.Fatalf("expected a null block rather than a made up one, got %d", *response.AsOf.Block)
	}
}

func TestTransactionListResponseSerialisesAnEmptyPageAsAnArray(t *testing.T) {
	response, err := dto.NewTransactionListResponse(pageWith(0, cursor.Key{}, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	encoded, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("could not encode: %v", err)
	}

	if !strings.Contains(string(encoded), `"items":[]`) {
		t.Fatalf("expected an empty array rather than null, got %s", encoded)
	}

	if !strings.Contains(string(encoded), `"next_cursor":null`) {
		t.Fatalf("expected a null cursor, got %s", encoded)
	}
}

func TestTransactionListResponseSurfacesAMaskingFailure(t *testing.T) {
	_, err := dto.NewTransactionListResponse(pageWith(2, cursor.Key{}, nil), failingMask)

	if err == nil {
		t.Fatal("expected a masking failure to surface rather than emit a blank id")
	}
}

func TestParseTransactionListRequestReadsEveryParameter(t *testing.T) {
	key := cursor.Key{Time: listMoment, ID: 42}

	values := url.Values{}
	values.Set(dto.ParamKind, "borrow,repay")
	values.Set(dto.ParamFrom, "2026-08-01T00:00:00Z")
	values.Set(dto.ParamTo, "2026-08-22T00:00:00Z")
	values.Set(dto.ParamCursor, cursor.Encode(key))
	values.Set(dto.ParamLimit, "50")

	request, err := dto.ParseTransactionListRequest("0xabc", values)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if request.Address != "0xabc" {
		t.Fatalf("expected the address to pass through, got %q", request.Address)
	}

	if len(request.Kinds) != 2 || request.Kinds[0] != domain.TransactionKindBorrow {
		t.Fatalf("unexpected kinds %v", request.Kinds)
	}

	if request.From == nil || request.To == nil {
		t.Fatalf("expected both bounds, got %v and %v", request.From, request.To)
	}

	if request.After.ID != 42 {
		t.Fatalf("expected the cursor to be decoded, got %+v", request.After)
	}

	if request.Limit != 50 {
		t.Fatalf("expected limit 50, got %d", request.Limit)
	}
}

func TestParseTransactionListRequestDefaultsToAnUnfilteredPage(t *testing.T) {
	request, err := dto.ParseTransactionListRequest("0xabc", url.Values{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(request.Kinds) != 0 {
		t.Fatalf("expected no kind filter, got %v", request.Kinds)
	}

	if request.From != nil || request.To != nil {
		t.Fatal("expected no time window")
	}

	if !request.After.IsZero() {
		t.Fatal("expected no cursor")
	}

	if request.Limit != 0 {
		t.Fatalf("expected the service to choose the page size, got %d", request.Limit)
	}
}

func TestParseTransactionListRequestAcceptsRepeatedKinds(t *testing.T) {
	values := url.Values{dto.ParamKind: []string{"borrow", "repay,deposit"}}

	request, err := dto.ParseTransactionListRequest("0xabc", values)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(request.Kinds) != 3 {
		t.Fatalf("expected three kinds, got %v", request.Kinds)
	}
}

func TestParseTransactionListRequestRejectsBadInput(t *testing.T) {
	cases := map[string]url.Values{
		"unknown kind":    {dto.ParamKind: []string{"teleport"}},
		"camel case kind": {dto.ParamKind: []string{"collateralAdded"}},
		"bad from":        {dto.ParamFrom: []string{"2026-08-01"}},
		"bad to":          {dto.ParamTo: []string{"yesterday"}},
		"bad limit":       {dto.ParamLimit: []string{"many"}},
		"negative limit":  {dto.ParamLimit: []string{"-5"}},
		"bad cursor":      {dto.ParamCursor: []string{"not-a-cursor"}},
		"forged cursor":   {dto.ParamCursor: []string{"dGFtcGVyZWQ"}},
	}

	for name, values := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := dto.ParseTransactionListRequest("0xabc", values)

			if !errors.Is(err, queryparam.ErrInvalid) {
				t.Fatalf("expected ErrInvalid, got %v", err)
			}

			var paramError *queryparam.ParamError
			if !errors.As(err, &paramError) || paramError.Param == "" {
				t.Fatalf("expected the error to name the offending parameter, got %v", err)
			}
		})
	}
}

func TestParseTransactionListRequestAcceptsAZeroLimit(t *testing.T) {
	request, err := dto.ParseTransactionListRequest("0xabc", url.Values{dto.ParamLimit: []string{"0"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if request.Limit != 0 {
		t.Fatalf("expected zero to fall through to the default page size, got %d", request.Limit)
	}
}

func TestParseTransactionListRequestIgnoresBlankValues(t *testing.T) {
	values := url.Values{
		dto.ParamKind:   []string{"", "  "},
		dto.ParamFrom:   []string{""},
		dto.ParamTo:     []string{"  "},
		dto.ParamCursor: []string{"   "},
		dto.ParamLimit:  []string{""},
	}

	request, err := dto.ParseTransactionListRequest("0xabc", values)
	if err != nil {
		t.Fatalf("a blank parameter should behave as absent, got %v", err)
	}

	if len(request.Kinds) != 0 || request.From != nil || request.To != nil || !request.After.IsZero() {
		t.Fatalf("expected an unfiltered request, got %+v", request)
	}
}

func TestActivityResponseCarriesItemsAndAsOf(t *testing.T) {
	block := int64(4_218)

	response, err := dto.NewActivityResponse(pageWith(2, cursor.Key{}, &block), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(response.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(response.Items))
	}

	if response.AsOf.Block == nil || *response.AsOf.Block != 4_218 {
		t.Fatalf("expected block 4218, got %v", response.AsOf.Block)
	}

	if response.AsOf.Time != "2026-08-22T09:30:00Z" {
		t.Fatalf("unexpected time %q", response.AsOf.Time)
	}
}

func TestActivityResponseHasNoCursorField(t *testing.T) {
	response, err := dto.NewActivityResponse(pageWith(1, cursor.Key{Time: listMoment, ID: 42}, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	encoded, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("could not encode: %v", err)
	}

	if strings.Contains(string(encoded), "next_cursor") {
		t.Fatalf("the dashboard list does not page, so no cursor should be emitted even if the page carries one, got %s", encoded)
	}
}

func TestActivityResponseSerialisesAnEmptyListAsAnArray(t *testing.T) {
	response, err := dto.NewActivityResponse(pageWith(0, cursor.Key{}, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	encoded, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("could not encode: %v", err)
	}

	if !strings.Contains(string(encoded), `"items":[]`) {
		t.Fatalf("expected an empty array rather than null, got %s", encoded)
	}
}

func TestActivityResponseMasksEveryItem(t *testing.T) {
	response, err := dto.NewActivityResponse(pageWith(3, cursor.Key{}, nil), maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, item := range response.Items {
		if !strings.HasPrefix(item.ID, "txn_") {
			t.Fatalf("expected a masked id, got %q", item.ID)
		}
	}
}

func TestActivityResponseSurfacesAMaskingFailure(t *testing.T) {
	if _, err := dto.NewActivityResponse(pageWith(2, cursor.Key{}, nil), failingMask); err == nil {
		t.Fatal("expected a masking failure to surface rather than emit a blank id")
	}
}

func TestActivityAndListShareTheItemShape(t *testing.T) {
	page := pageWith(1, cursor.Key{}, nil)

	list, err := dto.NewTransactionListResponse(page, maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	activity, err := dto.NewActivityResponse(page, maskAsHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if list.Items[0] != activity.Items[0] {
		t.Fatalf("one client type must cover both, got %+v and %+v", list.Items[0], activity.Items[0])
	}
}

func TestParseLimitReadsAndValidates(t *testing.T) {
	cases := map[string]struct {
		values url.Values
		want   int
	}{
		"absent": {values: url.Values{}, want: 0},
		"blank":  {values: url.Values{dto.ParamLimit: []string{"  "}}, want: 0},
		"zero":   {values: url.Values{dto.ParamLimit: []string{"0"}}, want: 0},
		"value":  {values: url.Values{dto.ParamLimit: []string{"12"}}, want: 12},
		"spaced": {values: url.Values{dto.ParamLimit: []string{" 7 "}}, want: 7},
		"large":  {values: url.Values{dto.ParamLimit: []string{"9999"}}, want: 9999},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := dto.ParseLimit(testCase.values)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != testCase.want {
				t.Fatalf("expected %d, got %d", testCase.want, got)
			}
		})
	}
}

func TestParseLimitRejectsBadValues(t *testing.T) {
	for _, raw := range []string{"many", "-1", "-999", "1.5"} {
		t.Run(raw, func(t *testing.T) {
			_, err := dto.ParseLimit(url.Values{dto.ParamLimit: []string{raw}})

			if !errors.Is(err, queryparam.ErrInvalid) {
				t.Fatalf("expected ErrInvalid for %q, got %v", raw, err)
			}

			var paramError *queryparam.ParamError
			if !errors.As(err, &paramError) || paramError.Param != dto.ParamLimit {
				t.Fatalf("expected the error to name the limit parameter, got %v", err)
			}
		})
	}
}
