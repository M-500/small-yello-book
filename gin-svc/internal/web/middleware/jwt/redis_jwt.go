package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisJWTHandler struct {
	client        redis.Cmdable
	signingMethod jwt.SigningMethod
}

func NewRedisJWTHandler(cli redis.Cmdable) Handler {
	return &redisJWTHandler{
		client:        cli,
		signingMethod: jwt.SigningMethodHS512,
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

func (r *redisJWTHandler) GenJWTToken(uid int64) (string, error) {
	uc := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		},
		Uid: uid,
	}
	token := jwt.NewWithClaims(r.signingMethod, uc)
	tokenStr, err := token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
