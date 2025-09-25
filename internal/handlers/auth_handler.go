package handlers

import (
	"database/sql"
	"time"

	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/models"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/repository"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("SUPER_SECRET_KEY") // TODO: ambil dari ENV

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
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var body models.RegisterRequest
	if err := c.BodyParser(&body); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to hash password", err.Error())
	}

	user := models.User{
		Name:         body.Name,
		Email:        body.Email,
		PasswordHash: string(hash),
	}

	if err := h.Repo.Create(c.Context(), &user); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
	}

	return utils.Success(c, fiber.StatusCreated, "User created successfully", fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and get JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body models.LoginRequest true "Login request"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body models.LoginRequest
	if err := c.BodyParser(&body); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	user, err := h.Repo.GetByEmail(c.Context(), body.Email)
	if err != nil || user == nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid credentials", "email not found")
	}

	// Compare password
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password)) != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid credentials", "wrong password")
	}

	// Generate JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(jwtSecret)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to generate token", err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "Login successful", fiber.Map{
		"token": signed,
	})
}
