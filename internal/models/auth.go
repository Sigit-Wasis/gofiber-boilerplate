// internal/models/auth.go
package models

// RegisterRequest is the payload for registering a new user
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest is the payload for logging in a user
type LoginRequest struct {
	Email    string `json:"email" db:"email" example:"admin@example.com"`
	Password string `json:"password" db:"password_hash" example:"admin123"`
}