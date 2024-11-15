package internal

import "errors"

var (
	ErrNoIngredientID         = errors.New("No Ingredient ID")
	ErrNoIngredientName       = errors.New("No Ingredient Name")
	ErrNoIngredientUnit       = errors.New("No Ingredient Unit")
	ErrIngredientAlreadyExist = errors.New("Ingredient with the same ID/Name already exists")
	ErrIngredientNotExist     = errors.New("There is no Ingredient with such ID")
	// ErrIngredientNoQuantity   = errors.New("Ingredient quantity must be specified to update")
	ErrInventoryIsEmpty = errors.New("Inventory is empty")
	ErrOrdersIsEmpty    = errors.New("No orders yet")
	ErrOrderNotExist    = errors.New("There is no Order with such ID")
	ErrOrderClosed      = errors.New("Order already closed")
	ErrMenuItemNotExist = errors.New("There is no MenuItem with such Product ID")
	ErrMenuIsEmpty      = errors.New("Menu is empty")
)
