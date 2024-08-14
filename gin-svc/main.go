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
	"gin-svc/internal/web/middleware/jwt"
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
	userCtl := InitUserController(db, redisCli)
	pubController := InitPubController(db, redisCli)
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

func InitUserController(db *gorm.DB, cli redis.Cmdable) controller.BaseController {
	userDao := dao.NewUserDao(db)
	userRepo := repo.NewUserRepoInterface(userDao)

	codeCache := cache.NewRedisVerificationCodeCache(cli)
	jwtHdl := jwt.NewRedisJWTHandler(cli)
	userSvc := service.NewUserSvc(userRepo, codeCache, jwtHdl)
	return controller.NewUserController(userSvc)
}

func InitPubController(db *gorm.DB, cli redis.Cmdable) controller.BaseController {
	userDao := dao.NewUserDao(db)
	userRepo := repo.NewUserRepoInterface(userDao)
	codeCache := cache.NewRedisVerificationCodeCache(cli)
	jwtHdl := jwt.NewRedisJWTHandler(cli)
	userSvc := service.NewUserSvc(userRepo, codeCache, jwtHdl)
	redisVerificationCodeCache := cache.NewRedisVerificationCodeCache(cli)
	emailSvc := service.NewSMTPService(redisVerificationCodeCache)
	return controller.NewPublicController(userSvc, emailSvc)
}

func InitRoleController(db *gorm.DB, cli redis.Cmdable) controller.BaseController {
	perDao := dao.NewPermissionDAO(db)
	roleDao := dao.NewRoleDAO(db)
	cac := cache.NewRoleCache(cli)
	roleRepo := repo.NewRoleRepo(perDao, roleDao, cac)
	roleSvc := service.NewRoleService(roleRepo)
	return controller.NewRoleController(roleSvc)
}
