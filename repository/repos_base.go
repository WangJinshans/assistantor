package repository

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	engine      *gorm.DB
	redisClient *redis.Client
)

func SetupEngine(e *gorm.DB) {
	engine = e
}

func SetupRedisClient(c *redis.Client) {
	redisClient = c
}

func GetEngine() (e *gorm.DB) {
	e = engine
	return
}

func GetRedisClient() (c *redis.Client) {
	c = redisClient
	return
}
