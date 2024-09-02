package web

import (
	"gin-svc/internel/web/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	helloCtl := controller.NewHelloCtl()
	r.GET("/hello", helloCtl.Hello)
}
