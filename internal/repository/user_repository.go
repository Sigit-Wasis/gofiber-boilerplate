package repository

import (
	"context"
	"database/sql"

	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, name, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	var u models.User
	err := r.DB.QueryRowContext(ctx,
		"SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1", id,
	).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &u, err
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.DB.QueryRowContext(ctx,
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at, updated_at",
		user.Name, user.Email,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	_, err := r.DB.ExecContext(ctx,
		"UPDATE users SET name=$1, email=$2, updated_at=NOW() WHERE id=$3",
		user.Name, user.Email, user.ID,
	)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)
	return err
}
