package rpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	cs "github.com/swagftw/cache-service/pkg/cacheService"
	rr "github.com/swagftw/cache-service/pkg/cacheService/repository/redis"
	cacheServiceRPC "github.com/swagftw/cache-service/transport/rpc/cache_service"
	"github.com/swagftw/cache-service/utl/config"
	"github.com/swagftw/cache-service/utl/storage/redis"
)

func Start(configPath string) {
	cfg := config.InitConfig(configPath)

	redisDB := redis.InitRedisDB(cfg)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	// creates new grpc server
	grpcServer := grpc.NewServer(grpc.ConnectionTimeout(time.Duration(cfg.Server.Timeout) * time.Second))

	// creates new cache service server
	cacheServiceServer := cacheServiceRPC.NewCacheServiceSrv(cs.InitCacheService(rr.InitRedisRepo(redisDB, cfg)))
	cacheServiceRPC.RegisterCacheServiceServer(grpcServer, cacheServiceServer)

	// check for interrupt signal
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// spawns new goroutine to serve grpc server
	go func() {
		log.Printf("Starting gRPC server on port %d\n", cfg.Server.Port)

		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatalf("Failed to serve: %v\n", err)
		}

		interrupt <- syscall.SIGTERM
	}()

	// waits for interrupt signal to gracefully shut down the server with a timeout of 5 seconds.
	if <-interrupt; true {
		log.Println("Shutting down gRPC server...")
		grpcServer.GracefulStop()
	}
}
