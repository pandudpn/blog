package dbc

import (
	"context"
	"log"
	"os"
	
	"github.com/go-redis/redis/v8"
)

// NewConnectionRedis to connect Redis
func NewConnectionRedis() *redis.Client {
	opt := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	
	client := redis.NewClient(opt)
	
	err := client.Ping(context.Background())
	if err != nil {
		log.Fatalln("failed to ping redis", err)
	}
	
	return client
}
