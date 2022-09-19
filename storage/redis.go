package storage

import "github.com/go-redis/redis/v7"

func NewClient(addr string) redis.Cmdable {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0, // use default DB
	})
	return rdb
}
