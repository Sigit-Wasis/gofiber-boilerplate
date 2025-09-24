package main

import (
	"log"

	_ "github.com/Sigit-Wasis/gofiber-boilerplate/docs"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"

	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/config"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/db"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/router"
)

// @title Go Fiber Postgres Boilerplate API
// @version 1.0
// @description Boilerplate API with Fiber, Postgres (sql), migrations, seeders, and Swagger
// @host localhost:8080
// @BasePath /
func main() {
cfg := config.Load()


// init db
if err := db.Init(cfg.DatabaseURL); err != nil {
log.Fatalf("failed to init db: %v", err)
}
defer db.Close()


app := fiber.New()


// routes
router.Setup(app)


// swagger route
app.Get("/swagger/*", swagger.New())


addr := ":" + cfg.AppPort
log.Printf("listening on %s", addr)
if err := app.Listen(addr); err != nil {
log.Fatalf("listen error: %v", err)
}
}