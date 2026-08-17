package cache_test

import (
	"context"
	"errors"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/farhapartex/lending-platform/core-service/internal/platform/cache"
)

type sample struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func redisURL() string {
	if url := os.Getenv("TEST_REDIS_URL"); url != "" {
		return url
	}

	return "redis://localhost:6379/9"
}

func newStore(t *testing.T) cache.Store {
	t.Helper()

	store, err := cache.NewRedis(cache.RedisParams{
		URL:            redisURL(),
		Namespace:      "test-" + strconv.FormatInt(time.Now().UnixNano(), 36),
		RequestTimeout: 2 * time.Second,
	})
	if err != nil {
		t.Fatalf("unexpected error building the store: %v", err)
	}

	if err := store.Ping(context.Background()); err != nil {
		_ = store.Close()
		t.Skipf("redis is not reachable at %s: %v", redisURL(), err)
	}

	t.Cleanup(func() {
		_ = store.Close()
	})

	return store
}

func TestNewRedisRejectsMissingURL(t *testing.T) {
	if _, err := cache.NewRedis(cache.RedisParams{}); !errors.Is(err, cache.ErrNotConfigured) {
		t.Fatalf("expected ErrNotConfigured, got %v", err)
	}
}

func TestNewRedisRejectsMalformedURL(t *testing.T) {
	_, err := cache.NewRedis(cache.RedisParams{URL: "not-a-redis-url"})
	if err == nil {
		t.Fatal("expected an error for a malformed url")
	}

	if errors.Is(err, cache.ErrNotConfigured) {
		t.Fatalf("expected a parse error rather than ErrNotConfigured, got %v", err)
	}
}

func TestSetThenGet(t *testing.T) {
	store := newStore(t)
	ctx := context.Background()

	if err := store.Set(ctx, "greeting", []byte("hello"), time.Minute); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	value, found, err := store.Get(ctx, "greeting")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !found {
		t.Fatal("expected the key to be found")
	}

	if string(value) != "hello" {
		t.Fatalf("expected hello, got %q", value)
	}
}

func TestGetReportsAMissWithoutError(t *testing.T) {
	store := newStore(t)

	value, found, err := store.Get(context.Background(), "never-written")
	if err != nil {
		t.Fatalf("expected a miss to be reported without an error, got %v", err)
	}

	if found {
		t.Fatal("expected the key to be missing")
	}

	if value != nil {
		t.Fatalf("expected no value, got %q", value)
	}
}

func TestValueExpires(t *testing.T) {
	store := newStore(t)
	ctx := context.Background()

	if err := store.Set(ctx, "brief", []byte("gone soon"), 50*time.Millisecond); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, found, _ := store.Get(ctx, "brief"); !found {
		t.Fatal("expected the key to exist before expiry")
	}

	time.Sleep(120 * time.Millisecond)

	if _, found, _ := store.Get(ctx, "brief"); found {
		t.Fatal("expected the key to be gone after expiry")
	}
}

