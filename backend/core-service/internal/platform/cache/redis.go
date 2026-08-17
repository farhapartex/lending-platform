package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisParams struct {
	URL            string
	Namespace      string
	RequestTimeout time.Duration
}

type redisStore struct {
	client         *redis.Client
	namespace      string
	requestTimeout time.Duration
}

const defaultRequestTimeout = 2 * time.Second

func NewRedis(params RedisParams) (Store, error) {
	if params.URL == "" {
		return nil, ErrNotConfigured
	}

	options, err := redis.ParseURL(params.URL)
	if err != nil {
		return nil, fmt.Errorf("REDIS_URL is not a usable redis url: %w", err)
	}

	timeout := params.RequestTimeout
	if timeout <= 0 {
		timeout = defaultRequestTimeout
	}

	return &redisStore{
		client:         redis.NewClient(options),
		namespace:      params.Namespace,
		requestTimeout: timeout,
	}, nil
}

func (s *redisStore) Get(ctx context.Context, key string) ([]byte, bool, error) {
	callCtx, cancel := context.WithTimeout(ctx, s.requestTimeout)
	defer cancel()

	value, err := s.client.Get(callCtx, s.qualify(key)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, false, nil
		}

		return nil, false, err
	}

	return value, true, nil
}

func (s *redisStore) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	callCtx, cancel := context.WithTimeout(ctx, s.requestTimeout)
	defer cancel()

	return s.client.Set(callCtx, s.qualify(key), value, ttl).Err()
}

func (s *redisStore) Delete(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	callCtx, cancel := context.WithTimeout(ctx, s.requestTimeout)
	defer cancel()

	qualified := make([]string, 0, len(keys))
	for _, key := range keys {
		qualified = append(qualified, s.qualify(key))
	}

	return s.client.Del(callCtx, qualified...).Err()
}

func (s *redisStore) Ping(ctx context.Context) error {
	callCtx, cancel := context.WithTimeout(ctx, s.requestTimeout)
	defer cancel()

	return s.client.Ping(callCtx).Err()
}

func (s *redisStore) Close() error {
	return s.client.Close()
}

func (s *redisStore) qualify(key string) string {
	if s.namespace == "" {
		return key
	}

	return s.namespace + ":" + key
}
