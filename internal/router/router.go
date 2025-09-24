package router

import (
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/handlers"
	"github.com/gofiber/fiber/v2"
)


func Setup(app *fiber.App) {
api := app.Group("/api")


api.Get("/health", handlers.Health)
}