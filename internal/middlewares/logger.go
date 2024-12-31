package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// logs details of each incoming request and response
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start time
		startTime := time.Now()

		// Process request
		c.Next()

		// End time
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Get status code
		statusCode := c.Writer.Status()

		// Get client IP address
		clientIP := c.ClientIP()

		// Log details about the request
		if statusCode >= 200 && statusCode < 300 {
			logger.Info("Request successful",
				zap.String("method", c.Request.Method),
				zap.String("url", c.Request.URL.Path),
				zap.String("client_ip", clientIP),
				zap.Int("status_code", statusCode),
				zap.Duration("duration", latency),
			)
		} else if statusCode >= 400 && statusCode < 500 {
			logger.Warn("Client error",
				zap.String("method", c.Request.Method),
				zap.String("url", c.Request.URL.Path),
				zap.String("client_ip", clientIP),
				zap.Int("status_code", statusCode),
				zap.Duration("duration", latency),
			)
		} else if statusCode >= 500 {
			logger.Error("Server error",
				zap.String("method", c.Request.Method),
				zap.String("url", c.Request.URL.Path),
				zap.String("client_ip", clientIP),
				zap.Int("status_code", statusCode),
				zap.Duration("duration", latency),
			)
		} else {
			logger.Info("Logging request",
				zap.String("method", c.Request.Method),
				zap.String("url", c.Request.URL.Path),
				zap.String("client_ip", clientIP),
				zap.Int("status_code", statusCode),
				zap.Duration("duration", latency),
			)
		}
	}
}
