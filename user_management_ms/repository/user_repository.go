package repository

import (
	"database/sql"
	"user_management_ms/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *model.User) error {
	query := `INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.CreatedAt).Scan(&user.ID)
	return err
}
