package cache

import (
	"time"
	
	"github.com/go-redis/redis/v8"
)

type cacheRepository struct {
	redis    *redis.Client
	timezone *time.Location
}

// New is constructor of Package Cache Repository
// and will return an instance of Redis Client (query into redis)
func New(r *redis.Client) *cacheRepository {
	timezone, _ := time.LoadLocation("Asia/Jakarta")
	return &cacheRepository{
		redis:    r,
		timezone: timezone,
	}
}
