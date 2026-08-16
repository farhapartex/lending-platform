package ethaddr_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/farhapartex/lending-platform/core-service/pkg/ethaddr"
)

const (
	lowercase = "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"
	uppercase = "0xF39FD6E51AAD88F6F4CE6AB8827279CFFFB92266"
	checksum  = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
)

func TestNormalize(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{name: "already lowercase", input: lowercase, want: lowercase},
		{name: "uppercase becomes lowercase", input: uppercase, want: lowercase},
		{name: "checksum becomes lowercase", input: checksum, want: lowercase},
		{name: "surrounding whitespace trimmed", input: "  " + checksum + "\n", want: lowercase},
		{name: "zero address is a valid shape", input: ethaddr.ZeroAddress, want: ethaddr.ZeroAddress},
		{name: "empty", input: "", wantErr: ethaddr.ErrEmpty},
		{name: "only whitespace", input: "   ", wantErr: ethaddr.ErrEmpty},
		{name: "missing prefix", input: strings.TrimPrefix(lowercase, "0x"), wantErr: ethaddr.ErrMissingPrefix},
		{name: "wrong prefix case", input: "0X" + strings.TrimPrefix(lowercase, "0x"), wantErr: ethaddr.ErrMissingPrefix},
		{name: "one character short", input: lowercase[:41], wantErr: ethaddr.ErrInvalidLength},
		{name: "one character long", input: lowercase + "a", wantErr: ethaddr.ErrInvalidLength},
		{name: "prefix only", input: "0x", wantErr: ethaddr.ErrInvalidLength},
		{name: "non hex character", input: "0xg39fd6e51aad88f6f4ce6ab8827279cfffb92266", wantErr: ethaddr.ErrInvalidHex},
		{name: "punctuation inside", input: "0xf39fd6e51aad88f6f4ce6ab8827279cfffb9226-", wantErr: ethaddr.ErrInvalidHex},
		{name: "multibyte char makes it too long in bytes", input: "0xf39fd6e51aad88f6f4ce6ab8827279cfffb9226é", wantErr: ethaddr.ErrInvalidLength},
		{name: "multibyte char at the right byte length", input: "0x" + strings.Repeat("a", 38) + "é", wantErr: ethaddr.ErrInvalidHex},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ethaddr.Normalize(testCase.input)

			if testCase.wantErr != nil {
				if !errors.Is(err, testCase.wantErr) {
					t.Fatalf("expected error %v, got %v", testCase.wantErr, err)
				}

				if got != "" {
					t.Fatalf("expected empty result on error, got %q", got)
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != testCase.want {
				t.Fatalf("expected %q, got %q", testCase.want, got)
			}
		})
	}
}

func TestNormalizeIsIdempotent(t *testing.T) {
	once, err := ethaddr.Normalize(checksum)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	twice, err := ethaddr.Normalize(once)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if once != twice {
		t.Fatalf("normalize is not idempotent: %q then %q", once, twice)
	}
}

func TestNormalizeNonZero(t *testing.T) {
	got, err := ethaddr.NormalizeNonZero(checksum)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got != lowercase {
		t.Fatalf("expected %q, got %q", lowercase, got)
	}

	if _, err := ethaddr.NormalizeNonZero(ethaddr.ZeroAddress); !errors.Is(err, ethaddr.ErrZeroNotAllowed) {
		t.Fatalf("expected ErrZeroNotAllowed, got %v", err)
	}

	if _, err := ethaddr.NormalizeNonZero(""); !errors.Is(err, ethaddr.ErrEmpty) {
		t.Fatalf("expected ErrEmpty to pass through, got %v", err)
	}
}

func TestChecksum(t *testing.T) {
	cases := []struct {
		name  string
		input string
	}{
		{name: "from lowercase", input: lowercase},
		{name: "from uppercase", input: uppercase},
		{name: "from checksum", input: checksum},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ethaddr.Checksum(testCase.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != checksum {
				t.Fatalf("expected %q, got %q", checksum, got)
			}
		})
	}

	if _, err := ethaddr.Checksum("nonsense"); err == nil {
		t.Fatal("expected an error for an invalid address")
	}
}

func TestChecksumRoundTripsThroughNormalize(t *testing.T) {
	checksummed, err := ethaddr.Checksum(lowercase)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	back, err := ethaddr.Normalize(checksummed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if back != lowercase {
		t.Fatalf("expected %q, got %q", lowercase, back)
	}
}

func TestIsValid(t *testing.T) {
	valid := []string{lowercase, uppercase, checksum, ethaddr.ZeroAddress, " " + checksum + " "}
	invalid := []string{"", "   ", "0x", "not-an-address", lowercase[:41], lowercase + "a"}

	for _, input := range valid {
		if !ethaddr.IsValid(input) {
			t.Fatalf("expected %q to be valid", input)
		}
	}

	for _, input := range invalid {
		if ethaddr.IsValid(input) {
			t.Fatalf("expected %q to be invalid", input)
		}
	}
}

func TestIsZero(t *testing.T) {
	if !ethaddr.IsZero(ethaddr.ZeroAddress) {
		t.Fatal("expected the zero address to be reported as zero")
	}

	if !ethaddr.IsZero(strings.ToUpper(ethaddr.ZeroAddress[2:])[:0] + ethaddr.ZeroAddress) {
		t.Fatal("expected the zero address to be reported as zero regardless of case")
	}

	if ethaddr.IsZero(lowercase) {
		t.Fatal("expected a real address not to be reported as zero")
	}

	if ethaddr.IsZero("nonsense") {
		t.Fatal("expected an invalid address not to be reported as zero")
	}
}

func TestEqual(t *testing.T) {
	cases := []struct {
		name  string
		left  string
		right string
		want  bool
	}{
		{name: "same case", left: lowercase, right: lowercase, want: true},
		{name: "different case", left: uppercase, right: lowercase, want: true},
		{name: "checksum versus lowercase", left: checksum, right: lowercase, want: true},
		{name: "whitespace tolerated", left: " " + checksum, right: lowercase + " ", want: true},
		{name: "different addresses", left: lowercase, right: ethaddr.ZeroAddress, want: false},
		{name: "left invalid", left: "nonsense", right: lowercase, want: false},
		{name: "right invalid", left: lowercase, right: "nonsense", want: false},
		{name: "both invalid", left: "nonsense", right: "nonsense", want: false},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := ethaddr.Equal(testCase.left, testCase.right); got != testCase.want {
				t.Fatalf("expected %v, got %v", testCase.want, got)
			}
		})
	}
}
