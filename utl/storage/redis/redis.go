package redis

import (
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/swagftw/cache-service/utl/config"
)

// InitRedisDB initializes a new redis client.
func InitRedisDB(cfg *config.Config) *redis.Client {
	port := strconv.Itoa(cfg.Redis.Port)
	host, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		host = "localhost"
	}
	rdb := redis.NewClient(&redis.Options{ //nolint:exhaustivestruct
		Addr:         host + ":" + port,
		DB:           0, // use default DB
		ReadTimeout:  time.Duration(cfg.Redis.Timeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Redis.Timeout) * time.Second,
	})

	return rdb
}
