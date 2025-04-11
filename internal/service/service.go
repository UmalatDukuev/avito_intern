package service

import (
	"avito_intern/internal/repository"
)

type Authorization interface {
	CreateUser(RegisterInput) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	GenerateDummyToken(userType string) (string, error)
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
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
