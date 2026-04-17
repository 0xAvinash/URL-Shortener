package database

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var RedisClient *redis.Client

func NewRedis() *redis.Client {
	redis_host := os.Getenv("REDIS_HOST")
	if redis_host == "" {
		redis_host = "localhost"
	}

	host := fmt.Sprintf("%s:6379", redis_host)
	RedisClient = redis.NewClient(&redis.Options{
		Addr: host,
	})

	return RedisClient
}
