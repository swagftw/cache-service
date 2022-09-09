package cacheService

import (
	"context"

	"github.com/swagftw/cache-service/types"
)

// CacheServiceSrv is the server API struct which implements CacheService.
type CacheServiceSrv struct {
	CacheService types.CacheService
	UnimplementedCacheServiceServer
}

// NewCacheServiceSrv creates a new CacheServiceSrv.
func NewCacheServiceSrv(cacheService types.CacheService) CacheServiceServer {
	return &CacheServiceSrv{
		CacheService: cacheService,
	}
}

func (c CacheServiceSrv) GetValue(ctx context.Context, request *GetRequest) (*GetResponse, error) {
	data, err := c.CacheService.Get(ctx, request.Key)
	if err != nil {
		return nil, err
	}

	return &GetResponse{
		Value: string(data),
	}, nil
}

func (c CacheServiceSrv) SetValue(ctx context.Context, request *SetRequest) (*SetResponse, error) {
	err := c.CacheService.Set(ctx, request.Key, []byte(request.Value))
	if err != nil {
		return nil, err
	}

	return &SetResponse{Message: "ok"}, nil
}

func (c CacheServiceSrv) mustEmbedUnimplementedCacheServiceServer() {
	//TODO implement me
	panic("implement me")
}
