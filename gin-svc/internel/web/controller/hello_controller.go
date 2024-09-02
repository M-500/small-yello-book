package controller

import "github.com/gin-gonic/gin"

type HelloCtl struct {
}

func NewHelloCtl() *HelloCtl {
	return &HelloCtl{}
}

func (h *HelloCtl) Hello(ctx *gin.Context) {
	ctx.JSON(200, "Hello World")
}
