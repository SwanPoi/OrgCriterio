package handler

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/SwanPoi/OrgCriterio/internal/logger"
)

func LogMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Info("HTTP Request",
			"status", status,
			"method", method,
			"path", path,
			"latency", latency,
		)
	}
}