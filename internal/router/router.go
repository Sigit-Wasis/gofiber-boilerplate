package router

import (
	"database/sql"

	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/handlers"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/middleware"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	// Auth routes (public)
	authHandler := handlers.NewAuthHandler(db)
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	// Protected routes group
	api := app.Group("/api", middleware.Protected())

	// User CRUD (hanya bisa diakses dengan token JWT)
	userHandler := handlers.NewUserHandler(db)
	api.Get("/users", userHandler.GetUsers)
	api.Get("/users/:id", userHandler.GetUser)
	api.Post("/users", userHandler.CreateUser)
	api.Put("/users/:id", userHandler.UpdateUser)
	api.Delete("/users/:id", userHandler.DeleteUser)

	// Example protected profile
	api.Get("/profile", func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		return utils.Success(c, fiber.StatusOK, "Welcome to your profile", fiber.Map{
			"user_id": userID,
		})
	})
}
