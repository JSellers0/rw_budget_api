package middleware

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupLogger creates log files and returns loggers
func SetupLogger(logDir string) (io.Writer, io.Writer, error) {
	// Create log directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	// Create access log file
	accessLogPath := filepath.Join(logDir, "access.log")
	accessLog, err := os.OpenFile(accessLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open access log: %w", err)
	}

	// Create error log file
	errorLogPath := filepath.Join(logDir, "error.log")
	errorLog, err := os.OpenFile(errorLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		accessLog.Close()
		return nil, nil, fmt.Errorf("failed to open error log: %w", err)
	}

	return accessLog, errorLog, nil
}

// AccessLogger logs all requests
func AccessLogger(writer io.Writer) gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Output: writer,
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[ACCESS] %s | %3d | %13v | %15s | %-7s %s\n",
				param.TimeStamp.Format(time.RFC3339),
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				param.Method,
				param.Path,
			)
		},
	})
}

// ErrorLogger logs errors (4xx, 5xx responses)
func ErrorLogger(writer io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Log if there's an error or status code >= 400
		if len(c.Errors) > 0 || c.Writer.Status() >= 400 {
			timestamp := time.Now().Format(time.RFC3339)
			logEntry := fmt.Sprintf("[ERROR] %s | %3d | %15s | %-7s %s | Errors: %v\n",
				timestamp,
				c.Writer.Status(),
				c.ClientIP(),
				c.Request.Method,
				c.Request.URL.Path,
				c.Errors.String(),
			)
			writer.Write([]byte(logEntry))
		}
	}
}
