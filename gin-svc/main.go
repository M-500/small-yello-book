//@Author: wulinlin
//@Description: main文件启动
//@File:  main
//@Version: 1.0.0
//@Date: 2024/07/15 17:43

package main

import (
	"flag"
	"gin-svc/internal"
	"gin-svc/internal/initialize"
	"gin-svc/internal/models"
	"gin-svc/internal/web"
)

var configFile = flag.String("config", "etc/dev.yaml", "配置文件路径")

// @title Swagger YBOOK API
// @version 0.0.1
// @description Y_BOOK 的API文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	config := initialize.SetUpConfig(*configFile)
	db := initialize.SetUpDB(&config.Database)
	redisCli := initialize.SetUpRedis(&config.Redis)
	logger := initialize.SetUpLogger(config)
	err := models.InitTables(db)
	if err != nil {
		return
	}
	app := &internal.App{
		Engine: nil,
		DB:     db,
		Cli:    redisCli,
		Lg:     logger,
		Cfg:    config,
	}
	app.Engine = web.SetupWebEngine(app)
	err = app.Engine.Run(":8122")
	if err != nil {
		panic(err)
	}
}
