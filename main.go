package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/go-redis/redis/v9"
	"gopkg.in/yaml.v3"
)

type config struct {
	RedisCfg redisConfig `yaml:"redis"`
}

type redisConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	defer lis.Close()

	// connectRedis()

	// fmt.Println("value of FOO: ", os.Getenv("FOO"))
	// fmt.Println("value of BAR: ", os.Getenv("BAR"))
	cfg := getConfig()

	fmt.Println("value user redis user: ", cfg.RedisCfg.User)
	fmt.Println("value user redis password: ", cfg.RedisCfg.Password)

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

func getConfig() *config {
	bytes, err := ioutil.ReadFile("configs/env.yaml")
	if err != nil {
		panic(err)
	}

	cfg := &config{}

	err = yaml.Unmarshal(bytes, cfg)

	fmt.Println(cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}
