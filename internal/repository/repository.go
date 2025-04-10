package repository

import (
	"avito_intern/models"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
}

type Authorization interface {
	CreateUser(models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
