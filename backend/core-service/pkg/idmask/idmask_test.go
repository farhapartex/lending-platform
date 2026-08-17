package idmask_test

import (
	"errors"
	"strconv"
	"strings"
	"testing"

	"github.com/farhapartex/lending-platform/core-service/pkg/idmask"
)

const secret = "local-development-id-mask-secret"

func newMasker(t *testing.T) *idmask.Masker {
	t.Helper()

	masker, err := idmask.New(secret)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	return masker
}

func TestNewRejectsEmptySecret(t *testing.T) {
	for _, input := range []string{"", "   ", "\t\n"} {
		if _, err := idmask.New(input); !errors.Is(err, idmask.ErrEmptySecret) {
			t.Fatalf("expected ErrEmptySecret for %q, got %v", input, err)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	masker := newMasker(t)

	ids := []int64{1, 2, 3, 42, 1000, 999999, 1 << 20, 1 << 40, 1<<62 - 1}
	kinds := []idmask.Kind{
		idmask.KindUser,
		idmask.KindMarket,
		idmask.KindPosition,
		idmask.KindTransaction,
		idmask.KindEvent,
		idmask.KindLiquidation,
		idmask.KindSnapshot,
		idmask.KindNotification,
	}

	for _, kind := range kinds {
		for _, id := range ids {
			t.Run(string(kind)+"/"+strconv.FormatInt(id, 10), func(t *testing.T) {
				token, err := masker.Mask(kind, id)
				if err != nil {
					t.Fatalf("unexpected mask error: %v", err)
				}

				back, err := masker.Unmask(kind, token)
				if err != nil {
					t.Fatalf("unexpected unmask error: %v", err)
				}

				if back != id {
					t.Fatalf("expected %d, got %d", id, back)
				}
			})
		}
	}
}

func TestRoundTripAcrossManySequentialIDs(t *testing.T) {
	masker := newMasker(t)

	seen := make(map[string]int64, 2000)

	for id := int64(1); id <= 2000; id++ {
		token, err := masker.Mask(idmask.KindTransaction, id)
		if err != nil {
			t.Fatalf("unexpected error at id %d: %v", id, err)
		}

		if previous, clash := seen[token]; clash {
			t.Fatalf("token %q produced by both %d and %d", token, previous, id)
		}
		seen[token] = id

		back, err := masker.Unmask(idmask.KindTransaction, token)
		if err != nil || back != id {
			t.Fatalf("round trip failed for %d: got %d with error %v", id, back, err)
		}
	}
}

func TestMaskIsDeterministic(t *testing.T) {
	masker := newMasker(t)

	first, err := masker.Mask(idmask.KindMarket, 7)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	second, err := masker.Mask(idmask.KindMarket, 7)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if first != second {
		t.Fatalf("expected a stable token, got %q then %q", first, second)
	}
}

func TestMaskRejectsBadInput(t *testing.T) {
	masker := newMasker(t)

	cases := []struct {
		name    string
		kind    idmask.Kind
		id      int64
		wantErr error
	}{
		{name: "zero id", kind: idmask.KindMarket, id: 0, wantErr: idmask.ErrIDOutOfRange},
		{name: "negative id", kind: idmask.KindMarket, id: -1, wantErr: idmask.ErrIDOutOfRange},
		{name: "empty kind", kind: "", id: 1, wantErr: idmask.ErrEmptyKind},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if _, err := masker.Mask(testCase.kind, testCase.id); !errors.Is(err, testCase.wantErr) {
				t.Fatalf("expected %v, got %v", testCase.wantErr, err)
			}
		})
	}
}

