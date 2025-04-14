package service

import (
	"avito_intern/internal/handler/entity"
	"avito_intern/internal/handler/response"
	"avito_intern/internal/repository"
	"time"
)

type Authorization interface {
	CreateUser(RegisterInput) (string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (string, string, error)
	GenerateDummyToken(role string) (string, error)
}

type PVZ interface {
	CreatePVZ(pvz entity.PVZ) (string, error)
	GetPVZList(startDate, endDate *time.Time, page, limit int) ([]entity.PVZResponse, error)
	GetByID(pvzID string) (*entity.PVZ, error)
}

type Reception interface {
	CreateReception(pvzID string) (string, error)
	CloseLastReception(pvzID string) (*response.ReceptionResponse, error)
}

type Product interface {
	AddProductToReception(pvzID string, productType string) (string, error)
	DeleteLastProduct(pvzID string) (string, error)
}

type RegisterInput struct {
	Email    string
	Password string
	Role     string
}

type LoginInput struct {
	Email    string
	Password string
}

type Service struct {
	Authorization
	PVZ
	Reception
	Product
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		PVZ:           NewPVZService(repo.PVZ),
		Reception:     NewReceptionService(repo.Reception),
		Product:       NewProductService(repo.Product),
	}
}
