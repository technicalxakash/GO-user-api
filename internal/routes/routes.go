package routes

import (
	"github.com/akash/go-user-api/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handler.UserHandler) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "User API is running"})
	})

	api := app.Group("/users")

	api.Post("/", h.CreateUser)
	api.Get("/", h.ListUsers)
	api.Get("/:id", h.GetUser)
	api.Put("/:id", h.UpdateUser)
	api.Delete("/:id", h.DeleteUser)
}
