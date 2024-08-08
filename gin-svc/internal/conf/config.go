package conf

import (
	"gin-svc/internal/conf/cfg"
)

type ConfigInstance struct {
	Database cfg.DatabaseCfg `mapstructure:"database"`
}
