package service

import (
	"avito_intern/internal/handler/entity"
	"avito_intern/internal/repository"
)

type Authorization interface {
	CreateUser(RegisterInput) (string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (string, string, error)
	GenerateDummyToken(userType string) (string, error)
}

type PVZ interface {
	CreatePVZ(pvz entity.PVZ) (string, error)
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
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		PVZ:           NewPVZService(repo.PVZRepo),
	}
}
