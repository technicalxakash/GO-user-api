package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/akash/go-user-api/config"
	"github.com/akash/go-user-api/internal/handler"
	"github.com/akash/go-user-api/internal/logger"
	"github.com/akash/go-user-api/internal/middleware"
	"github.com/akash/go-user-api/internal/repository"
	"github.com/akash/go-user-api/internal/routes"
)

func main() {

	logger.Init()

	db := config.ConnectDB()

	app := fiber.New()

	app.Use(middleware.RequestMiddleware())

	repo := repository.NewUserRepo(db)
	handler := handler.NewUserHandler(repo)

	routes.SetupRoutes(app, handler)

	app.Listen(":3000")
}