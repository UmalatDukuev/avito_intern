package repository

import (
	"avito_intern/internal/handler/entity"
	"avito_intern/internal/handler/response"
	"time"

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
	GetPVZWithDetails(startDate, endDate *time.Time, page, limit int) ([]entity.PVZResponse, error)
	GetByID(pvzID string) (*entity.PVZ, error)
}

type Reception interface {
	CreateReception(pvzID string) (string, error)
	GetActiveReception(pvzID string) (*response.ReceptionResponse, error)
	UpdateReceptionStatus(reception *response.ReceptionResponse) error
}

type Product interface {
	AddProductToReception(product entity.Product) (string, error)
	GetActiveReception(pvzID string) (*response.ReceptionResponse, error)
	GetLastProduct(receptionID string) (*entity.Product, error)
	DeleteProduct(productID string) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		PVZ:           NewPVZPostgres(db),
		Reception:     NewReceptionPostgres(db),
		Product:       NewProductPostgres(db),
	}
}
