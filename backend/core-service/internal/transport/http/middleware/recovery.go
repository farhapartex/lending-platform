package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"github.com/farhapartex/lending-platform/core-service/internal/transport/http/dto"
)

func Recovery(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				log.Error(
					"panic recovered",
					slog.String("request_id", RequestIDFrom(c)),
					slog.String("method", c.Request.Method),
					slog.String("path", c.Request.URL.Path),
					slog.Any("panic", recovered),
					slog.String("stack", string(debug.Stack())),
				)

				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					dto.NewErrorResponse(dto.CodeInternalError, "Something went wrong on our side.", RequestIDFrom(c)),
				)
			}
		}()

		c.Next()
	}
}
