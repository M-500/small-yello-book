//@Author: wulinlin
//@Description: main文件启动
//@File:  main
//@Version: 1.0.0
//@Date: 2024/07/15 17:43

package main

import (
	"flag"
	"fmt"
	"gin-svc/internal/ioc"
	"gin-svc/internal/models"
	"gin-svc/internal/repo"
	"gin-svc/internal/repo/cache"
	"gin-svc/internal/repo/dao"
	"gin-svc/internal/service"
	"gin-svc/internal/web"
	"gin-svc/internal/web/controller"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var configFile = flag.String("config", "etc/dev.yaml", "配置文件路径")

func main() {
	config := ioc.SetUpConfig(*configFile)
	fmt.Println(config)
	db := ioc.SetUpDB(&config.Database)
	redisCli := ioc.SetUpRedis(&config.Redis)
	err := models.InitTables(db)
	if err != nil {
		return
	}
	userCtl := InitUserController(db)
	pubController := InitPubController(db)
	roleCtl := InitRoleController(db, redisCli)
	engine := web.SetupWebEngine(
		userCtl,
		pubController,
		roleCtl)
	err = engine.Run(":8122")
	if err != nil {
		panic(err)
	}
}

func InitUserController(db *gorm.DB) controller.BaseController {
	userDao := dao.NewUserDao(db)
	userRepo := repo.NewUserRepoInterface(userDao)
	userSvc := service.NewUserSvc(userRepo)
	return controller.NewUserController(userSvc)
}

func InitPubController(db *gorm.DB) controller.BaseController {
	userDao := dao.NewUserDao(db)
	userRepo := repo.NewUserRepoInterface(userDao)
	userSvc := service.NewUserSvc(userRepo)
	return controller.NewPublicController(userSvc)
}

func InitRoleController(db *gorm.DB, cli redis.Cmdable) controller.BaseController {
	perDao := dao.NewPermissionDAO(db)
	roleDao := dao.NewRoleDAO(db)
	cac := cache.NewRoleCache(cli)
	roleRepo := repo.NewRoleRepo(perDao, roleDao, cac)
	roleSvc := service.NewRoleService(roleRepo)
	return controller.NewRoleController(roleSvc)
}
