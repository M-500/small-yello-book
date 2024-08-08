package web

import (
	"gin-svc/internal/web/controller"
	"github.com/gin-gonic/gin"
)

func SetupWebEngine() *gin.Engine {
	engine := gin.Default()
	rg := engine.Group("/api/v1")
	{
		controller.NewUserController().RegisterRoute(rg)
	}

	return engine
}