func TestDelete(t *testing.T) {
	store := newStore(t)
	ctx := context.Background()

	for _, key := range []string{"first", "second"} {
		if err := store.Set(ctx, key, []byte("value"), time.Minute); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if err := store.Delete(ctx, "first", "second"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, key := range []string{"first", "second"} {
		if _, found, _ := store.Get(ctx, key); found {
			t.Fatalf("expected %s to be deleted", key)
		}
	}
}

func TestDeleteWithNoKeysIsHarmless(t *testing.T) {
	store := newStore(t)

	if err := store.Delete(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNamespaceKeepsStoresApart(t *testing.T) {
	ctx := context.Background()

	first, err := cache.NewRedis(cache.RedisParams{URL: redisURL(), Namespace: "alpha"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func() { _ = first.Close() }()

	if err := first.Ping(ctx); err != nil {
		t.Skipf("redis is not reachable: %v", err)
	}

	second, err := cache.NewRedis(cache.RedisParams{URL: redisURL(), Namespace: "beta"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func() { _ = second.Close() }()

	if err := first.Set(ctx, "shared", []byte("from alpha"), time.Minute); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Cleanup(func() {
		_ = first.Delete(ctx, "shared")
	})

	if _, found, _ := second.Get(ctx, "shared"); found {
		t.Fatal("expected namespaces to keep keys apart")
	}
}

func TestGetOrSetLoadsOnceThenServesFromCache(t *testing.T) {
	store := newStore(t)
	ctx := context.Background()

	calls := 0
	loader := func(context.Context) (sample, error) {
		calls++

		return sample{Name: "market", Count: 42}, nil
	}

	first, err := cache.GetOrSet(ctx, store, "market", time.Minute, loader)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	second, err := cache.GetOrSet(ctx, store, "market", time.Minute, loader)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if calls != 1 {
		t.Fatalf("expected the loader to run once, it ran %d times", calls)
	}

	if first != second {
		t.Fatalf("expected identical values, got %+v and %+v", first, second)
	}

	if second.Name != "market" || second.Count != 42 {
		t.Fatalf("unexpected cached value %+v", second)
	}
}

func TestGetOrSetPropagatesLoaderErrors(t *testing.T) {
	store := newStore(t)

	wanted := errors.New("database is down")

	_, err := cache.GetOrSet(context.Background(), store, "failing", time.Minute,
		func(context.Context) (sample, error) {
			return sample{}, wanted
		})

	if !errors.Is(err, wanted) {
		t.Fatalf("expected the loader error to surface, got %v", err)
	}
}

func TestGetOrSetSurvivesAnUnreachableCache(t *testing.T) {
	store, err := cache.NewRedis(cache.RedisParams{
		URL:            "redis://127.0.0.1:1/0",
		RequestTimeout: 200 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func() { _ = store.Close() }()

	calls := 0

	value, err := cache.GetOrSet(context.Background(), store, "market", time.Minute,
		func(context.Context) (sample, error) {
			calls++

			return sample{Name: "fallback", Count: 1}, nil
		})

	if err != nil {
		t.Fatalf("expected an unreachable cache to degrade rather than fail, got %v", err)
	}

	if calls != 1 {
		t.Fatalf("expected the loader to run, it ran %d times", calls)
	}

	if value.Name != "fallback" {
		t.Fatalf("expected the loaded value, got %+v", value)
	}
}

func TestGetOrSetFallsBackWhenCachedBytesAreCorrupt(t *testing.T) {
	store := newStore(t)
	ctx := context.Background()

	if err := store.Set(ctx, "corrupt", []byte("{not json"), time.Minute); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	calls := 0

	value, err := cache.GetOrSet(ctx, store, "corrupt", time.Minute,
		func(context.Context) (sample, error) {
			calls++

			return sample{Name: "rebuilt", Count: 7}, nil
		})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if calls != 1 {
		t.Fatalf("expected the loader to rebuild the value, it ran %d times", calls)
	}

	if value.Name != "rebuilt" {
		t.Fatalf("expected the rebuilt value, got %+v", value)
	}
}

func TestGetOrSetWithoutAStoreStillLoads(t *testing.T) {
	calls := 0

	value, err := cache.GetOrSet(context.Background(), nil, "market", time.Minute,
		func(context.Context) (sample, error) {
			calls++

			return sample{Name: "direct", Count: 3}, nil
		})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if calls != 1 || value.Name != "direct" {
		t.Fatalf("expected a direct load, got %d calls and %+v", calls, value)
	}
}

func TestGetOrSetSkipsCachingWhenKeyOrTTLIsUnset(t *testing.T) {
	store := newStore(t)
	ctx := context.Background()

	calls := 0
	loader := func(context.Context) (sample, error) {
		calls++

		return sample{Name: "uncached", Count: 1}, nil
	}

	if _, err := cache.GetOrSet(ctx, store, "", time.Minute, loader); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := cache.GetOrSet(ctx, store, "zero-ttl", 0, loader); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := cache.GetOrSet(ctx, store, "zero-ttl", 0, loader); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if calls != 3 {
		t.Fatalf("expected every call to load fresh, got %d", calls)
	}
}

func TestGetOrSetRejectsMissingLoader(t *testing.T) {
	store := newStore(t)

	if _, err := cache.GetOrSet[sample](context.Background(), store, "key", time.Minute, nil); err == nil {
		t.Fatal("expected an error when no loader is supplied")
	}
}

func TestGetOrSetHandlesUnserialisableValues(t *testing.T) {
	store := newStore(t)

	calls := 0

	value, err := cache.GetOrSet(context.Background(), store, "channel", time.Minute,
		func(context.Context) (chan int, error) {
			calls++

			return make(chan int), nil
		})

	if err != nil {
		t.Fatalf("expected an unserialisable value to be returned rather than failing, got %v", err)
	}

	if calls != 1 || value == nil {
		t.Fatalf("expected the loaded value, got %d calls", calls)
	}
}

func TestNoopStore(t *testing.T) {
	store := cache.NewNoop()
	ctx := context.Background()

	if err := store.Set(ctx, "key", []byte("value"), time.Minute); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, found, err := store.Get(ctx, "key"); err != nil || found {
		t.Fatalf("expected the noop store to never find anything, got found=%v err=%v", found, err)
	}

	if err := store.Delete(ctx, "key"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := store.Ping(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := store.Close(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNoopStoreStillLoadsThroughGetOrSet(t *testing.T) {
	calls := 0

	for range 3 {
		if _, err := cache.GetOrSet(context.Background(), cache.NewNoop(), "key", time.Minute,
			func(context.Context) (sample, error) {
				calls++

				return sample{Name: "always fresh"}, nil
			}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	if calls != 3 {
		t.Fatalf("expected three loads through the noop store, got %d", calls)
	}
}

func TestKey(t *testing.T) {
	cases := []struct {
		name  string
		parts []string
		want  string
	}{
		{name: "no parts", parts: nil, want: ""},
		{name: "single part", parts: []string{"markets"}, want: "markets"},
		{name: "two parts", parts: []string{"markets", "1"}, want: "markets:1"},
		{name: "many parts", parts: []string{"accounts", "0xabc", "transactions", "page1"}, want: "accounts:0xabc:transactions:page1"},
		{name: "empty part is kept", parts: []string{"markets", ""}, want: "markets:"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := cache.Key(testCase.parts...); got != testCase.want {
				t.Fatalf("expected %q, got %q", testCase.want, got)
			}
		})
	}
}
