package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type redisJWTHandler struct {
	client redis.Cmdable
}

func NewRedisJWTHandler(cli redis.Cmdable) Handler {
	return &redisJWTHandler{
		client: cli,
	}
}

func (r *redisJWTHandler) ClearToken(ctx *gin.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *redisJWTHandler) ExtractToken(ctx *gin.Context) string {
	//TODO implement me
	panic("implement me")
}

func (r *redisJWTHandler) SetLoginToken(ctx *gin.Context, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *redisJWTHandler) setJwtToken(ctx *gin.Context, uid int64, ssid string) error {
	//TODO implement me
	panic("implement me")
}

func (r *redisJWTHandler) getJwtToken(ctx *gin.Context, ssid string) error {
	//TODO implement me
	panic("implement me")
}
