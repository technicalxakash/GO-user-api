package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/akash/go-user-api/internal/logger"
	"github.com/akash/go-user-api/internal/models"
	"github.com/akash/go-user-api/internal/service"
)

// AuthMiddleware validates JWT token
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID, _ := c.Locals("request_id").(string)

		// Get token from Authorization header or cookie
		authHeader := c.Get("Authorization")
		var token string

		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				token = parts[1]
			}
		}

		// Try to get from cookie if not in header
		if token == "" {
			token = c.Cookies("token")
		}

		if token == "" {
			logger.Log.Warn("Missing authentication token",
				zap.String("request_id", requestID),
				zap.String("path", c.Path()))
			return sendAuthError(c, 401, "NO_TOKEN", "Missing authentication token")
		}

		// Verify token
		claims, err := service.VerifyToken(token)
		if err != nil {
			logger.Log.Warn("Invalid token",
				zap.Error(err),
				zap.String("request_id", requestID))
			return sendAuthError(c, 401, "INVALID_TOKEN", "Invalid or expired token")
		}

		// Store user info in context
		c.Locals("user_id", claims.UserID)
		c.Locals("user_email", claims.Email)
		c.Locals("user_role", claims.Role)

		return c.Next()
	}
}

// sendAuthError sends authentication error response
func sendAuthError(c *fiber.Ctx, statusCode int, code, message string) error {
	requestID, _ := c.Locals("request_id").(string)
	errResp := models.ErrorResponse{}
	errResp.Error.Message = message
	errResp.Error.Code = code
	errResp.Error.RequestID = requestID

	return c.Status(statusCode).JSON(errResp)
}
