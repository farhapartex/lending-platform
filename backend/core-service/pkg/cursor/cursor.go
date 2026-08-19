package cursor

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	version   = "v1"
	separator = "."
)

var (
	ErrMalformed = errors.New("cursor is not readable")
	ErrEmpty     = errors.New("cursor must not be empty")
)

type Key struct {
	Time time.Time
	ID   int64
}

func (k Key) IsZero() bool {
	return k.Time.IsZero() && k.ID == 0
}

func Encode(key Key) string {
	payload := strings.Join([]string{
		version,
		strconv.FormatInt(key.Time.UTC().UnixMicro(), 10),
		strconv.FormatInt(key.ID, 10),
	}, separator)

	return base64.RawURLEncoding.EncodeToString([]byte(payload))
}

func Decode(raw string) (Key, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return Key{}, ErrEmpty
	}

	decoded, err := base64.RawURLEncoding.DecodeString(trimmed)
	if err != nil {
		return Key{}, ErrMalformed
	}

	parts := strings.Split(string(decoded), separator)
	if len(parts) != 3 || parts[0] != version {
		return Key{}, ErrMalformed
	}

	micros, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return Key{}, ErrMalformed
	}

	id, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return Key{}, ErrMalformed
	}

	if id < 0 {
		return Key{}, fmt.Errorf("%w: identifier must not be negative", ErrMalformed)
	}

	return Key{Time: time.UnixMicro(micros).UTC(), ID: id}, nil
}

func DecodeOptional(raw string) (Key, error) {
	if strings.TrimSpace(raw) == "" {
		return Key{}, nil
	}

	return Decode(raw)
}
