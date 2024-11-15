package dal

import (
	"encoding/json"
	"fmt"
	"os"

	"hot-coffee/models"
)

type InventoryRepoInterface interface {
	ReadInventory() ([]models.InventoryItem, error)
	WriteInventory(listOfInventory []models.InventoryItem) error
}

type inventoryRepo struct {
	path string
}

func NewInventoryRepo(path string) *inventoryRepo {
	return &inventoryRepo{path: path}
}

func (r *inventoryRepo) ReadInventory() ([]models.InventoryItem, error) {
	var listOfInventory []models.InventoryItem
	jsonContent, err := os.ReadFile(r.path + "/inventory_item.json")
	if err != nil {
		if os.IsNotExist(err) {
			os.WriteFile(r.path+"/inventory_item.json", []byte{}, 0o644)
		} else {
			return listOfInventory, err
		}
	}

	if len(jsonContent) > 0 {
		err = json.Unmarshal(jsonContent, &listOfInventory)
		if err != nil {
			return listOfInventory, err
		}
	}

	fmt.Println("ReadInventory complete")
	return listOfInventory, nil
}

func (r *inventoryRepo) WriteInventory(listOfInventory []models.InventoryItem) error {
	jsonData, err := json.MarshalIndent(listOfInventory, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.path+"/inventory_item.json", jsonData, 0o644)
	if err != nil {
		return err
	}

	fmt.Println("WriteInventory complete")
	return nil
}
