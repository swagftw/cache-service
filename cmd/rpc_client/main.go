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
		get(client, args[1])
	} else if args[0] == "set" && len(args) == 3 {
		set(client, args[1], args[2])
	} else if args[0] == "getuser" && len(args) == 3 {
		rollNum, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			log.Printf("Failed to parse roll number: %v\n", err)

			return
		}

		getUser(client, args[1], rollNum)
	} else if args[0] == "setuser" && len(args) == 4 {
		rollNum, err := strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			log.Printf("Failed to parse roll number: %v\n", err)

			return
		}

		setUser(client, args[1], args[2], rollNum)
	} else {
		log.Printf("Invalid command: %s\n", args[0])
	}
}

func get(client cacheService.CacheServiceClient, key string) {
	response, err := client.GetValue(context.TODO(), &cacheService.GetRequest{Key: key})
	if err != nil {
		log.Printf("Failed to get value: %v\n", err)

		return
	}

	log.Printf("Value: %s\n", response.Value)
}

func set(client cacheService.CacheServiceClient, key string, value string) {
	_, err := client.SetValue(context.TODO(), &cacheService.SetRequest{Key: key, Value: value})
	if err != nil {
		log.Printf("Failed to set value: %v\n", err)

		return
	}

	log.Println("OK")
}

func setUser(client cacheService.CacheServiceClient, name, class string, rollNum int64) {
	resp, err := client.SetUser(context.TODO(), &cacheService.User{
		Name:     name,
		Class:    class,
		RollNum:  rollNum,
		Metadata: nil,
	})

	if err != nil {
		log.Printf("Failed to set user: %v\n", err)

		return
	}

	log.Println(resp)
}

func getUser(client cacheService.CacheServiceClient, name string, rollNum int64) {
	resp, err := client.GetUser(context.TODO(), &cacheService.GetUserRequest{
		Name:    name,
		RollNum: rollNum,
	})

	if err != nil {
		log.Printf("Failed to get user: %v\n", err)

		return
	}

	log.Println(resp)
}
