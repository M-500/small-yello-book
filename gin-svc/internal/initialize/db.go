package initialize

import (
	"gin-svc/internal/conf/cfg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化数据库连接

func SetUpDB(cfg *cfg.DatabaseCfg) *gorm.DB {
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	//config.NamingStrategy = schema.NamingStrategy{
	//	TablePrefix:   "bbs_",
	//	SingularTable: cfg.ServiceDebug,
	//}
	db, err := gorm.Open(mysql.Open(cfg.DSN), config)
	if err != nil {
		panic(err)
	}
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	}
	//err = db.Use(prometheus.New(prometheus.Config{
	//	DBName:          cfg.ServiceName,
	//	RefreshInterval: 15,    // 插件采集数据的频率
	//	StartServer:     false, // 无需重新启动
	//	MetricsCollector: []prometheus.MetricsCollector{
	//		&prometheus.MySQL{
	//			VariableNames: []string{"thread_running"},
	//		},
	//	},
	//}))

	return db
}
