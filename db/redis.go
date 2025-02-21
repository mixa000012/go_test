package db

import (
	"github.com/redis/go-redis/v9"
)

func InitRedis(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return client
}
