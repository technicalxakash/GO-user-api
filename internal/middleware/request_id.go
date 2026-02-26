package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/akash/go-user-api/internal/logger"
)

// RequestIDMiddleware generates and tracks request IDs
func RequestIDMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get request ID from header or generate new one
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Store in context
		c.Locals("request_id", requestID)

		// Set response header
		c.Set("X-Request-ID", requestID)

		// Log request with request ID
		logger.Log.Info("Incoming request",
			zap.String("request_id", requestID),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
		)

		return c.Next()
	}
}
