package internal

import (
	"gin-svc/internal/conf"
	lg "gin-svc/pkg/ylog"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Engine      *gin.Engine
	DB          *gorm.DB
	Cli         redis.Cmdable
	Cfg         *conf.ConfigInstance
	Lg          lg.Logger
	MinioClient *minio.Client
}

func NewApp(engine *gin.Engine, DB *gorm.DB, cli redis.Cmdable, cfg *conf.ConfigInstance, client *minio.Client) *App {
	return &App{
		Engine:      engine,
		DB:          DB,
		Cli:         cli,
		Cfg:         cfg,
		MinioClient: client,
	}
}
