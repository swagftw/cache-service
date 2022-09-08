package types

import "context"

type CacheService interface {
	// Set method sets data to cache
	Set(ctx context.Context, key string, value []byte) error

	// Get method gets data from cache
	Get(ctx context.Context, key string) ([]byte, error)
}
