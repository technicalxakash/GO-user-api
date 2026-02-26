package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/akash/go-user-api/internal/models"
)

// sendErrorResponse sends a standardized error response
func sendErrorResponse(c *fiber.Ctx, statusCode int, code, message string) error {
	requestID, ok := c.Locals("request_id").(string)
	if !ok {
		requestID = ""
	}

	errResp := models.ErrorResponse{}
	errResp.Error.Message = message
	errResp.Error.Code = code
	errResp.Error.RequestID = requestID

	return c.Status(statusCode).JSON(errResp)
}
