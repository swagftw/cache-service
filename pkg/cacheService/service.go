package cacheService

import (
	"context"
)

type (
	Repository interface {
		GetValue(ctx context.Context, key string) (string, error)
		SetValue(ctx context.Context, key string, value string) error
		GetUser(ctx context.Context, name string, rollNo int64) ([]byte, error)
		SetUser(ctx context.Context, key string, user []byte) error
	}
)
