package conf

type Config struct {
	Server ServerCfg `mapstructure:"server" json:"server"`
}

type ServerCfg struct {
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
