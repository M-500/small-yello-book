package conf

import (
	"gin-svc/internal/conf/cfg"
)

type ConfigInstance struct {
	Database cfg.DatabaseCfg `mapstructure:"database"`
	Redis    cfg.RedisCfg    `mapstructure:"redis"`
	Email    cfg.EmailCfg    `mapstructure:"email"`
	Logs     cfg.LogCfg      `mapstructure:"logs"`
	Server   cfg.ServerCfg   `mapstructure:"server"`
	Jwt      cfg.JwtCfg      `mapstructure:"jwt"`
	Minio cfg.MinioCfg `mapstructure:"minio"`
}
