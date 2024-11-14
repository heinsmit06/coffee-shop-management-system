package dal

import "hot-coffee/models"

type MenuRepositoryInterface interface {
	GetAll() ([]models.MenuItem, error)
}

type menuRepo struct {
	dirPath string
}

func NewMenuRepo(path string) *menuRepo {
	return &menuRepo{dirPath: path}
}

func (r *menuRepo) GetAll() ([]models.MenuItem, error) {
	return []models.MenuItem{}, nil
}
