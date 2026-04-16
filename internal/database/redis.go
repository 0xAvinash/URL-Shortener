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

	host := fmt.Sprintf("%s:6379", redis_host)
	RedisClient = redis.NewClient(&redis.Options{
		Addr: host,
	})

	return RedisClient
}
