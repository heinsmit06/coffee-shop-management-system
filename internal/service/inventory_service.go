package service

import (
	"encoding/json"
	"net/http"

	"hot-coffee/internal"
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
	listOfInventoryItems, err := s.inventoryRepo.ReadInventory()
	if err != nil {
		return err
	}

	var NewInventoryItem models.InventoryItem

	err = json.NewDecoder(r.Body).Decode(&NewInventoryItem)
	if err != nil {
		return err
	}

	if NewInventoryItem.IngredientID == "" {
		return internal.ErrNoIngredientID
	}
	if NewInventoryItem.Name == "" {
		return internal.ErrNoIngredientName
	}
	if NewInventoryItem.Unit == "" {
		return internal.ErrNoIngredientUnit
	}

	for _, InventoryItem := range listOfInventoryItems {
		if InventoryItem.IngredientID == NewInventoryItem.IngredientID ||
			InventoryItem.Name == NewInventoryItem.Name {
			return internal.ErrIngredientAlreadyExist
		}
	}

	listOfInventoryItems = append(listOfInventoryItems, NewInventoryItem)

	s.inventoryRepo.WriteInventory(listOfInventoryItems)
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
