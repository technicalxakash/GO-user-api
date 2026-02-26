package main

import (
	"log"

	"github.com/akash/go-user-api/config"
	"github.com/akash/go-user-api/internal/handler"
	"github.com/akash/go-user-api/internal/logger"
	"github.com/akash/go-user-api/internal/middleware"
	"github.com/akash/go-user-api/internal/repository"
	"github.com/akash/go-user-api/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	logger.Init()
	println(" Logger initialized")

	db := config.ConnectDB()
	println(" Database connected")

	app := fiber.New()
	println(" Fiber app created")

	// Global middleware
	app.Use(middleware.RequestIDMiddleware())
	app.Use(middleware.RequestMiddleware())
	println(" Global middleware added")

	// Initialize repository and handlers
	repo := repository.NewUserRepo(db)
	println(" User repo initialized")

	userHandler := handler.NewUserHandler(repo)
	println(" User handler initialized")

	authHandler := handler.NewAuthHandler(repo)
	println(" Auth handler initialized")

	// Setup routes
	routes.SetupRoutes(app, userHandler, authHandler)
	println(" Routes setup complete")

	println("🚀 Server listening on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(" Server failed:", err)
	}
}
