package web

import (
	_ "gin-svc/docs"
	"gin-svc/internal/web/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupWebEngine(userCtl, pubCtl controller.BaseController) *gin.Engine {
	engine := gin.Default()
	rg := engine.Group("")

	publicGroup := rg.Group("/na")
	{
		publicGroup.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		pubCtl.RegisterRoute(publicGroup)
	}

	privateGroup := rg.Group("/api/v1")
	{
		userCtl.RegisterRoute(privateGroup)
	}

	return engine
}
