package dal

import (
	"encoding/json"
	"os"

	"hot-coffee/models"
)

type InventoryRepoInterface interface {
	ReadInventory() ([]models.Order, error)
	WriteInventory(listOfInventory []models.Order) error
}

type inventoryRepo struct {
	path string
}

func NewInventoryRepo(path string) *inventoryRepo {
	return &inventoryRepo{path: path}
}

func (r *inventoryRepo) ReadInventory() ([]models.Order, error) {
	var listOfInventory []models.Order
	jsonContent, err := os.ReadFile(r.path)
	if err != nil {
		return listOfInventory, err
	}

	if len(jsonContent) > 0 {
		err = json.Unmarshal(jsonContent, &listOfInventory)
		if err != nil {
			return listOfInventory, err
		}
	}

	return listOfInventory, nil
}

func (r *inventoryRepo) WriteInventory(listOfInventory []models.Order) error {
	jsonData, err := json.MarshalIndent(listOfInventory, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
