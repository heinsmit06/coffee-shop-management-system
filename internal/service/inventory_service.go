package service

import (
	"encoding/json"
	"net/http"
	"strings"

	"hot-coffee/internal"
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type InventoryServiceInterface interface {
	AddInventory(r *http.Request) error
	GetAll() ([]byte, error)
	GetOne(r *http.Request) ([]byte, error)
	Update(r *http.Request) error
	Delete(r *http.Request) error
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

func (s *inventoryService) GetAll() ([]byte, error) {
	return s.inventoryRepo.GetAll()
}

func (s *inventoryService) GetOne(r *http.Request) ([]byte, error) {
	splitedURL := strings.Split(r.URL.Path, "/")
	id := splitedURL[len(splitedURL)-1]

	listOfInventoryItems, err := s.inventoryRepo.ReadInventory()
	if err != nil {
		return nil, err
	}

	for i, inventoryItem := range listOfInventoryItems {
		if inventoryItem.IngredientID == id {
			return json.MarshalIndent(listOfInventoryItems[i], "", " ")
		}
	}

	return nil, internal.ErrIngredientNotExist
}

func (s *inventoryService) Update(r *http.Request) error {
	splitedURL := strings.Split(r.URL.Path, "/")
	id := splitedURL[len(splitedURL)-1]

	var UpdatedInventoryItem models.InventoryItem

	err := json.NewDecoder(r.Body).Decode(&UpdatedInventoryItem)
	if err != nil {
		return err
	}

	listOfInventoryItems, err := s.inventoryRepo.ReadInventory()
	if err != nil {
		return err
	}

	for i, inventoryItem := range listOfInventoryItems {
		if inventoryItem.IngredientID == id {
			listOfInventoryItems[i].Name = UpdatedInventoryItem.Name
			listOfInventoryItems[i].Unit = UpdatedInventoryItem.Unit
			listOfInventoryItems[i].Quantity = UpdatedInventoryItem.Quantity
			s.inventoryRepo.WriteInventory(listOfInventoryItems)
			return nil
		}
	}

	return internal.ErrIngredientNotExist
}

func (s *inventoryService) Delete(r *http.Request) error {
	splitedURL := strings.Split(r.URL.Path, "/")
	id := splitedURL[len(splitedURL)-1]

	listOfInventoryItems, err := s.inventoryRepo.ReadInventory()
	if err != nil {
		return err
	}

	if len(listOfInventoryItems) < 1 {
		return internal.ErrInventoryIsEmpty
	}

	for i, inventoryItem := range listOfInventoryItems {
		if inventoryItem.IngredientID == id {
			listOfInventoryItems = append(listOfInventoryItems[:i], listOfInventoryItems[i+1:]...)
			s.inventoryRepo.WriteInventory(listOfInventoryItems)
			return nil
		}
	}

	return internal.ErrIngredientNotExist
}
