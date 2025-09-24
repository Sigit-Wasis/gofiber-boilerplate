package router

import (
	"database/sql"

	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// User CRUD
	userHandler := handlers.NewUserHandler(db)
	app.Get("/users", userHandler.GetUsers)
	app.Get("/users/:id", userHandler.GetUser)
	app.Post("/users", userHandler.CreateUser)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
}