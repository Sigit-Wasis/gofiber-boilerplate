package handlers

import (
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/db"
	"github.com/gofiber/fiber/v2"
)

// Health godoc
// @Summary Health check
// @Description get health status
// @Tags health
// @Success 200 {object} map[string]interface{}
// @Router /api/health [get]
func Health(c *fiber.Ctx) error {
	d := db.GetDB()
	
	if d == nil {
		return c.Status(500).JSON(fiber.Map{"status": "db not initialized"})
	}

	if err := d.Ping(); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "db error", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "ok"})
}