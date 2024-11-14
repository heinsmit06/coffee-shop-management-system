package internal

import "errors"

var (
	ErrNoIngredientID         = errors.New("No Ingredient ID")
	ErrNoIngredientName       = errors.New("No Ingredient Name")
	ErrNoIngredientUnit       = errors.New("No Ingredient Unit")
	ErrIngredientAlreadyExist = errors.New("Ingredient with the same ID/Name already exists")
)
