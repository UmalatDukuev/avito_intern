package service

import (
	"avito_intern/internal/handler/entity"
	"avito_intern/internal/repository"
	"fmt"
	"time"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) AddProductToReception(pvzID string, productType string) (string, error) {
	activeReception, err := s.repo.GetActiveReception(pvzID)
	if err != nil {
		return "", fmt.Errorf("error checking active reception: %v", err)
	}

	if activeReception == nil {
		return "", fmt.Errorf("no active reception for this PVZ")
	}

	product := entity.Product{
		Type:     productType,
		DateTime: time.Now(),
		ID:       activeReception.ID,
	}

	productID, err := s.repo.AddProductToReception(product)
	if err != nil {
		return "", fmt.Errorf("failed to add product: %v", err)
	}

	return productID, nil
}

func (s *ProductService) DeleteLastProduct(pvzID string) (string, error) {
	activeReception, err := s.repo.GetActiveReception(pvzID)
	if err != nil {
		return "", fmt.Errorf("failed to get active reception: %v", err)
	}
	if activeReception == nil {
		return "", fmt.Errorf("no active reception for this PVZ")
	}

	lastProduct, err := s.repo.GetLastProduct(activeReception.ID)
	if err != nil {
		return "", fmt.Errorf("failed to get last product: %v", err)
	}
	if lastProduct == nil {
		return "", fmt.Errorf("no products to delete")
	}

	err = s.repo.DeleteProduct(lastProduct.ID)
	if err != nil {
		return "", fmt.Errorf("failed to delete product: %v", err)
	}

	return lastProduct.ID, nil
}
