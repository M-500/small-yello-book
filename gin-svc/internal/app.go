package internal

import (
	"gin-svc/internal/conf"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Engine *gin.Engine
	DB     *gorm.DB
	Cli    redis.Cmdable
	Cfg    *conf.ConfigInstance
}

func NewApp(engine *gin.Engine, DB *gorm.DB, cli redis.Cmdable, cfg *conf.ConfigInstance) *App {
	return &App{
		Engine: engine,
		DB:     DB,
		Cli:    cli,
		Cfg:    cfg}
}
