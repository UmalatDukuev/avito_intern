package service

import (
	"avito_intern/internal/handler/entity"
	"avito_intern/internal/repository"
)

type Authorization interface {
	CreateUser(RegisterInput) (string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (string, string, error)
	GenerateDummyToken(role string) (string, error)
}

type PVZ interface {
	CreatePVZ(pvz entity.PVZ) (string, error)
}

type Reception interface {
	CreateReception(pvzID string) (string, error)
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
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		PVZ:           NewPVZService(repo.PVZ),
		Reception:     NewReceptionService(repo.Reception),
	}
}
