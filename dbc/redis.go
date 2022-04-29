package dbc

import (
	"os"
	
	"github.com/go-redis/redis/v8"
)

// NewConnectionRedis to connect Redis
func NewConnectionRedis() (*redis.Client, error) {
	opt := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	
	client := redis.NewClient(opt)
	
	return client, nil
}
