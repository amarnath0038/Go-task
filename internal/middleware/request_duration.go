package middleware

import (
	"time"

	"user-crud/internal/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestDuration() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		// fetch req id set by previous middleware
		requestID, _ := c.Locals(RequestIDKey).(string)

		logger.Log.Info("request completed",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
			zap.String("request_id", requestID),
		)
		return err
	}
}
