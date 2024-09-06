package jwt

import (
	"gin-svc/internal/conf"
	lg "gin-svc/pkg/ylog"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

type redisJWTHandler struct {
	client        redis.Cmdable
	signingMethod jwt.SigningMethod
	JWTKey        []byte
	Cfg           *conf.ConfigInstance
	Lg            lg.Logger
}

var _ Handler = &redisJWTHandler{} // 没别的意思，确保实现接口而已

func NewRedisJWTHandler(cli redis.Cmdable, cfg *conf.ConfigInstance) Handler {
	return &redisJWTHandler{
		client:        cli,
		signingMethod: jwt.SigningMethodHS512,
		JWTKey:        []byte(cfg.Jwt.Secret),
		Cfg:           cfg,
	}
}

func (r *redisJWTHandler) ClearToken(ctx *gin.Context) error {
	//TODO implement me
	panic("implement me")
}

// ExtractToken
//
//	@Description: 提取TOKEN 根据约定，token 在 Authorization 头部
//	@receiver r
//	@param ctx
//	@return string
func (r *redisJWTHandler) ExtractToken(ctx *gin.Context) string {
	authCode := ctx.GetHeader(r.Cfg.Jwt.JwtHeaderKey)
	if authCode == "" {
		return authCode
	}
	segs := strings.Split(authCode, " ")
	if len(segs) != 2 {
		return ""
	}
	return segs[1]
}

func (r *redisJWTHandler) SetLoginToken(ctx *gin.Context, uid string) error {
	//TODO implement me
	panic("implement me")
}

func (r *redisJWTHandler) setJwtToken(ctx *gin.Context, uid string, ssid string) error {
	uc := UserClaims{
		Uid: uid,
		//UserAgent: ctx.GetHeader("User-Agent"),
		RegisteredClaims: jwt.RegisteredClaims{
			// 1 分钟过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(r.Cfg.Jwt.Expire))),
		},
	}
	token := jwt.NewWithClaims(r.signingMethod, uc)
	tokenStr, err := token.SignedString([]byte(r.Cfg.Jwt.Secret))
	if err != nil {
		return err
	}
	ctx.Header("x-jwt-token", tokenStr)
	return nil
}

func (r *redisJWTHandler) getJwtToken(ctx *gin.Context, ssid string) error {
	//TODO implement me
	panic("implement me")
}

// GenJWTToken
//
//	@Description: 生成token
//	@receiver r
//	@param uid
//	@return string
//	@return error
func (r *redisJWTHandler) GenJWTToken(uid string) (string, error) {
	uc := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		},
		Uid: uid,
	}
	token := jwt.NewWithClaims(r.signingMethod, uc)
	tokenStr, err := token.SignedString([]byte(r.Cfg.Jwt.Secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
