package main

import (
	"log"

	_ "github.com/Sigit-Wasis/gofiber-boilerplate/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/config"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/db"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/router"
)

// @title Go Fiber Postgres Boilerplate API
// @version 1.0
// @description Boilerplate API with Fiber, Postgres (sql), migrations, seeders, and Swagger
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer {your JWT token}" to authenticate requests
func main() {
	cfg := config.Load()

	// init db
	if err := db.Init(cfg.DatabaseURL); err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer db.Close()

	app := fiber.New()

	// enable CORS for Swagger UI
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // ubah ke origin spesifik kalau mau lebih aman
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// serve docs as static
	app.Static("/docs", "./docs")

	// swagger UI
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// routes
	router.SetupRoutes(app, db.GetDB())

	addr := ":" + cfg.AppPort
	log.Printf("listening on %s", addr)

	if err := app.Listen(addr); err != nil {
		log.Fatalf("listen error: %v", err)
	}
}
