package service

import (
	"avito_intern/internal/handler/response"
	"avito_intern/internal/repository"
	"fmt"
	"time"
)

type ReceptionService struct {
	repo repository.Reception
}

func NewReceptionService(repo repository.Reception) *ReceptionService {
	return &ReceptionService{repo: repo}
}

func (s *ReceptionService) CreateReception(pvzID string) (string, error) {
	existingReception, err := s.repo.GetActiveReception(pvzID)
	if err != nil {
		return "", fmt.Errorf("error checking for existing reception: %v", err)
	}

	if existingReception != nil {
		return "", fmt.Errorf("there is already an active reception for this PVZ")
	}

	id, err := s.repo.CreateReception(pvzID)
	if err != nil {
		return "", fmt.Errorf("failed to create reception: %v", err)
	}

	return id, nil
}

func (s *ReceptionService) CloseLastReception(pvzID string) (*response.ReceptionResponse, error) {
	activeReception, err := s.repo.GetActiveReception(pvzID)
	if err != nil {
		return nil, fmt.Errorf("failed to get active reception: %v", err)
	}
	if activeReception == nil {
		return nil, fmt.Errorf("no active reception found for this PVZ")
	}

	activeReception.Status = "close"
	activeReception.ClosedAt = time.Now()

	err = s.repo.UpdateReceptionStatus(activeReception)
	if err != nil {
		return nil, fmt.Errorf("failed to close reception: %v", err)
	}

	return activeReception, nil
}
