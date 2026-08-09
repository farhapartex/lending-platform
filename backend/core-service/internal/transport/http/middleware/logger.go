package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		route := c.FullPath()
		if route == "" {
			route = "unmatched"
		}

		attrs := []any{
			slog.String("request_id", RequestIDFrom(c)),
			slog.String("method", c.Request.Method),
			slog.String("route", route),
			slog.Int("status", c.Writer.Status()),
			slog.Duration("duration", time.Since(start)),
			slog.String("client_ip", c.ClientIP()),
		}

		if len(c.Errors) > 0 {
			attrs = append(attrs, slog.String("errors", c.Errors.String()))
		}

		switch {
		case c.Writer.Status() >= 500:
			log.Error("request failed", attrs...)
		case c.Writer.Status() >= 400:
			log.Warn("request rejected", attrs...)
		default:
			log.Info("request completed", attrs...)
		}
	}
}
