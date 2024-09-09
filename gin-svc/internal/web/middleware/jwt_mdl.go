package middleware

import (
	"gin-svc/internal/conf/cfg"
	myJwt "gin-svc/internal/web/middleware/jwt"
	"gin-svc/pkg/constant"
	"gin-svc/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// JWT 权限校验中间件 采用Builder模式

type JWTMiddlewareBuilder struct {
	myJwt.Handler
	cfg  cfg.JwtCfg
	path []string // 不需要权限校验的路径
}

func NewJwtMiddlewareBuilder(hdl myJwt.Handler, cfg cfg.JwtCfg) *JWTMiddlewareBuilder {
	return &JWTMiddlewareBuilder{
		Handler: hdl,
		cfg:     cfg,
	}
}

func (j *JWTMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.URL.Path
		// 判断uri是否在path中
		for _, p := range j.path {
			if p == uri || utils.KeyMatch2(uri, p) {
				ctx.Next() // 直接放行
				return
			}
		}
		tokenStr := j.ExtractToken(ctx)
		var uc myJwt.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return []byte(j.cfg.Secret), nil
		})
		if err != nil {
			// token 不对，token 是伪造的
			//ctx.AbortWithStatus(401)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token校验失败",
			})
			ctx.Abort()
			return
		}
		if token == nil || !token.Valid {
			// 在这里发现 access_token 过期了，生成一个新的 access_token

			// token 解析出来了，但是 token 可能是非法的，或者过期了的
			//ctx.AbortWithStatus(http.StatusUnauthorized)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token过期",
			})
			ctx.Abort()
			return
		}

		ctx.Set(constant.CtxUserId, uc.Uid)
		ctx.Set(constant.CtxUserKey, uc)
		ctx.Next()
	}
}

func (j *JWTMiddlewareBuilder) IgnorePath(path string) *JWTMiddlewareBuilder {
	j.path = append(j.path, path)
	return j
}
