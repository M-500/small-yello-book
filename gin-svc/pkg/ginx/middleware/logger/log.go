package logger

import (
	"gin-svc/pkg/ylog"
	"github.com/gin-gonic/gin"
	"time"
)

func LogHdl(lg ylog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		lg.Info(path,
			ylog.Int("status", status),
			ylog.String("method", c.Request.Method),
			ylog.String("path", path),
			ylog.String("query", raw),
			ylog.String("ip", c.ClientIP()),
			ylog.String("user-agent", c.Request.UserAgent()),
			ylog.Duration("latency", latency),
			ylog.String("error", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}
}
