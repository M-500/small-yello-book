package initialize

import (
	"gin-svc/internal/conf"
	"github.com/spf13/viper"
	"log"
)

func SetUpConfig(path string) *conf.ConfigInstance {
	var appCfg conf.ConfigInstance
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("ebip: Failed to read config file: %v", err)
	}

	err = viper.Unmarshal(&appCfg)
	if err != nil {
		log.Fatalf("ebip: Failed to unmarshal config file: %v", err)
	}
	return &appCfg
}
