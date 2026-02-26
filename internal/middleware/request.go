package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestMiddleware() fiber.Handler {

	return func(c *fiber.Ctx) error {

		start := time.Now()

		requestID := uuid.New().String()
		c.Set("X-Request-ID", requestID)

		err := c.Next()

		duration := time.Since(start)
		println("Request:", c.Path(), duration.String())

		return err
	}
}