package repository

import (
	"avito_intern/internal/handler/entity"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	PVZ
	Reception
	Product
}

type Authorization interface {
	CreateUser(entity.User) (string, error)
	GetUser(email, password string) (entity.User, error)
}

type PVZ interface {
	CreatePVZ(pvz entity.PVZ) (string, error)
}

type Reception interface {
	CreateReception(pvzID string) (string, error)
	GetActiveReception(pvzID string) (*entity.Reception, error)
}

type Product interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		PVZ:           NewPVZPostgres(db),
		Reception:     NewReceptionPostgres(db),
		Product:       NewProductPostgres(db),
	}
}
