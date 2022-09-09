package redis

import (
	"github.com/go-redis/redis/v9"
	"strconv"
	"time"

	"github.com/swagftw/cache-service/utl/config"
)

// InitRedisDB initializes a new redis client.
func InitRedisDB(cfg *config.Config) *redis.Client {
	port := strconv.Itoa(cfg.Redis.Port)
	rdb := redis.NewClient(&redis.Options{ //nolint:exhaustivestruct
		Addr:         ":" + port,
		DB:           0, // use default DB
		ReadTimeout:  time.Duration(cfg.Redis.Timeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Redis.Timeout) * time.Second,
	})

	return rdb
}
