package cache

import (
	"context"
	"time"
)

type noopStore struct{}

func NewNoop() Store {
	return noopStore{}
}

func (noopStore) Get(context.Context, string) ([]byte, bool, error) {
	return nil, false, nil
}

func (noopStore) Set(context.Context, string, []byte, time.Duration) error {
	return nil
}

func (noopStore) Delete(context.Context, ...string) error {
	return nil
}

func (noopStore) Ping(context.Context) error {
	return nil
}

func (noopStore) Close() error {
	return nil
}
