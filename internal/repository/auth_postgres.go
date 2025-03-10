package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/petprojects9964409/todo_app/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	const op = "Repository.CreateUser"
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}
