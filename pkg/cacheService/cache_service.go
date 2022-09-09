package cacheService

import (
	"context"
	"encoding/json"
	"fmt"

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

func (s service) SetUser(ctx context.Context, user *types.User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s:%s:%d", user.Name, user.Class, user.RollNum)
	key = s.prefixKey(ctx, key)

	return s.Repository.SetUser(ctx, key, bytes)
}

func (s service) GetUser(ctx context.Context, name string, rollNo int64) (*types.User, error) {
	userBytes, err := s.Repository.GetUser(ctx, name, rollNo)
	if err != nil {
		return nil, err
	}

	user := new(types.User)

	err = json.Unmarshal(userBytes, user)
	if err != nil {
		return nil, err
	}

	return user, nil
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
