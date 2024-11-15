package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"hot-coffee/internal"
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
	internal.Logger.Info("AddMenu called", "method", "AddMenu")

	menuItems, err := s.menuRepo.ReadMenu()
	if err != nil {
		internal.Logger.Error("Failed to read menu items from storage", "error", err)
		return err
	}

	var menuItem models.MenuItem
	err = json.NewDecoder(r.Body).Decode(&menuItem)
	if err != nil {
		if err == io.EOF {
			internal.Logger.Warn("Empty request body", "error", err)
			return io.EOF
		}
		internal.Logger.Error("Failed to decode request body", "error", err)
		return err
	}

	// checking if a menu item with the same ID already exists
	for _, existingItem := range menuItems {
		if existingItem.ID == menuItem.ID {
			internal.Logger.Warn("Menu item with the same ID already exists", "menuItemID", menuItem.ID)
			return fmt.Errorf("menu item with ID '%s' already exists", menuItem.ID)
		}
	}

	// check if the ingredients exist in inventory.json storage
	inventoryItems, err := s.menuRepo.ReadInventory()
	if err != nil {
		internal.Logger.Error("Failed to read inventory items from storage", "error", err)
		return err
	}

	// creating a map to lookup inventory items by their IngredientID
	inventoryMap := make(map[string]models.InventoryItem)
	for _, item := range inventoryItems {
		inventoryMap[item.IngredientID] = item
	}

	// checking each ingredient in the menu item against the inventory
	for _, ingredient := range menuItem.Ingredients {
		if _, exists := inventoryMap[ingredient.IngredientID]; !exists {
			internal.Logger.Warn("Ingredient not found in inventory.json", "ingredientID", ingredient.IngredientID)
			return fmt.Errorf("ingredient '%s' is not present in inventory.json", ingredient.IngredientID)
		}
	}

	// checking the fields for emptiness or correctness
	if menuItem.ID == "" {
		internal.Logger.Error("Missing required field", "field", "product_id")
		return fmt.Errorf("missing: product_id")
	}
	if menuItem.Name == "" {
		internal.Logger.Error("Missing required field", "field", "name")
		return fmt.Errorf("missing: name")
	}
	if menuItem.Description == "" {
		internal.Logger.Error("Missing required field", "field", "description")
		return fmt.Errorf("missing: description")
	}
	if menuItem.Price <= 0 {
		internal.Logger.Error("Invalid field value", "field", "price", menuItem.Price)
		return fmt.Errorf("invalid price: must be greater than zero")
	}
	if len(menuItem.Ingredients) == 0 {
		internal.Logger.Error("Missing required field", "field", "ingredients")
		return fmt.Errorf("missing: ingredients ")
	}

	menuItems = append(menuItems, menuItem)
	err = s.menuRepo.WriteMenu(menuItems)
	if err != nil {
		internal.Logger.Error("Failed to write menu items to storage", "error", err)
		return err
	}

	internal.Logger.Info("Menu item added successfully", "menuItemID", menuItem.ID)
	return nil
}

func (s *menuService) GetAll() ([]models.MenuItem, error) {
	internal.Logger.Info("GetAll called", "method", "GetAll")

	menuItems, err := s.menuRepo.ReadMenu()
	if err != nil {
		internal.Logger.Error("Failed to read menu items from storage", "error", err)
		return menuItems, err
	}

	internal.Logger.Info("Successfully retrieved all menu items", "count", len(menuItems))
	return menuItems, nil
}

func (s *menuService) GetOne(id string) (models.MenuItem, error) {
	internal.Logger.Info("GetOne called", "method", "GetOne", "menuItemID", id)

	var menuItem models.MenuItem
	menuItems, err := s.menuRepo.ReadMenu()
	if err != nil {
		internal.Logger.Error("Failed to read menu items from storage", "error", err)
		return menuItem, err
	}

	for _, item := range menuItems {
		if item.ID == id {
			menuItem = item
			internal.Logger.Info("Menu item found", "menuItemID", id)
			return menuItem, nil
		}
	}

	internal.Logger.Warn("Menu item not found", "menuItemID", id)
	return menuItem, fmt.Errorf("menu item with ID %s not found", id)
}

