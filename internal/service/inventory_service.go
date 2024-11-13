package service

import (
	"net/http"

	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type InventoryServiceInterface interface {
	AddInventory(r *http.Request) error
	GetAll() ([]models.InventoryItem, error)
	GetOne(id string) (models.InventoryItem, error)
	Update(r *http.Request, id string) error
	Delete(id string) error
}

type inventoryService struct {
	inventoryRepo dal.InventoryRepoInterface
}

func NewInventoryService(inventoryRepo dal.InventoryRepoInterface) *inventoryService {
	return &inventoryService{inventoryRepo: inventoryRepo}
}

func (s *inventoryService) AddInventory(r *http.Request) error {
	return nil
}

func (s *inventoryService) GetAll() ([]models.InventoryItem, error) {
	return []models.InventoryItem{}, nil
}

func (s *inventoryService) GetOne(id string) (models.InventoryItem, error) {
	return models.InventoryItem{}, nil
}

func (s *inventoryService) Update(r *http.Request, id string) error {
	return nil
}

func (s *inventoryService) Delete(id string) error {
	return nil
}
