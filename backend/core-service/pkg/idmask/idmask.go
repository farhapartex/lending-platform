package idmask

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base32"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
)

type Kind string

const (
	KindUser         Kind = "usr"
	KindMarket       Kind = "mkt"
	KindPosition     Kind = "pos"
	KindTransaction  Kind = "txn"
	KindEvent        Kind = "evt"
	KindLiquidation  Kind = "liq"
	KindSnapshot     Kind = "snp"
	KindNotification Kind = "ntf"
)

const (
	separator   = "_"
	feistelPass = 4
	tagBytes    = 4
	payloadSize = 8 + tagBytes
)

var (
	ErrEmptySecret  = errors.New("id mask secret must not be empty")
	ErrEmptyKind    = errors.New("id mask kind must not be empty")
	ErrIDOutOfRange = errors.New("id must be a positive integer")
	ErrInvalidToken = errors.New("id is not a valid masked identifier")
	ErrWrongKind    = errors.New("id belongs to a different resource type")
)

var encoding = base32.StdEncoding.WithPadding(base32.NoPadding)

type Masker struct {
	key []byte
}

func New(secret string) (*Masker, error) {
	trimmed := strings.TrimSpace(secret)
	if trimmed == "" {
		return nil, ErrEmptySecret
	}

	digest := sha256.Sum256([]byte(trimmed))
	key := make([]byte, len(digest))
	copy(key, digest[:])

	return &Masker{key: key}, nil
}

func (m *Masker) Mask(kind Kind, id int64) (string, error) {
	if kind == "" {
		return "", ErrEmptyKind
	}

	if id < 1 {
		return "", fmt.Errorf("%w: got %d", ErrIDOutOfRange, id)
	}

	permuted := m.permute(kind, uint64(id))

	payload := make([]byte, payloadSize)
	binary.BigEndian.PutUint64(payload[:8], permuted)
	copy(payload[8:], m.tag(kind, payload[:8]))

	return string(kind) + separator + strings.ToLower(encoding.EncodeToString(payload)), nil
}

func (m *Masker) Unmask(kind Kind, token string) (int64, error) {
	if kind == "" {
		return 0, ErrEmptyKind
	}

	trimmed := strings.TrimSpace(token)
	if trimmed == "" {
		return 0, ErrInvalidToken
	}

	prefix, encoded, found := strings.Cut(trimmed, separator)
	if !found || encoded == "" {
		return 0, ErrInvalidToken
	}

	if prefix != string(kind) {
		return 0, fmt.Errorf("%w: expected %s", ErrWrongKind, kind)
	}

	payload, err := encoding.DecodeString(strings.ToUpper(encoded))
	if err != nil {
		return 0, ErrInvalidToken
	}

	if len(payload) != payloadSize {
		return 0, ErrInvalidToken
	}

	expectedTag := m.tag(kind, payload[:8])
	if subtle.ConstantTimeCompare(payload[8:], expectedTag) != 1 {
		return 0, ErrInvalidToken
	}

	id := m.reverse(kind, binary.BigEndian.Uint64(payload[:8]))
	if id < 1 || id > uint64(1)<<62 {
		return 0, ErrInvalidToken
	}

	return int64(id), nil
}

func (m *Masker) MustMask(kind Kind, id int64) string {
	masked, err := m.Mask(kind, id)
	if err != nil {
		panic(err)
	}

	return masked
}

func (m *Masker) permute(kind Kind, value uint64) uint64 {
	left, right := uint32(value>>32), uint32(value)

	for round := range feistelPass {
		left, right = right, left^m.round(kind, round, right)
	}

	return uint64(left)<<32 | uint64(right)
}

func (m *Masker) reverse(kind Kind, value uint64) uint64 {
	left, right := uint32(value>>32), uint32(value)

	for round := feistelPass - 1; round >= 0; round-- {
		previousRight := left
		previousLeft := right ^ m.round(kind, round, previousRight)
		left, right = previousLeft, previousRight
	}

	return uint64(left)<<32 | uint64(right)
}

func (m *Masker) round(kind Kind, round int, half uint32) uint32 {
	mac := hmac.New(sha256.New, m.key)
	mac.Write([]byte(kind))
	mac.Write([]byte{byte(round)})

	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, half)
	mac.Write(buffer)

	return binary.BigEndian.Uint32(mac.Sum(nil)[:4])
}

func (m *Masker) tag(kind Kind, permuted []byte) []byte {
	mac := hmac.New(sha256.New, m.key)
	mac.Write([]byte("tag"))
	mac.Write([]byte(kind))
	mac.Write(permuted)

	return mac.Sum(nil)[:tagBytes]
}
