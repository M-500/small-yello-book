package main

import (
	"flag"
	"fmt"
	"gin-svc/internel/initialize"
	"gin-svc/internel/web"
	"github.com/gin-gonic/gin"
)

var configFile = flag.String("config", "etc/local.yaml", "配置文件路径")

func main() {

	cfg := initialize.MustSetupCfg(*configFile)
	engine := gin.Default()
	web.RegisterRouters(engine) // 注册路由
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	engine.Run(addr)
}
