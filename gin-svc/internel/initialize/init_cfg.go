package initialize

import (
	"gin-svc/internel/conf"
	"github.com/spf13/viper"
)

func SetupCfg(path string) (conf.Config, error) {
	var cfg conf.Config
	// 读取配置文件
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {

		return cfg, err
	}
	// 将配置信息保存到全局变量
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func MustSetupCfg(path string) conf.Config {
	cfg, err := SetupCfg(path)
	if err != nil {
		panic(err)
	}
	return cfg
}
