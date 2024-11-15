package dal

import (
	"encoding/json"
	"os"

	"hot-coffee/internal"
	"hot-coffee/models"
)

type MenuRepositoryInterface interface {
	ReadMenu() ([]models.MenuItem, error)
	WriteMenu(listOfMenu []models.MenuItem) error
	ReadInventory() ([]models.InventoryItem, error)
}

type menuRepo struct {
	path string
}

func NewMenuRepo(path string) *menuRepo {
	return &menuRepo{path: path}
}

func (r *menuRepo) ReadMenu() ([]models.MenuItem, error) {
	internal.Logger.Info("ReadMenu called", "method", "ReadMenu")
	var listOfMenu []models.MenuItem

	jsonContent, err := os.ReadFile(r.path + "menu_items.json")
	if os.IsNotExist(err) {
		internal.Logger.Warn("menu_items.json not found, creating new file", "path", r.path+"menu_items.json")

		file, err := os.OpenFile(r.path+"menu_items.json", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			internal.Logger.Error("Failed to create menu_items.json", "error", err)
			return listOfMenu, err
		}
		file.Close()

		internal.Logger.Info("Created new menu_items.json", "path", r.path+"menu_items.json")
		return listOfMenu, nil
	} else if err != nil {
		internal.Logger.Error("Failed to read menu_items.json", "error", err)
		return listOfMenu, err
	}

	if len(jsonContent) > 0 {
		err = json.Unmarshal(jsonContent, &listOfMenu)
		if err != nil {
			internal.Logger.Error("Failed to unmarshal JSON content", "error", err)
			return listOfMenu, err
		}
	} else {
		internal.Logger.Info("menu_items.json is empty")
	}

	return listOfMenu, nil
}

func (r *menuRepo) WriteMenu(listOfMenu []models.MenuItem) error {
	internal.Logger.Info("WriteMenu called", "method", "WriteMenu", "count", len(listOfMenu))

	jsonData, err := json.MarshalIndent(listOfMenu, "", " ")
	if err != nil {
		internal.Logger.Error("Failed to marshal menu items to JSON", "error", err)
		return err
	}

	err = os.WriteFile(r.path+"menu_items.json", jsonData, 0644)
	if err != nil {
		internal.Logger.Error("Failed to write menu items to menu_items.json", "error", err)
		return err
	}

	internal.Logger.Info("Successfully wrote menu items to menu_items.json", "path", r.path+"menu_items.json")
	return nil
}

func (r *menuRepo) ReadInventory() ([]models.InventoryItem, error) {
	internal.Logger.Info("ReadInventory called", "method", "ReadInventory")
	var listOfInventory []models.InventoryItem

	jsonContent, err := os.ReadFile(r.path + "inventory.json")
	if err != nil {
		internal.Logger.Error("Failed to read inventory.json", "error", err)
		return listOfInventory, err
	}

	if len(jsonContent) > 0 {
		err = json.Unmarshal(jsonContent, &listOfInventory)
		if err != nil {
			internal.Logger.Error("Failed to unmarshal JSON content", "error", err)
			return listOfInventory, err
		}
	} else {
		internal.Logger.Info("inventory.json is empty")
	}

	return listOfInventory, nil
}
