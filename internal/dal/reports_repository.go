package dal

import (
	"encoding/json"
	"os"

	"hot-coffee/models"
)

type ReportsRepoInterface interface {
	ReadMenu() ([]models.MenuItem, error)
	ReadOrders() ([]models.Order, error)
}

type reportsRepo struct {
	path string
}

func NewReportsRepo(path string) *reportsRepo {
	return &reportsRepo{path: path}
}

func (r *reportsRepo) ReadMenu() ([]models.MenuItem, error) {
	var listOfMenu []models.MenuItem
	jsonContent, err := os.ReadFile(r.path + "/menu_items.json")
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

func (r *reportsRepo) ReadOrders() ([]models.Order, error) {
	var listOfOrder []models.Order
	jsonContent, err := os.ReadFile(r.path + "/orders.json")
	if err != nil {
		return listOfOrder, err
	}

	if len(jsonContent) > 0 {
		err = json.Unmarshal(jsonContent, &listOfOrder)
		if err != nil {
			return listOfOrder, err
		}
	}

	return listOfOrder, nil
}
