package queryparam

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var ErrInvalid = errors.New("query parameter is not valid")

type ParamError struct {
	Param  string
	Reason string
}

func (e *ParamError) Error() string {
	return fmt.Sprintf("%s: %s: %s", ErrInvalid, e.Param, e.Reason)
}

func (e *ParamError) Unwrap() error {
	return ErrInvalid
}

func Message(err error) string {
	var paramError *ParamError
	if errors.As(err, &paramError) {
		return fmt.Sprintf("The %s parameter %s.", paramError.Param, paramError.Reason)
	}

	return "That request could not be read."
}

func invalid(param, reason string) error {
	return &ParamError{Param: param, Reason: reason}
}

func String(values url.Values, key string) string {
	return strings.TrimSpace(values.Get(key))
}

func Int(values url.Values, key string, fallback int) (int, error) {
	raw := String(values, key)
	if raw == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(raw)
	if err != nil {
		return 0, invalid(key, "must be a whole number")
	}

	return parsed, nil
}

func Time(values url.Values, key string) (*time.Time, error) {
	raw := String(values, key)
	if raw == "" {
		return nil, nil
	}

	parsed, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return nil, invalid(key, "must be an RFC 3339 timestamp such as 2026-08-20T03:20:11Z")
	}

	utc := parsed.UTC()

	return &utc, nil
}

func List(values url.Values, key string) []string {
	items := make([]string, 0)
	seen := make(map[string]struct{})

	for _, raw := range values[key] {
		for _, part := range strings.Split(raw, ",") {
			trimmed := strings.TrimSpace(part)
			if trimmed == "" {
				continue
			}

			if _, found := seen[trimmed]; found {
				continue
			}

			seen[trimmed] = struct{}{}
			items = append(items, trimmed)
		}
	}

	return items
}
