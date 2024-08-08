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
)

var configFile = flag.String("config", "etc/dev.yaml", "配置文件路径")

func main() {
	config := ioc.SetUpConfig(*configFile)
	fmt.Println(config)
	db := ioc.SetUpDB(&config.Database)
	err := models.InitTables(db)
	if err != nil {
		return
	}
}
