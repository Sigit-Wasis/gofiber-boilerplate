package middleware

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func generateToken(t *testing.T, claims jwt.MapClaims) string {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    s, err := token.SignedString(JWT_SECRET)
    assert.NoError(t, err)
    return s
}

func TestProtected_MissingToken(t *testing.T) {
    app := fiber.New()
    app.Use(Protected())
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("ok")
    })

    req := httptest.NewRequest("GET", "/", nil)
    resp, _ := app.Test(req)
    assert.NotNil(t, resp)
    assert.Equal(t, 401, resp.StatusCode)
}

func TestProtected_InvalidToken(t *testing.T) {
    app := fiber.New()
    app.Use(Protected())
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("ok")
    })

    req := httptest.NewRequest("GET", "/", nil)
    req.Header.Set("Authorization", "Bearer invalidtoken")
    resp, _ := app.Test(req)
    assert.Equal(t, 401, resp.StatusCode)
}

func TestProtected_ValidToken(t *testing.T) {
    app := fiber.New()
    app.Use(Protected())
    app.Get("/", func(c *fiber.Ctx) error {
        userID := c.Locals("user_id")
        return c.JSON(fiber.Map{"user_id": userID})
    })

    claims := jwt.MapClaims{
        "user_id": "123",
        "exp":     time.Now().Add(time.Hour).Unix(),
    }
    token := generateToken(t, claims)

    req := httptest.NewRequest("GET", "/", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    resp, _ := app.Test(req)
    assert.Equal(t, 200, resp.StatusCode)
}

func TestProtected_InvalidClaims(t *testing.T) {
    app := fiber.New()
    app.Use(Protected())
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("ok")
    })

    // Create a token with jwt.RegisteredClaims to simulate an invalid claim type for middleware expecting jwt.MapClaims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
    })
    if s, err := token.SignedString(JWT_SECRET); err != nil {
        t.Fatalf("failed to sign token: %v", err)
    } else {
        req := httptest.NewRequest("GET", "/", nil)
        req.Header.Set("Authorization", "bearer "+s)
        resp, err := app.Test(req)
        if err != nil {
            t.Fatalf("failed to test app: %v", err)
        } else if resp.StatusCode != 401 {
            t.Fatalf("expected 401 status code, got %d", resp.StatusCode)
        }
    }
}
