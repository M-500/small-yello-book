package middleware

import (
	myJwt "gin-svc/internal/web/middleware/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// JWT 权限校验中间件 采用Builder模式

type JWTMiddlewareBuilder struct {
	myJwt.Handler
	path []string // 不需要权限校验的路径
}

func NewJwtMiddlewareBuilder(hdl myJwt.Handler) *JWTMiddlewareBuilder {
	return &JWTMiddlewareBuilder{
		Handler: hdl,
	}
}

func (j *JWTMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.URL.Path
		// 判断uri是否在path中
		for _, p := range j.path {
			if p == uri {
				ctx.Next() // 直接放行
				return
			}
		}
		tokenStr := j.ExtractToken(ctx)
		var uc myJwt.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return myJwt.JWTKey, nil
		})
		if err != nil {
			// token 不对，token 是伪造的
			ctx.AbortWithStatus(401)
			return
		}
		if token == nil || !token.Valid {
			// 在这里发现 access_token 过期了，生成一个新的 access_token

			// token 解析出来了，但是 token 可能是非法的，或者过期了的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("uid", uc.Uid)
		ctx.Next()
	}
}

func (j *JWTMiddlewareBuilder) IgnorePath(path string) *JWTMiddlewareBuilder {
	j.path = append(j.path, path)
	return j
}
