package dal

import (
	"encoding/json"
	"fmt"
	"os"

	"hot-coffee/models"
)

type OrderRepoInterface interface {
	ReadOrders() ([]models.Order, error)
	WriteOrders(listOfOrders []models.Order) error
}

type orderRepo struct {
	path string
}

// CUNSTRUCTURE
func NewOrderRepo(path string) *orderRepo {
	return &orderRepo{path: path}
}

// Methods:
func (r *orderRepo) ReadOrders() ([]models.Order, error) {
	var listOfOrders []models.Order
	jsonContent, err := os.ReadFile(r.path)
	if err != nil {
		return listOfOrders, err
	}

	if len(jsonContent) > 0 {
		err = json.Unmarshal(jsonContent, &listOfOrders)
		if err != nil {
			fmt.Println("readorders Unmarshal")
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

	err = os.WriteFile(r.path, jsonData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("writeOrders")
	return nil
}
