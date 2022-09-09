package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/swagftw/cache-service/pkg/cacheService"
	"github.com/swagftw/cache-service/utl/config"
)

type repo struct {
	db  *redis.Client
	cfg *config.Config
}

func InitRedisRepo(redisDB *redis.Client, cfg *config.Config) cacheService.Repository {
	return &repo{db: redisDB, cfg: cfg}
}

func (r repo) GetValue(ctx context.Context, key string) (string, error) {
	return r.db.Get(ctx, key).Result()
}

func (r repo) SetValue(ctx context.Context, key string, value string) error {
	return r.db.Set(ctx, key, value, 0).Err()
}
