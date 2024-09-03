package initialize

import (
	"gin-svc/internal/conf"
	lg "gin-svc/pkg/ylog"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// 主要用来初始化日志

func SetUpLogger(cfg *conf.ConfigInstance) lg.Logger {
	isProduction := !cfg.Logs.IsDev

	errorLogWriter := getLogWriter(cfg.Logs.ErrFileName, cfg.Logs.MaxSize, cfg.Logs.MaxBackups, cfg.Logs.MaxAge, cfg.Logs.Compress)
	infoLogWriter := getLogWriter(cfg.Logs.FileName, cfg.Logs.MaxSize, cfg.Logs.MaxBackups, cfg.Logs.MaxAge, cfg.Logs.Compress)

	// 配置不同级别的核心
	errorCore := zapcore.NewCore(getEncoder(isProduction), errorLogWriter, zapcore.ErrorLevel)
	infoCore := zapcore.NewCore(getEncoder(isProduction), infoLogWriter, zapcore.InfoLevel)
	// 在开发环境下，额外创建一个 consoleCore，将日志输出到控制台
	if !isProduction {
		consoleCore := zapcore.NewCore(getEncoder(isProduction), zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
		logger := zap.New(zapcore.NewTee(infoCore, errorCore, consoleCore), zap.AddCaller())
		defer logger.Sync()
		// 为 Error 及以上级别的日志自动添加堆栈信息
		logger = logger.WithOptions(zap.AddStacktrace(zapcore.ErrorLevel))
		return lg.NewZapLogger(logger)
	}

	// 生产环境只写入文件，且自动添加堆栈信息
	logger := zap.New(zapcore.NewTee(infoCore, errorCore), zap.AddCaller())
	defer logger.Sync()

	// 为 Error 及以上级别的日志自动添加堆栈信息
	logger = logger.WithOptions(zap.AddStacktrace(zapcore.ErrorLevel))
	return lg.NewZapLogger(logger)
}

// getEncoder 根据环境返回不同的日志编码器
func getEncoder(isProduction bool) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",       // 时间键名
		LevelKey:       "level",      // 级别键名
		NameKey:        "logger",     // 日志器键名
		CallerKey:      "caller",     // 调用者键名
		MessageKey:     "msg",        // 消息键名
		StacktraceKey:  "stacktrace", // 堆栈键名
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,                // 级别大写
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.DateTime), // ISO 8601时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 简单文件路径
	}

	if isProduction {
		return zapcore.NewJSONEncoder(encoderConfig) // 生产环境使用JSON编码
	}
	return zapcore.NewConsoleEncoder(encoderConfig) // 开发环境使用文本编码
}

// getLogWriter 根据日志级别返回不同的日志写入器
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,   // 每个日志文件最大10MB
		MaxBackups: maxBackup, // 最多保留5个备份日志文件
		MaxAge:     maxAge,    // 最长保存30天
		Compress:   compress,  // 是否压缩备份日志文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
