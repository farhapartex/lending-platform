package cursor_test

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/farhapartex/lending-platform/core-service/pkg/cursor"
)

func TestRoundTrip(t *testing.T) {
	cases := []struct {
		name string
		key  cursor.Key
	}{
		{name: "typical", key: cursor.Key{Time: time.Date(2026, 8, 19, 3, 25, 46, 0, time.UTC), ID: 42}},
		{name: "zero id", key: cursor.Key{Time: time.Date(2026, 8, 19, 3, 25, 46, 0, time.UTC), ID: 0}},
		{name: "unix epoch", key: cursor.Key{Time: time.Unix(0, 0).UTC(), ID: 1}},
		{name: "microsecond precision", key: cursor.Key{Time: time.UnixMicro(1786966713123456).UTC(), ID: 7}},
		{name: "large id", key: cursor.Key{Time: time.Now().UTC().Truncate(time.Microsecond), ID: 1 << 40}},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			encoded := cursor.Encode(testCase.key)

			decoded, err := cursor.Decode(encoded)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !decoded.Time.Equal(testCase.key.Time.Truncate(time.Microsecond)) {
				t.Fatalf("expected %s, got %s", testCase.key.Time, decoded.Time)
			}

			if decoded.ID != testCase.key.ID {
				t.Fatalf("expected id %d, got %d", testCase.key.ID, decoded.ID)
			}
		})
	}
}

func TestEncodeNormalisesToUTC(t *testing.T) {
	zone := time.FixedZone("UTC+6", 6*60*60)
	local := time.Date(2026, 8, 19, 9, 25, 46, 0, zone)

	decoded, err := cursor.Decode(cursor.Encode(cursor.Key{Time: local, ID: 1}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if decoded.Time.Location() != time.UTC {
		t.Fatalf("expected UTC, got %s", decoded.Time.Location())
	}

	if !decoded.Time.Equal(local) {
		t.Fatalf("expected the same instant, got %s", decoded.Time)
	}
}

func TestEncodeIsOpaque(t *testing.T) {
	encoded := cursor.Encode(cursor.Key{Time: time.Date(2026, 8, 19, 3, 25, 46, 0, time.UTC), ID: 42})

	if strings.Contains(encoded, "42") {
		t.Fatalf("expected the id to be hidden, got %q", encoded)
	}

	if strings.Contains(encoded, "2026") {
		t.Fatalf("expected the timestamp to be hidden, got %q", encoded)
	}

	if strings.ContainsAny(encoded, "+/=") {
		t.Fatalf("expected a url safe encoding, got %q", encoded)
	}
}

func TestEncodeIsStable(t *testing.T) {
	key := cursor.Key{Time: time.Date(2026, 8, 19, 3, 25, 46, 0, time.UTC), ID: 42}

	if cursor.Encode(key) != cursor.Encode(key) {
		t.Fatal("expected encoding to be deterministic")
	}
}

func TestDecodeRejectsBadInput(t *testing.T) {
	valid := cursor.Encode(cursor.Key{Time: time.Now().UTC(), ID: 5})

	cases := []struct {
		name    string
		input   string
		wantErr error
	}{
		{name: "empty", input: "", wantErr: cursor.ErrEmpty},
		{name: "whitespace", input: "   ", wantErr: cursor.ErrEmpty},
		{name: "not base64", input: "!!!not-base64!!!", wantErr: cursor.ErrMalformed},
		{name: "base64 but not a cursor", input: "aGVsbG8", wantErr: cursor.ErrMalformed},
		{name: "too few parts", input: encodeRaw("v1.123"), wantErr: cursor.ErrMalformed},
		{name: "too many parts", input: encodeRaw("v1.123.4.5"), wantErr: cursor.ErrMalformed},
		{name: "wrong version", input: encodeRaw("v2.123.4"), wantErr: cursor.ErrMalformed},
		{name: "time not a number", input: encodeRaw("v1.soon.4"), wantErr: cursor.ErrMalformed},
		{name: "id not a number", input: encodeRaw("v1.123.last"), wantErr: cursor.ErrMalformed},
		{name: "negative id", input: encodeRaw("v1.123.-4"), wantErr: cursor.ErrMalformed},
		{name: "truncated valid cursor", input: valid[:len(valid)-3], wantErr: cursor.ErrMalformed},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			key, err := cursor.Decode(testCase.input)

			if !errors.Is(err, testCase.wantErr) {
				t.Fatalf("expected %v, got %v", testCase.wantErr, err)
			}

			if !key.IsZero() {
				t.Fatalf("expected a zero key on error, got %+v", key)
			}
		})
	}
}

func TestDecodeToleratesSurroundingWhitespace(t *testing.T) {
	encoded := cursor.Encode(cursor.Key{Time: time.Now().UTC(), ID: 9})

	decoded, err := cursor.Decode("  " + encoded + "\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if decoded.ID != 9 {
		t.Fatalf("expected id 9, got %d", decoded.ID)
	}
}

func TestDecodeOptional(t *testing.T) {
	key, err := cursor.DecodeOptional("")
	if err != nil {
		t.Fatalf("expected an empty cursor to be allowed, got %v", err)
	}

	if !key.IsZero() {
		t.Fatalf("expected a zero key, got %+v", key)
	}

	if _, err := cursor.DecodeOptional("   "); err != nil {
		t.Fatalf("expected whitespace to be treated as absent, got %v", err)
	}

	encoded := cursor.Encode(cursor.Key{Time: time.Now().UTC(), ID: 3})

	present, err := cursor.DecodeOptional(encoded)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if present.ID != 3 {
		t.Fatalf("expected id 3, got %d", present.ID)
	}

	if _, err := cursor.DecodeOptional("garbage!!"); !errors.Is(err, cursor.ErrMalformed) {
		t.Fatalf("expected ErrMalformed, got %v", err)
	}
}

func TestKeyIsZero(t *testing.T) {
	if !(cursor.Key{}).IsZero() {
		t.Fatal("expected an empty key to report zero")
	}

	if (cursor.Key{ID: 1}).IsZero() {
		t.Fatal("expected a key with an id not to report zero")
	}

	if (cursor.Key{Time: time.Now()}).IsZero() {
		t.Fatal("expected a key with a time not to report zero")
	}
}

func encodeRaw(payload string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(payload))
}
