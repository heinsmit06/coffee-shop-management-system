package dal

import (
	"encoding/json"
	"fmt"
	"os"

	"hot-coffee/models"
)

type MenuRepositoryInterface interface {
	ReadMenu() ([]models.MenuItem, error)
	WriteMenu(listOfMenu []models.MenuItem) error
}

type menuRepo struct {
	path string
}

func NewMenuRepo(path string) *menuRepo {
	return &menuRepo{path: path}
}

func (r *menuRepo) ReadMenu() ([]models.MenuItem, error) {
	var listOfMenu []models.MenuItem
	jsonContent, err := os.ReadFile(r.path + "menu_items.json")
	if os.IsNotExist(err) {
		file, err := os.OpenFile(r.path+"menu_items.json", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return listOfMenu, err
		}
		file.Close()

		return listOfMenu, nil
	} else if err != nil {
		return listOfMenu, err
	}

	if len(jsonContent) > 0 {
		err = json.Unmarshal(jsonContent, &listOfMenu)
		if err != nil {
			fmt.Println("this")
			return listOfMenu, err
		}
	}

	return listOfMenu, nil
}

func (r *menuRepo) WriteMenu(listOfMenu []models.MenuItem) error {
	jsonData, err := json.MarshalIndent(listOfMenu, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.path+"menu_items.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
