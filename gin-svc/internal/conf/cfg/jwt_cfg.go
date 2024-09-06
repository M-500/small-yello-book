package cfg

type JwtCfg struct {
	Secret       string `mapstructure:"secret"`
	Expire       int    `mapstructure:"expire"`
	JwtHeaderKey string `mapstructure:"jwt_header_key"`
	JwtPrefix    string `mapstructure:"jwt_prefix"`
}
