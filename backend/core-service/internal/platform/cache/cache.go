package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

var ErrNotConfigured = errors.New("cache is not configured")

type Store interface {
	Get(ctx context.Context, key string) ([]byte, bool, error)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, keys ...string) error
	Ping(ctx context.Context) error
	Close() error
}

type Stats struct {
	Hits   int64
	Misses int64
	Errors int64
}

func GetOrSet[T any](
	ctx context.Context,
	store Store,
	key string,
	ttl time.Duration,
	loader func(context.Context) (T, error),
) (T, error) {
	var empty T

	if loader == nil {
		return empty, errors.New("cache loader must not be nil")
	}

	if store != nil && key != "" {
		if cached, found, err := store.Get(ctx, key); err == nil && found {
			var decoded T

			if json.Unmarshal(cached, &decoded) == nil {
				return decoded, nil
			}
		}
	}

	loaded, err := loader(ctx)
	if err != nil {
		return empty, err
	}

	if store == nil || key == "" || ttl <= 0 {
		return loaded, nil
	}

	encoded, err := json.Marshal(loaded)
	if err != nil {
		return loaded, nil
	}

	_ = store.Set(ctx, key, encoded, ttl)

	return loaded, nil
}

func Key(parts ...string) string {
	if len(parts) == 0 {
		return ""
	}

	key := parts[0]
	for _, part := range parts[1:] {
		key += ":" + part
	}

	return key
}
