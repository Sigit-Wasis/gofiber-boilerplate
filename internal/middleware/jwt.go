package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte("SUPER_SECRET_KEY")

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return c.Status(401).JSON(fiber.Map{"error": "missing token"})
		}

		const bearerPrefix = "Bearer "
		if len(auth) > len(bearerPrefix) && auth[:len(bearerPrefix)] == bearerPrefix {
			auth = auth[len(bearerPrefix):]
		}

		token, err := jwt.Parse(auth, func(token *jwt.Token) (any, error) {
			return JWT_SECRET, nil
		})
		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "invalid token claims"})
		}

		if userID, ok := claims["user_id"]; ok {
			c.Locals("user_id", userID)
		}

		return c.Next()
	}
}