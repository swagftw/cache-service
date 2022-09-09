package cacheService

import (
	"context"

	"github.com/swagftw/cache-service/types"
)

func InitCacheService(repo Repository) types.CacheService {
	return &service{
		Repository: repo,
	}
}

type service struct {
	Repository Repository
}

func (s service) Set(ctx context.Context, key string, value []byte) error {
	key = s.prefixKey(ctx, key)

	return s.Repository.SetValue(ctx, key, string(value))
}

func (s service) Get(ctx context.Context, key string) ([]byte, error) {
	key = s.prefixKey(ctx, key)

	value, err := s.Repository.GetValue(ctx, key)
	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

// prefixKey adds a prefix to the key.
// This can also be done by decorator pattern.
func (s service) prefixKey(ctx context.Context, key string) string {
	return "swapnil:" + key
}