func (s *menuService) Update(r *http.Request, id string) error {
	internal.Logger.Info("Update called", "method", "Update", "menuItemID", id)

	var menuItem map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&menuItem)
	if err != nil {
		internal.Logger.Error("Failed to decode JSON request body", "error", err)
		return fmt.Errorf("error decoding JSON: %v", err)
	}

	// reading the menu_items.json storage
	menuItems, err := s.menuRepo.ReadMenu()
	if err != nil {
		internal.Logger.Error("Failed to read menu items from storage", "error", err)
		return err
	}

	var found bool
	// iterating through the fields in the map to see if they are empty or not
	for i, item := range menuItems {
		if item.ID == id {
			found = true
			internal.Logger.Info("Menu item found", "menuItemID", id)

			// only updating the fields that are present in the request
			// returning an error if request contains ID field, because it cannot be changed
			if _, prs := menuItem["product_id"].(string); prs {
				internal.Logger.Warn("Attempted to modify product_id", "menuItemID", id)
				return fmt.Errorf("product_id cannot be changed")
			}

			if name, prs := menuItem["name"].(string); prs {
				item.Name = name
			} else if _, exists := menuItem["name"]; exists {
				internal.Logger.Error("Invalid type for 'name'", "menuItemID", id)
				return fmt.Errorf("invalid type for 'name': expected a string")
			}

			if description, prs := menuItem["description"].(string); prs {
				item.Description = description
			} else if _, exists := menuItem["description"]; exists {
				internal.Logger.Error("Invalid type for 'description'", "menuItemID", id)
				return fmt.Errorf("invalid type for 'description': expected a string")
			}

			if price, prs := menuItem["price"].(float64); prs {
				item.Price = price
			} else if _, exists := menuItem["price"]; exists {
				internal.Logger.Error("Invalid type for 'price'", "menuItemID", id)
				return fmt.Errorf("invalid type for 'description': expected float64")
			}

			if ingredients, ok := menuItem["ingredients"].([]interface{}); ok {
				inventoryItems, err := s.menuRepo.ReadInventory()
				if err != nil {
					internal.Logger.Error("Failed to read inventory items from storage", "error", err)
					return fmt.Errorf("failed to read inventory")
				}

				// creating a map to lookup inventory items by their IngredientID
				inventoryMap := make(map[string]models.InventoryItem)
				for _, invItem := range inventoryItems {
					inventoryMap[invItem.IngredientID] = invItem
				}

				for _, ing := range ingredients {
					// each ing should be a map[string]interface{}
					ingredientMap, ok := ing.(map[string]interface{})
					if !ok {
						internal.Logger.Error("Invalid ingredient format", "menuItemID", id)
						return fmt.Errorf("invalid ingredient format")
					}

					var ingredientID string
					var quantity float64

					// getting ingredient_id and quantity from the map
					if id, ok := ingredientMap["ingredient_id"].(string); ok {
						ingredientID = id
					} else {
						internal.Logger.Error("Missing or invalid ingredient_id in ingredients", "menuItemID", id)
						return fmt.Errorf("missing or invalid ingredient_id in ingredients")
					}
					if qty, ok := ingredientMap["quantity"].(float64); ok {
						quantity = qty
					} else {
						internal.Logger.Error("Missing or invalid quantity in ingredients", "menuItemID", id)
						return fmt.Errorf("missing or invalid quantity in ingredients")
					}

					// Check if the ingredient exists in the inventory
					if _, exists := inventoryMap[ingredientID]; !exists {
						internal.Logger.Warn("Ingredient not found in inventory.json", "menuItemID", id, "ingredientID", ingredientID)
						return fmt.Errorf("ingredient '%s' not found in inventory.json", ingredientID)
					}

					// Check if this ingredient already exists in item.Ingredients
					found := false
					for i, existingIngredient := range item.Ingredients {
						if existingIngredient.IngredientID == ingredientID {
							// Update the quantity of the existing ingredient
							item.Ingredients[i].Quantity = quantity
							found = true
							internal.Logger.Info("Updated ingredient quantity", "menuItemID", id, "ingredientID", ingredientID, "newQuantity", quantity)
							break
						}
					}

					// If ingredient not found, append it as a new ingredient
					if !found {
						item.Ingredients = append(item.Ingredients, models.MenuItemIngredient{
							IngredientID: ingredientID,
							Quantity:     quantity,
						})
						internal.Logger.Info("Added new ingredient", "menuItemID", id, "ingredientID", ingredientID, "quantity", quantity)
					}
				}
			}

			menuItems[i] = item
			break
		}
	}

	if !found {
		internal.Logger.Warn("Menu item not found", "menuItemID", id)
		return fmt.Errorf("menu item with ID %s not found", id)
	}

	err = s.menuRepo.WriteMenu(menuItems)
	if err != nil {
		internal.Logger.Error("Failed to write updated menu items to storage", "error", err)
		return fmt.Errorf("error writing updated menu items to storage: %v", err)
	}

	internal.Logger.Info("Menu item updated successfully", "menuItemID", id)
	return nil
}

func (s *menuService) Delete(id string) error {
	internal.Logger.Info("Delete called", "method", "Delete", "menuItemID", id)

	menuItems, err := s.menuRepo.ReadMenu()
	if err != nil {
		internal.Logger.Error("Failed to read menu items from storage", "error", err)
		return err
	}

	var updatedMenuItems []models.MenuItem
	var found bool

	for _, item := range menuItems {
		if item.ID != id {
			updatedMenuItems = append(updatedMenuItems, item)
		} else {
			internal.Logger.Info("Menu item found for deletion", "menuItemID", id)
			found = true
		}
	}

	if !found {
		internal.Logger.Warn("Menu item not found", "menuItemID", id)
		return fmt.Errorf("menu Item with ID '%v' is not found", id)
	}

	err = s.menuRepo.WriteMenu(updatedMenuItems)
	if err != nil {
		internal.Logger.Error("Failed to write updated menu items to storage", "error", err)
		return fmt.Errorf("error writing the json storage")
	}

	internal.Logger.Info("Menu item deleted successfully", "menuItemID", id)
	return nil
}
