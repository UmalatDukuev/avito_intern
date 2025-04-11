package repository

import (
	"avito_intern/internal/handler/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (email, role, password) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Email, user.Role, user.Password)
	if err := row.Scan(&id); err != nil {
		return "0", err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}
