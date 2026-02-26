package routes

import (
	"github.com/akash/go-user-api/internal/handler"
	"github.com/akash/go-user-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handler.UserHandler, ah *handler.AuthHandler) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "User API is running"})
	})

	// Auth routes (public)
	auth := app.Group("/auth")
	auth.Post("/signup", ah.Signup)
	auth.Post("/login", ah.Login)

	// Protected user routes
	api := app.Group("/users")
	api.Post("/", h.CreateUser)
	api.Get("/", h.ListUsers)
	api.Get("/:id", h.GetUser)
	api.Put("/:id", h.UpdateUser)
	api.Delete("/:id", h.DeleteUser)

	// Protected profile route
	profile := app.Group("/user")
	profile.Use(middleware.AuthMiddleware())
	profile.Get("/profile", ah.GetProfile)

	// Admin routes
	admin := app.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminOnly())
	// Add admin-specific routes here
}
