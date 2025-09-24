package handlers

import (
	"database/sql"
	"strconv"

	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/models"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{Repo: repository.NewUserRepository(db)}
}

// GetUsers godoc
// @Summary Get all users
// @Description Retrieve list of users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.Repo.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Retrieve a single user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.Repo.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if user == nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)
}

// CreateUser godoc
// @Summary Create new user
// @Description Add a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User Data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var u models.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.Repo.Create(c.Context(), &u); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(u)
}

// UpdateUser godoc
// @Summary Update existing user
// @Description Update user data by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.User true "User Data"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var u models.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	u.ID = id
	if err := h.Repo.Update(c.Context(), &u); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(u)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Remove user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 204
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.Repo.Delete(c.Context(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
