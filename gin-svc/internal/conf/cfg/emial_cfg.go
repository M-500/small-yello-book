package cfg

type EmailCfg struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Addr      string `mapstructure:"addr"`
	SecretKey string `mapstructure:"secret_key"`
}
