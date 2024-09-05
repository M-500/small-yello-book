package cfg

type ServerCfg struct {
	Addr           int    `mapstructure:"addr"`
	Mode           string `mapstructure:"mode"`
	FileUploadHost string `mapstructure:"file_upload_host"`
}