func TestTokenHasKindPrefixAndHidesTheID(t *testing.T) {
	masker := newMasker(t)

	token, err := masker.Mask(idmask.KindLiquidation, 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.HasPrefix(token, string(idmask.KindLiquidation)+"_") {
		t.Fatalf("expected a liq_ prefix, got %q", token)
	}

	if strings.Contains(token, "12345") {
		t.Fatalf("expected the raw id to be hidden, got %q", token)
	}
}

func TestSequentialIDsDoNotProduceSequentialTokens(t *testing.T) {
	masker := newMasker(t)

	first, err := masker.Mask(idmask.KindMarket, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	second, err := masker.Mask(idmask.KindMarket, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	firstBody := strings.TrimPrefix(first, string(idmask.KindMarket)+"_")
	secondBody := strings.TrimPrefix(second, string(idmask.KindMarket)+"_")

	shared := 0
	for index := 0; index < len(firstBody) && index < len(secondBody); index++ {
		if firstBody[index] == secondBody[index] {
			shared++
		}
	}

	if shared > len(firstBody)/2 {
		t.Fatalf("tokens for 1 and 2 look too similar: %q and %q", first, second)
	}
}

func TestSameIDDiffersByKind(t *testing.T) {
	masker := newMasker(t)

	marketToken, err := masker.Mask(idmask.KindMarket, 99)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	transactionToken, err := masker.Mask(idmask.KindTransaction, 99)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	marketBody := strings.TrimPrefix(marketToken, string(idmask.KindMarket)+"_")
	transactionBody := strings.TrimPrefix(transactionToken, string(idmask.KindTransaction)+"_")

	if marketBody == transactionBody {
		t.Fatal("expected different kinds to produce different token bodies")
	}
}

func TestUnmaskRejectsWrongKind(t *testing.T) {
	masker := newMasker(t)

	token, err := masker.Mask(idmask.KindMarket, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := masker.Unmask(idmask.KindTransaction, token); !errors.Is(err, idmask.ErrWrongKind) {
		t.Fatalf("expected ErrWrongKind, got %v", err)
	}
}

func TestUnmaskRejectsMalformedTokens(t *testing.T) {
	masker := newMasker(t)

	valid, err := masker.Mask(idmask.KindMarket, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	body := strings.TrimPrefix(valid, "mkt_")

	cases := []struct {
		name  string
		token string
	}{
		{name: "empty", token: ""},
		{name: "whitespace", token: "   "},
		{name: "no separator", token: "mkt" + body},
		{name: "empty body", token: "mkt_"},
		{name: "not base32", token: "mkt_!!!!!!!!!!!!!!!!!!!!"},
		{name: "body too short", token: "mkt_" + body[:len(body)-4]},
		{name: "body too long", token: "mkt_" + body + "aaaaaaaa"},
		{name: "tampered body", token: "mkt_" + flipFirstRune(body)},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := masker.Unmask(idmask.KindMarket, testCase.token)

			if !errors.Is(err, idmask.ErrInvalidToken) {
				t.Fatalf("expected ErrInvalidToken, got %v", err)
			}
		})
	}
}

func TestUnmaskRejectsEmptyKind(t *testing.T) {
	masker := newMasker(t)

	if _, err := masker.Unmask("", "mkt_aaaaaaaaaaaaaaaaaaaa"); !errors.Is(err, idmask.ErrEmptyKind) {
		t.Fatalf("expected ErrEmptyKind, got %v", err)
	}
}

func TestUnmaskIsCaseInsensitiveAndTrims(t *testing.T) {
	masker := newMasker(t)

	token, err := masker.Mask(idmask.KindEvent, 314)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	prefix, body, _ := strings.Cut(token, "_")
	upper := prefix + "_" + strings.ToUpper(body)

	back, err := masker.Unmask(idmask.KindEvent, "  "+upper+"  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if back != 314 {
		t.Fatalf("expected 314, got %d", back)
	}
}

func TestTokensFromADifferentSecretAreRejected(t *testing.T) {
	masker := newMasker(t)

	other, err := idmask.New("a completely different secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	token, err := other.Mask(idmask.KindMarket, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := masker.Unmask(idmask.KindMarket, token); !errors.Is(err, idmask.ErrInvalidToken) {
		t.Fatalf("expected a foreign token to be rejected, got %v", err)
	}
}

func TestDifferentSecretsProduceDifferentTokens(t *testing.T) {
	first := newMasker(t)

	second, err := idmask.New("another secret entirely")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	firstToken, err := first.Mask(idmask.KindMarket, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	secondToken, err := second.Mask(idmask.KindMarket, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if firstToken == secondToken {
		t.Fatal("expected different secrets to produce different tokens")
	}
}

func TestGuessedTokensAreOverwhelminglyRejected(t *testing.T) {
	masker := newMasker(t)

	accepted := 0

	for attempt := range 5000 {
		token, err := masker.Mask(idmask.KindMarket, int64(attempt+1))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		guess := "mkt_" + flipFirstRune(strings.TrimPrefix(token, "mkt_"))

		if _, err := masker.Unmask(idmask.KindMarket, guess); err == nil {
			accepted++
		}
	}

	if accepted > 0 {
		t.Fatalf("expected tampered tokens to be rejected, %d were accepted", accepted)
	}
}

func TestMustMask(t *testing.T) {
	masker := newMasker(t)

	token := masker.MustMask(idmask.KindMarket, 1)
	if token == "" {
		t.Fatal("expected a token")
	}

	defer func() {
		if recover() == nil {
			t.Fatal("expected MustMask to panic on an invalid id")
		}
	}()

	masker.MustMask(idmask.KindMarket, 0)
}

func flipFirstRune(body string) string {
	if body == "" {
		return body
	}

	replacement := byte('a')
	if body[0] == 'a' {
		replacement = 'b'
	}

	return string(replacement) + body[1:]
}
