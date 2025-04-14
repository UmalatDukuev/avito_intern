package service

import (
	"avito_intern/internal/handler/entity"
	"avito_intern/internal/repository"
	"fmt"
	"time"
)

type PVZService struct {
	repo repository.PVZ
}

func NewPVZService(repo repository.PVZ) *PVZService {
	return &PVZService{repo: repo}
}

func (s *PVZService) CreatePVZ(pvz entity.PVZ) (string, error) {
	if pvz.City != "Москва" && pvz.City != "Санкт-Петербург" && pvz.City != "Казань" {
		return "", fmt.Errorf("invalid city, must be Москва, Санкт-Петербург, or Казань")
	}

	pvzS := entity.PVZ{
		City:             pvz.City,
		RegistrationDate: time.Now(),
	}

	return s.repo.CreatePVZ(pvzS)
}
