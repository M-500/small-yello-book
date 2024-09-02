package initialize

import (
	lg "gin-svc/pkg/logger"
	"go.uber.org/zap"
)

// 主要用来初始化日志

func SetUpLogger() lg.Logger {
	config := zap.NewDevelopmentConfig()
	logger, _ := config.Build()
	return lg.NewZapLogger(logger)
}
