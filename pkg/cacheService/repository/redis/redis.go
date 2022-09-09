package redis

import (
	"context"
	"fmt"
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

func (r repo) GetUser(ctx context.Context, name string, rollNo int64) ([]byte, error) {
	keys, err := r.db.Keys(ctx, fmt.Sprintf("*%s:*:%d*", name, rollNo)).Result()
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, fmt.Errorf("user not found") // not a good practice to create errors like this
	} else if len(keys) > 1 {
		return nil, fmt.Errorf("multiple users found")
	}

	user, err := r.db.Get(ctx, keys[0]).Result()
	if err != nil {
		return nil, err
	}

	return []byte(user), nil
}

func (r repo) SetUser(ctx context.Context, key string, user []byte) error {
	return r.db.Set(ctx, key, user, 0).Err()
}
