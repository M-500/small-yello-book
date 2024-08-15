package middleware

import "github.com/gin-gonic/gin"

// 权限校验中间件

type PermissionMiddleware struct {
	path []string // 不需要权限校验的路径
}

func NewPermissionMdl() *PermissionMiddleware {
	return &PermissionMiddleware{}
}

func (m *PermissionMiddleware) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (m *PermissionMiddleware) IgnorePath(path string) *PermissionMiddleware {
	m.path = append(m.path, path)
	return m
}
