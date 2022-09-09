package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cacheService "github.com/swagftw/cache-service/transport/rpc/cache_service"
	"github.com/swagftw/cache-service/utl/config"
)

func main() {
	args := os.Args[1:]

	cfg := config.InitConfig("./utl/config/config.local.yaml")

	connection, err := grpc.Dial("localhost:"+strconv.Itoa(cfg.Server.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v\n", err)
	}

	defer func(connection *grpc.ClientConn) {
		err = connection.Close()
		if err != nil {
			log.Fatalf("Failed to close connection: %v\n", err)
		}
	}(connection)

	client := cacheService.NewCacheServiceClient(connection)

	if args[0] == "get" && len(args) == 2 {
		Get(client, args[1])
	} else if args[0] == "set" && len(args) == 3 {
		Set(client, args[1], args[2])
	} else {
		log.Printf("Invalid command: %s\n", args[0])
	}
}

func Get(client cacheService.CacheServiceClient, key string) {
	response, err := client.GetValue(context.TODO(), &cacheService.GetRequest{Key: key})
	if err != nil {
		log.Printf("Failed to get value: %v\n", err)

		return
	}

	log.Printf("Value: %s\n", response.Value)
}

func Set(client cacheService.CacheServiceClient, key string, value string) {
	_, err := client.SetValue(context.TODO(), &cacheService.SetRequest{Key: key, Value: value})
	if err != nil {
		log.Printf("Failed to set value: %v\n", err)
		
		return
	}

	log.Println("OK")
}
