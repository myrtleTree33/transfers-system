package internal

import "github.com/gin-gonic/gin"

// Retrieve trace ID from Gin context
func GetTraceId(c *gin.Context) string {
	traceId, _ := c.Get("X-Trace-Id")
	return traceId.(string)
}
