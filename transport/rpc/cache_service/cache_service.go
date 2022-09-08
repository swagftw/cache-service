package cache_service

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
	//TODO implement me
	panic("implement me")
}

func (c CacheServiceSrv) SetValue(ctx context.Context, request *SetRequest) (*SetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CacheServiceSrv) mustEmbedUnimplementedCacheServiceServer() {
	//TODO implement me
	panic("implement me")
}
