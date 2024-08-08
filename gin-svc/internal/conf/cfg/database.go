package cfg

type DatabaseCfg struct {
	DSN         string `mapstructure:"dsn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
}
