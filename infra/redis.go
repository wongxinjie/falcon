package infra

import (
	"github.com/go-redis/redis/v8"

	"falcon/config"
)

func NewRedisClient(c *config.RedisConf) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})

	return rdb, nil
}
