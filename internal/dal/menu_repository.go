package dal

import (
	"encoding/json"
	"os"

	"hot-coffee/models"
)

type MenuRepositoryInterface interface {
	ReadMenu() ([]models.Order, error)
	WriteMenu(listOfMenu []models.Order) error
}

type menuRepo struct {
	path string
}

func NewMenuRepo(path string) *menuRepo {
	return &menuRepo{path: path}
}

func (r *menuRepo) ReadMenu() ([]models.Order, error) {
	var listOfMenu []models.Order
	jsonContent, err := os.ReadFile(r.path)
	if err != nil {
		return listOfMenu, err
	}

	if len(jsonContent) > 0 {
		err = json.Unmarshal(jsonContent, &listOfMenu)
		if err != nil {
			return listOfMenu, err
		}
	}

	return listOfMenu, nil
}

func (r *menuRepo) WriteMenu(listOfMenu []models.Order) error {
	jsonData, err := json.MarshalIndent(listOfMenu, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
