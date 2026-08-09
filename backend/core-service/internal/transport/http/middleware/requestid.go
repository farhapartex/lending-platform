package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

const (
	HeaderRequestID  = "X-Request-Id"
	ContextRequestID = "request_id"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(HeaderRequestID)
		if requestID == "" {
			requestID = generate()
		}

		c.Set(ContextRequestID, requestID)
		c.Header(HeaderRequestID, requestID)
		c.Next()
	}
}

func RequestIDFrom(c *gin.Context) string {
	if value, ok := c.Get(ContextRequestID); ok {
		if requestID, ok := value.(string); ok {
			return requestID
		}
	}

	return ""
}

func generate() string {
	buf := make([]byte, 12)
	if _, err := rand.Read(buf); err != nil {
		return "unknown"
	}

	return hex.EncodeToString(buf)
}
