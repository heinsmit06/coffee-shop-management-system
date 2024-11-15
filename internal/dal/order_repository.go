package dal

import (
	"encoding/json"
	"os"

	"hot-coffee/internal"
	"hot-coffee/models"
)

type OrderRepoInterface interface {
	GetAll() ([]byte, error)
	ReadOrders() ([]models.Order, error)
	WriteOrders(listOfOrders []models.Order) error
	ReadInventory() ([]models.InventoryItem, error)
	ReadMenu() ([]models.MenuItem, error)
	WriteInventory(listOfInventory []models.InventoryItem) error
}

type orderRepo struct {
	path string
}

// CUNSTRUCTURE
func NewOrderRepo(path string) *orderRepo {
	return &orderRepo{path: path}
}

// Methods:
func (r *orderRepo) GetAll() ([]byte, error) {
	jsonContent, err := os.ReadFile(r.path + "/orders.json")
	if err != nil {
		if os.IsNotExist(err) {
			// os.WriteFile(r.path+"/inventory.json", []byte{}, 0644)
			return []byte{}, internal.ErrOrdersIsEmpty
		} else {
			return []byte{}, err
		}
	}

	return jsonContent, nil
}

func (r *orderRepo) ReadOrders() ([]models.Order, error) {
	listOfOrders := []models.Order{}
	jsonContent, err := os.ReadFile(r.path + "/orders.json")
	if err != nil {
		if os.IsNotExist(err) {
			os.WriteFile(r.path+"/orders.json", []byte{}, 0o644)
		} else {
			return listOfOrders, err
		}
	}

	if len(jsonContent) > 0 {
		err = json.Unmarshal(jsonContent, &listOfOrders)
		if err != nil {
			return listOfOrders, err
		}
	}

	return listOfOrders, nil
}

func (r *orderRepo) WriteOrders(listOfOrders []models.Order) error {
	jsonData, err := json.MarshalIndent(listOfOrders, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.path+"/orders.json", jsonData, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepo) ReadInventory() ([]models.InventoryItem, error) {
	listOfInventory := []models.InventoryItem{}
	jsonContent, err := os.ReadFile(r.path + "/inventory.json")
	if err != nil {
		if os.IsNotExist(err) {
			os.WriteFile(r.path+"/inventory.json", []byte{}, 0o644)
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

	return listOfInventory, nil
}

func (r *orderRepo) ReadMenu() ([]models.MenuItem, error) {
	listOfMenu := []models.MenuItem{}
	jsonContent, err := os.ReadFile(r.path + "menu_items.json")
	if err != nil {
		if os.IsNotExist(err) {
			os.WriteFile(r.path+"/menu_items.json", []byte{}, 0o644)
		} else {
			return listOfMenu, err
		}
	}

	if len(jsonContent) > 0 {
		err = json.Unmarshal(jsonContent, &listOfMenu)
		if err != nil {
			return listOfMenu, err
		}
	}

	return listOfMenu, nil
}

func (r *orderRepo) WriteInventory(listOfInventory []models.InventoryItem) error {
	jsonData, err := json.MarshalIndent(listOfInventory, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.path+"/inventory.json", jsonData, 0o644)
	if err != nil {
		return err
	}

	return nil
}
