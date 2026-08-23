package queryparam_test

import (
	"errors"
	"net/url"
	"testing"
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/queryparam"
)

func valuesFrom(t *testing.T, rawQuery string) url.Values {
	t.Helper()

	values, err := url.ParseQuery(rawQuery)
	if err != nil {
		t.Fatalf("could not build the test query %q: %v", rawQuery, err)
	}

	return values
}

func TestStringTrimsAndDefaultsToEmpty(t *testing.T) {
	values := valuesFrom(t, "cursor=%20abc%20&blank=%20%20")

	if got := queryparam.String(values, "cursor"); got != "abc" {
		t.Fatalf("expected the value to be trimmed, got %q", got)
	}

	if got := queryparam.String(values, "blank"); got != "" {
		t.Fatalf("expected a whitespace only value to read as empty, got %q", got)
	}

	if got := queryparam.String(values, "absent"); got != "" {
		t.Fatalf("expected an absent key to read as empty, got %q", got)
	}
}

func TestIntReadsAValueOrFallsBack(t *testing.T) {
	values := valuesFrom(t, "limit=50&spaced=%2012%20&empty=")

	cases := []struct {
		key  string
		want int
	}{
		{key: "limit", want: 50},
		{key: "spaced", want: 12},
		{key: "empty", want: 25},
		{key: "absent", want: 25},
	}

	for _, testCase := range cases {
		t.Run(testCase.key, func(t *testing.T) {
			got, err := queryparam.Int(values, testCase.key, 25)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != testCase.want {
				t.Fatalf("expected %d, got %d", testCase.want, got)
			}
		})
	}
}

func TestIntAcceptsZeroAndNegatives(t *testing.T) {
	values := valuesFrom(t, "zero=0&negative=-5")

	zero, err := queryparam.Int(values, "zero", 25)
	if err != nil || zero != 0 {
		t.Fatalf("expected zero to be read as a value rather than absent, got %d and %v", zero, err)
	}

	negative, err := queryparam.Int(values, "negative", 25)
	if err != nil || negative != -5 {
		t.Fatalf("expected the parser to pass a negative through for the caller to judge, got %d and %v", negative, err)
	}
}

func TestIntRejectsNonNumbers(t *testing.T) {
	for _, raw := range []string{"limit=many", "limit=12.5", "limit=1e3", "limit=%2B", "limit=12abc"} {
		t.Run(raw, func(t *testing.T) {
			_, err := queryparam.Int(valuesFrom(t, raw), "limit", 25)

			if !errors.Is(err, queryparam.ErrInvalid) {
				t.Fatalf("expected ErrInvalid for %q, got %v", raw, err)
			}

			var paramError *queryparam.ParamError
			if !errors.As(err, &paramError) || paramError.Param != "limit" {
				t.Fatalf("expected the error to name the parameter, got %v", err)
			}
		})
	}
}

func TestTimeReadsRFC3339AsUTC(t *testing.T) {
	values := valuesFrom(t, "from=2026-08-20T03:20:11Z&offset=2026-08-20T05:20:11%2B02:00&empty=")

	from, err := queryparam.Time(values, "from")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if from == nil || !from.Equal(time.Date(2026, 8, 20, 3, 20, 11, 0, time.UTC)) {
		t.Fatalf("unexpected timestamp %v", from)
	}

	offset, err := queryparam.Time(values, "offset")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if offset == nil || !offset.Equal(time.Date(2026, 8, 20, 3, 20, 11, 0, time.UTC)) {
		t.Fatalf("expected an offset timestamp to be converted to utc, got %v", offset)
	}

	if offset.Location() != time.UTC {
		t.Fatalf("expected the returned time to carry utc, got %s", offset.Location())
	}
}

func TestTimeReturnsNilWhenAbsent(t *testing.T) {
	values := valuesFrom(t, "empty=&spaces=%20%20")

	for _, key := range []string{"empty", "spaces", "absent"} {
		got, err := queryparam.Time(values, key)
		if err != nil {
			t.Fatalf("unexpected error for %q: %v", key, err)
		}

		if got != nil {
			t.Fatalf("expected nil for %q, got %v", key, got)
		}
	}
}

func TestTimeRejectsOtherFormats(t *testing.T) {
	for _, raw := range []string{"from=2026-08-20", "from=20/08/2026", "from=today", "from=1755660011"} {
		t.Run(raw, func(t *testing.T) {
			_, err := queryparam.Time(valuesFrom(t, raw), "from")

			if !errors.Is(err, queryparam.ErrInvalid) {
				t.Fatalf("expected ErrInvalid for %q, got %v", raw, err)
			}
		})
	}
}

func TestListSplitsOnCommasAndRepeats(t *testing.T) {
	values := valuesFrom(t, "kind=borrow,repay&kind=deposit")

	got := queryparam.List(values, "kind")
	want := []string{"borrow", "repay", "deposit"}

	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}

	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("expected %v in order, got %v", want, got)
		}
	}
}

func TestListTrimsDropsBlanksAndDeduplicates(t *testing.T) {
	values := valuesFrom(t, "kind=%20borrow%20,,repay,borrow&kind=%20&kind=repay")

	got := queryparam.List(values, "kind")
	want := []string{"borrow", "repay"}

	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}

	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("expected %v, got %v", want, got)
		}
	}
}

func TestListReturnsEmptySliceNotNil(t *testing.T) {
	got := queryparam.List(valuesFrom(t, "other=1"), "kind")

	if got == nil {
		t.Fatal("expected an empty slice so callers can range without a nil check")
	}

	if len(got) != 0 {
		t.Fatalf("expected no items, got %v", got)
	}
}

func TestParamErrorMessageNamesTheParameterAndReason(t *testing.T) {
	err := &queryparam.ParamError{Param: "limit", Reason: "must be a whole number"}

	if got := err.Error(); got != "query parameter is not valid: limit: must be a whole number" {
		t.Fatalf("unexpected message %q", got)
	}

	if !errors.Is(err, queryparam.ErrInvalid) {
		t.Fatal("expected the sentinel to be reachable through errors.Is")
	}
}
