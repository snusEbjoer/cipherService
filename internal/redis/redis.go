package redis

import (
	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr(),
		DB:   database(),
	})
	return rdb
}
