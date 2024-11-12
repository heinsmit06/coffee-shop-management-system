package service

import (
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type MenuServiceInterface interface {
	Add(models.MenuItem) error
	///
	///
	//
}

type menuService struct {
	menuRepo dal.MenuRepositoryInterface
}

func NewMenuService(menuRepo dal.MenuRepositoryInterface) *menuService {
	return &menuService{menuRepo: menuRepo}
}

func (s *menuService) Add(newMenu models.MenuItem) error {
	// Valiadtionnasd
	// s.menuRepo.Add(newMenu)
	return nil
}

// SOLID && Dependency injection
/*
	S - Single responsibility
	O - Open/Closed principle
	L - Liskov's subtitution
	I - Interface segregation
	D - Dependency inversion
*/
