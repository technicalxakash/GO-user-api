package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/akash/go-user-api/internal/logger"
	"github.com/akash/go-user-api/internal/models"
)

// RoleMiddleware checks if user has required role
func RoleMiddleware(requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID, _ := c.Locals("request_id").(string)
		userRole, ok := c.Locals("user_role").(string)

		if !ok {
			logger.Log.Warn("Missing user role in context",
				zap.String("request_id", requestID))
			return sendRBACError(c, 401, "UNAUTHORIZED", "User not authenticated")
		}

		// Check if user role is allowed
		allowed := false
		for _, role := range requiredRoles {
			if userRole == role {
				allowed = true
				break
			}
		}

		if !allowed {
			logger.Log.Warn("Unauthorized role access",
				zap.String("request_id", requestID),
				zap.String("user_role", userRole),
				zap.Strings("required_roles", requiredRoles),
				zap.String("path", c.Path()))
			return sendRBACError(c, 403, "FORBIDDEN", "You do not have permission to access this resource")
		}

		logger.Log.Info("Role authorization success",
			zap.String("request_id", requestID),
			zap.String("user_role", userRole))

		return c.Next()
	}
}

// AdminOnly is a shortcut for RoleMiddleware with admin role
func AdminOnly() fiber.Handler {
	return RoleMiddleware("admin")
}

// sendRBACError sends role-based access control error
func sendRBACError(c *fiber.Ctx, statusCode int, code, message string) error {
	requestID, _ := c.Locals("request_id").(string)
	errResp := models.ErrorResponse{}
	errResp.Error.Message = message
	errResp.Error.Code = code
	errResp.Error.RequestID = requestID

	return c.Status(statusCode).JSON(errResp)
}
