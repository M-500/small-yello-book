package web

import (
	"gin-svc/internal/web/controller"
	"github.com/gin-gonic/gin"
)

func SetupWebEngine(userCtl controller.BaseController) *gin.Engine {
	engine := gin.Default()
	rg := engine.Group("/api/v1")
	{
		userCtl.RegisterRoute(rg)
	}

	return engine
}
