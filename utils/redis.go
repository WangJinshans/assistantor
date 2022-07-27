package utils

import (
	"assistantor/config"
	"github.com/go-redis/redis"
)

func GetRedisClient(conf config.RedisConfig) (redisClient *redis.Client) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		DB:       conf.DB,
		Password: conf.Password,
	})
	return
}
