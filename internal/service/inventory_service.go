package service

import (
	"net/http"

	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type InventoryServiceInterface interface {
	Add(r *http.Request) error
	GetAll() ([]models.InventoryItem, error)
	GetOne() (models.InventoryItem, error)
	Update(r *http.Request) error
	Delete(r *http.Request) error
}

type inventoryService struct {
	inventoryRepo dal.InventoryRepoInterface
}

func NewInventoryService(inventoryRepo dal.InventoryRepoInterface) *inventoryService {
	return &inventoryService{inventoryRepo: inventoryRepo}
}

func (s *inventoryService) Add(r *http.Request) error {
	return nil
}

func (s *inventoryService) GetAll() ([]models.InventoryItem, error) {
	return []models.InventoryItem{}, nil
}

func (s *inventoryService) GetOne() (models.InventoryItem, error) {
	return models.InventoryItem{}, nil
}

func (s *inventoryService) Update(r *http.Request) error {
	return nil
}

func (s *inventoryService) Delete(r *http.Request) error {
	return nil
}
