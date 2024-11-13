package service

import (
	"net/http"

	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type MenuServiceInterface interface {
	AddMenu(models.MenuItem) error
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
	return nil
}

func (s *menuService) GetAll() ([]models.MenuItem, error) {
	return []models.MenuItem{}, nil
}

func (s *menuService) GetOne(id string) (models.MenuItem, error) {
	return models.MenuItem{}, nil
}

func (s *menuService) Update(r *http.Request, id string) error {
	return nil
}

func (s *menuService) Delete(id string) error {
	return nil
}
