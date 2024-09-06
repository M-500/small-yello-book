package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
)

// 跨域中间件

func CorsMdl() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // 允许的前端来源
		//AllowOrigins:     []string{"http://localhost:5173"},                   // 允许前端的地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},          // 允许的响应头
		AllowCredentials: true,                                                // 允许客户端发送Cookie等凭证信息
		MaxAge:           12 * time.Hour,                                      // 预检请求的缓存时间
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				//if strings.Contains(origin, "localhost") {
				return true
			}
			if strings.Contains(origin, ":1") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
	})
}
