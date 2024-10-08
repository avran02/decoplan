package repository

import (
	"github.com/redis/go-redis/v9"
)

type Redis interface{}

type redisImpl struct {
	db *redis.Client
}

func (r *redisImpl) SaveMsg() {
}

func (r *redisImpl) DeleteMsg() {}

func (r *redisImpl) GetMsgs() {}

func NewRedis() Redis {
	return &redisImpl{
		connectRedis(),
	}
}

func connectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
