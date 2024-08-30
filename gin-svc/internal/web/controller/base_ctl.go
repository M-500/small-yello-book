package controller

import "github.com/gin-gonic/gin"

type BaseController interface {
	RegisterRoute(group *gin.RouterGroup)
}

// 用于在编译期间确保 实现了BaseController接口
var _ BaseController = (*UserController)(nil)
