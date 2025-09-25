package handlers

import (
	"database/sql"
	"time"

	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/models"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("SUPER_SECRET_KEY") // sebaiknya dari ENV

type AuthHandler struct {
	Repo *repository.UserRepository
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{Repo: repository.NewUserRepository(db)}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param body body models.RegisterRequest true "Register request"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var body models.RegisterRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:         body.Name,
		Email:        body.Email,
		PasswordHash: string(hash),
	}

	if err := h.Repo.Create(c.Context(), &user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "user created"})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and get JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body models.LoginRequest true "Login request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body models.LoginRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.Repo.GetByEmail(c.Context(), body.Email)
	if err != nil || user == nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials Email"})
	}

	// Compare password
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password)) != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials Password"})
	}

	// Generate JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString(jwtSecret)

	return c.JSON(fiber.Map{"token": signed})
}
