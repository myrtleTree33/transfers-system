package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// WithTraceId is a middleware that adds a trace id to the request
func WithTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("X-Trace-Id")
		if traceId == "" || len(traceId) != 36 {
			traceId = uuid.New().String()
		}
		c.Set("X-Trace-Id", traceId)
		c.Writer.Header().Add("X-Trace-Id", traceId)
		c.Next()
	}
}
