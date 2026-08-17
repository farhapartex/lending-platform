package idmask

import (
	"encoding/binary"
	"errors"
	"strings"
	"testing"
)

func mintToken(t *testing.T, masker *Masker, kind Kind, raw uint64) string {
	t.Helper()

	permuted := masker.permute(kind, raw)

	payload := make([]byte, payloadSize)
	binary.BigEndian.PutUint64(payload[:8], permuted)
	copy(payload[8:], masker.tag(kind, payload[:8]))

	return string(kind) + separator + strings.ToLower(encoding.EncodeToString(payload))
}

func TestUnmaskRejectsOutOfRangeIDEvenWithAValidTag(t *testing.T) {
	masker, err := New("internal-test-secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cases := []struct {
		name string
		raw  uint64
	}{
		{name: "zero", raw: 0},
		{name: "just above the ceiling", raw: uint64(1)<<62 + 1},
		{name: "maximum uint64", raw: ^uint64(0)},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			token := mintToken(t, masker, KindMarket, testCase.raw)

			if _, err := masker.Unmask(KindMarket, token); !errors.Is(err, ErrInvalidToken) {
				t.Fatalf("expected ErrInvalidToken for raw value %d, got %v", testCase.raw, err)
			}
		})
	}
}

func TestPermuteAndReverseAreInverses(t *testing.T) {
	masker, err := New("internal-test-secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	values := []uint64{0, 1, 2, 255, 256, 1 << 31, 1<<32 - 1, 1 << 32, 1 << 61, ^uint64(0)}

	for _, kind := range []Kind{KindMarket, KindUser, KindEvent} {
		for _, value := range values {
			permuted := masker.permute(kind, value)

			if back := masker.reverse(kind, permuted); back != value {
				t.Fatalf("kind %s value %d permuted to %d and reversed to %d", kind, value, permuted, back)
			}
		}
	}
}

func TestPermuteSpreadsAdjacentValues(t *testing.T) {
	masker, err := New("internal-test-secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	first := masker.permute(KindMarket, 1)
	second := masker.permute(KindMarket, 2)

	difference := first ^ second
	bits := 0

	for difference > 0 {
		bits += int(difference & 1)
		difference >>= 1
	}

	if bits < 16 {
		t.Fatalf("expected adjacent ids to differ in many bits, only %d differed", bits)
	}
}

func TestRoundFunctionDependsOnKindAndRound(t *testing.T) {
	masker, err := New("internal-test-secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	base := masker.round(KindMarket, 0, 12345)

	if masker.round(KindUser, 0, 12345) == base {
		t.Fatal("expected the round function to depend on the kind")
	}

	if masker.round(KindMarket, 1, 12345) == base {
		t.Fatal("expected the round function to depend on the round number")
	}

	if masker.round(KindMarket, 0, 12346) == base {
		t.Fatal("expected the round function to depend on the input half")
	}
}
