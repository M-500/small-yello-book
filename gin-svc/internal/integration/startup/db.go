package startup

import (
	"gin-svc/internal/conf"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitRoleWebServer(path string) *gin.Engine {
	//config := ioc.SetUpConfig(path)

	return gin.Default()
}

func InitTestDB(cfg *conf.ConfigInstance) *gorm.DB {
	config := &gorm.Config{}
	db, err := gorm.Open(mysql.Open(cfg.Database.DSN), config)
	if err != nil {
		panic(err)
	}
	return db
}
