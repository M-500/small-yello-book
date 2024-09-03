package cfg

type LogCfg struct {
	IsDev       bool   `mapstructure:"is_dev"`
	FileName    string `mapstructure:"file_name"`
	ErrFileName string `mapstructure:"err_file_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}
