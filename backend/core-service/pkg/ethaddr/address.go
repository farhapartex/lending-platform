package ethaddr

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

const (
	Length      = 42
	ZeroAddress = "0x0000000000000000000000000000000000000000"
	hexPrefix   = "0x"
)

var (
	ErrEmpty          = errors.New("ethereum address must not be empty")
	ErrMissingPrefix  = errors.New("ethereum address must start with 0x")
	ErrInvalidLength  = errors.New("ethereum address must be 42 characters")
	ErrInvalidHex     = errors.New("ethereum address must be hexadecimal")
	ErrZeroNotAllowed = errors.New("ethereum address must not be the zero address")
)

func Normalize(raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)

	if trimmed == "" {
		return "", ErrEmpty
	}

	if !strings.HasPrefix(trimmed, hexPrefix) {
		return "", ErrMissingPrefix
	}

	if len(trimmed) != Length {
		return "", ErrInvalidLength
	}

	if !isHexBody(trimmed[len(hexPrefix):]) {
		return "", ErrInvalidHex
	}

	return strings.ToLower(trimmed), nil
}

func NormalizeNonZero(raw string) (string, error) {
	normalized, err := Normalize(raw)
	if err != nil {
		return "", err
	}

	if normalized == ZeroAddress {
		return "", ErrZeroNotAllowed
	}

	return normalized, nil
}

func NormalizeWithChecksum(raw string) (normalized string, checksum string, err error) {
	normalized, err = Normalize(raw)
	if err != nil {
		return "", "", err
	}

	return normalized, common.HexToAddress(normalized).Hex(), nil
}

func Checksum(raw string) (string, error) {
	normalized, err := Normalize(raw)
	if err != nil {
		return "", err
	}

	return common.HexToAddress(normalized).Hex(), nil
}

func IsValid(raw string) bool {
	_, err := Normalize(raw)

	return err == nil
}

func IsZero(raw string) bool {
	normalized, err := Normalize(raw)

	return err == nil && normalized == ZeroAddress
}

func Equal(left, right string) bool {
	normalizedLeft, err := Normalize(left)
	if err != nil {
		return false
	}

	normalizedRight, err := Normalize(right)
	if err != nil {
		return false
	}

	return normalizedLeft == normalizedRight
}

func isHexBody(body string) bool {
	for _, char := range body {
		switch {
		case char >= '0' && char <= '9':
		case char >= 'a' && char <= 'f':
		case char >= 'A' && char <= 'F':
		default:
			return false
		}
	}

	return true
}
