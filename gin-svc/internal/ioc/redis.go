package ioc

import (
	"gin-svc/internal/conf/cfg"
	"github.com/redis/go-redis/v9"
)

func SetUpRedis(cfg *cfg.RedisCfg) redis.Cmdable {
	addr := cfg.Addr
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
	})
	return redisClient
}
