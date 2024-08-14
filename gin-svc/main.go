//@Author: wulinlin
//@Description: main文件启动
//@File:  main
//@Version: 1.0.0
//@Date: 2024/07/15 17:43

package main

import (
	"flag"
	"gin-svc/internal"
	"gin-svc/internal/ioc"
	"gin-svc/internal/models"
	"gin-svc/internal/web"
)

var configFile = flag.String("config", "etc/dev.yaml", "配置文件路径")

// @title Swagger YBOOK API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	config := ioc.SetUpConfig(*configFile)
	db := ioc.SetUpDB(&config.Database)
	redisCli := ioc.SetUpRedis(&config.Redis)
	err := models.InitTables(db)
	if err != nil {
		return
	}
	app := &internal.App{
		Engine: nil,
		DB:     db,
		Cli:    redisCli,
		Cfg:    config,
	}
	app.Engine = web.SetupWebEngine(app)
	err = app.Engine.Run(":8122")
	if err != nil {
		panic(err)
	}
}
