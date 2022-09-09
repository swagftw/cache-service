package cacheService

import "context"

type (
	Repository interface {
		GetValue(ctx context.Context, key string) (string, error)
		SetValue(ctx context.Context, key string, value string) error
	}
)
