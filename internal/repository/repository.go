package repository

import (
	"avito_intern/internal/handler/entity"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
}

type Authorization interface {
	CreateUser(entity.User) (string, error)
	GetUser(email, password string) (entity.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
