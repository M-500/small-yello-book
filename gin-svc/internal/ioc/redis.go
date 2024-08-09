package ioc

import (
	"context"
	"fmt"
	"gin-svc/internal/conf/cfg"
	"github.com/redis/go-redis/v9"
)

func SetUpRedis(cfg *cfg.RedisCfg) redis.Cmdable {
	addr := fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return redisClient
}
