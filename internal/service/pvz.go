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

func (s *PVZService) GetPVZList(startDate, endDate *time.Time, page, limit int) ([]entity.PVZResponse, error) {
	pvzList, err := s.repo.GetPVZWithDetails(startDate, endDate, page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get pvz list: %v", err)
	}

	return pvzList, nil
}

func (s *PVZService) GetByID(pvzID string) (*entity.PVZ, error) {
	return s.repo.GetByID(pvzID)
}
