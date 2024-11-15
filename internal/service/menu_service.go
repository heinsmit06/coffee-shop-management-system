package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type MenuServiceInterface interface {
	AddMenu(*http.Request) error
	GetAll() ([]models.MenuItem, error)
	GetOne(id string) (models.MenuItem, error)
	Update(r *http.Request, id string) error
	Delete(id string) error
}

type menuService struct {
	menuRepo dal.MenuRepositoryInterface
}

func NewMenuService(menuRepo dal.MenuRepositoryInterface) *menuService {
	return &menuService{menuRepo: menuRepo}
}

func (s *menuService) AddMenu(r *http.Request) error {
	var menuItem models.MenuItem
	menuItems, err := s.menuRepo.ReadMenu()
	if err != nil {
		fmt.Println("qw")
		return err
	}

	err = json.NewDecoder(r.Body).Decode(&menuItem)
	if err != nil {
		fmt.Println("fgsdg")
		return err
	}

	// checking the fields for emptiness or correctness
	if menuItem.ID == "" {
		return fmt.Errorf("missing: product_id")
	}
	if menuItem.Name == "" {
		return fmt.Errorf("missing: name")
	}
	if menuItem.Description == "" {
		return fmt.Errorf("missing: description")
	}
	if menuItem.Price <= 0 {
		return fmt.Errorf("invalid price: must be greater than zero")
	}
	if len(menuItem.Ingredients) == 0 {
		return fmt.Errorf("missing: ingredients ")
	}

	menuItems = append(menuItems, menuItem)
	err = s.menuRepo.WriteMenu(menuItems)
	if err != nil {
		return nil
	}

	return nil
}

func (s *menuService) GetAll() ([]models.MenuItem, error) {
	menuItems, err := s.menuRepo.ReadMenu()
	if err != nil {
		return menuItems, err
	}
	return menuItems, nil
}

func (s *menuService) GetOne(id string) (models.MenuItem, error) {
	var menuItem models.MenuItem
	menuItems, err := s.menuRepo.ReadMenu()
	if err != nil {
		return menuItem, err
	}

	for _, item := range menuItems {
		if item.ID == id {
			menuItem = item
		}
	}

	return menuItem, nil
}

func (s *menuService) Update(r *http.Request, id string) error {
	return nil
}

func (s *menuService) Delete(id string) error {
	return nil
}
