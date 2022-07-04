package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-redis/redis/v9"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	defer lis.Close()

	connectRedis()

	for {
		_, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Println("hello golang application")
	}
}

func connectRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-svc:6379",
		Password: "",
		DB:       0,
	})

	fmt.Println(client)
	if client == nil {
		log.Fatal("cannot connect to Redis")
	}

	err := client.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect to Redis successfully")
}
