package dbc

import (
	"fmt"
	"os"
	
	"github.com/go-redis/redis/v8"
)

// NewConnectionRedis to connect Redis
func NewConnectionRedis() (*redis.Client, error) {
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	
	client := redis.NewClient(opt)
	
	return client, nil
}
