package service

import (
	"encoding/json"
	"fmt"
	"io"
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
		if err == io.EOF {
			return io.EOF
		}
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
	var menuItem map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&menuItem)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}

	// reading the menu_items.json storage
	menuItems, err := s.menuRepo.ReadMenu()
	if err != nil {
		return err
	}

	var found bool
	// iterating through the fields in the map to see if they are empty or not
	for i, item := range menuItems {
		if item.ID == id {
			found = true

			// only updating the fields that are present in the request
			// returning an error if request contains ID field, because it cannot be changed
			if _, prs := menuItem["product_id"].(string); prs {
				return fmt.Errorf("product_id cannot be changed")
			}

			if name, prs := menuItem["name"].(string); prs {
				item.Name = name
			} else if _, exists := menuItem["name"]; exists {
				return fmt.Errorf("invalid type for 'name': expected a string")
			}

			if description, prs := menuItem["description"].(string); prs {
				item.Description = description
			} else if _, exists := menuItem["description"]; exists {
				return fmt.Errorf("invalid type for 'description': expected a string")
			}

			if price, prs := menuItem["price"].(float64); prs {
				item.Price = price
			} else if _, exists := menuItem["price"]; exists {
				return fmt.Errorf("invalid type for 'description': expected float64")
			}

			if ingredients, ok := menuItem["ingredients"].([]interface{}); ok {
				for _, ing := range ingredients {
					// Each ing should be a map[string]interface{}
					ingredientMap, ok := ing.(map[string]interface{})
					if !ok {
						return fmt.Errorf("invalid ingredient format")
					}

					var ingredientID string
					var quantity float64

					// Get ingredient_id and quantity from the map
					if id, ok := ingredientMap["ingredient_id"].(string); ok {
						ingredientID = id
					} else {
						return fmt.Errorf("missing or invalid ingredient_id in ingredients")
					}
					if qty, ok := ingredientMap["quantity"].(float64); ok {
						quantity = qty
					} else {
						return fmt.Errorf("missing or invalid quantity in ingredients")
					}

					// Check if this ingredient already exists in item.Ingredients
					found := false
					for i, existingIngredient := range item.Ingredients {
						if existingIngredient.IngredientID == ingredientID {
							// Update the quantity of the existing ingredient
							item.Ingredients[i].Quantity = quantity
							found = true
							break
						}
					}

					// If ingredient not found, append it as a new ingredient
					if !found {
						item.Ingredients = append(item.Ingredients, models.MenuItemIngredient{
							IngredientID: ingredientID,
							Quantity:     quantity,
						})
					}
				}
			}

			menuItems[i] = item
			break
		}
	}

	if !found {
		return fmt.Errorf("menu item with ID %s not found", id)
	}

	err = s.menuRepo.WriteMenu(menuItems)
	if err != nil {
		return fmt.Errorf("error writing updated menu items to storage: %v", err)
	}

	return nil
}

func (s *menuService) Delete(id string) error {
	menuItems, err := s.menuRepo.ReadMenu()
	if err != nil {
		return err
	}

	var updatedMenuItems []models.MenuItem
	var found bool

	for _, item := range menuItems {
		if item.ID != id {
			updatedMenuItems = append(updatedMenuItems, item)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("menu Item with ID '%v' is not found", id)
	}

	err = s.menuRepo.WriteMenu(updatedMenuItems)
	if err != nil {
		return fmt.Errorf("error writing the json storage")
	}

	return nil
}
