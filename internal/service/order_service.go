package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"hot-coffee/internal"
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type OrderServiceInterface interface {
	Create(*http.Request) error
	GetAll() ([]byte, error)
	GetOne(r *http.Request) ([]byte, error)
	Update(*http.Request) error
	Delete(*http.Request) error
	Close(*http.Request) error
	Test()
}

type orderService struct {
	orderRepo dal.OrderRepoInterface
}

func NewOrderService(orderRepo dal.OrderRepoInterface) *orderService {
	return &orderService{orderRepo: orderRepo}
}

func (s *orderService) Create(r *http.Request) error {
	listOfOrders, err := s.orderRepo.ReadOrders()
	if err != nil {
		return err
	}

	listOfMenuItems, err := s.orderRepo.ReadMenu()
	if err != nil {
		return err
	}

	listOfInventoryItems, err := s.orderRepo.ReadInventory()
	if err != nil {
		return err
	}

	if len(listOfMenuItems) < 1 {
		return internal.ErrMenuIsEmpty
	}

	if len(listOfInventoryItems) < 1 {
		return internal.ErrInventoryIsEmpty
	}

	var newOrder models.Order

	err = json.NewDecoder(r.Body).Decode(&newOrder)

	for _, orderItem := range newOrder.Items {
		for j, menuItem := range listOfMenuItems {
			if orderItem.ProductID == menuItem.ID {
				for _, ingredientItem := range menuItem.Ingredients {
					for i, ingredientItemInv := range listOfInventoryItems {
						if ingredientItem.IngredientID == ingredientItemInv.IngredientID {
							if ingredientItem.Quantity*float64(
								orderItem.Quantity,
							) <= ingredientItemInv.Quantity {
								listOfInventoryItems[i].Quantity -= ingredientItem.Quantity * float64(
									orderItem.Quantity,
								)
							} else {

								ingredientName := ingredientItemInv.Name
								required := ingredientItem.Quantity * float64(orderItem.Quantity)
								available := ingredientItemInv.Quantity
								unit := ingredientItemInv.Unit
								return fmt.Errorf("Insufficient inventory for ingredient '%s'. Required: %.2f%s, Available: %.2f%s.", ingredientName, required, unit, available, unit)
							}
							break
						} else if i == len(listOfInventoryItems)-1 {
							return fmt.Errorf("There is no ingredient %s(IngredientID) for %s(ProductID)", ingredientItem.IngredientID, orderItem.ProductID)
							// return internal.ErrIngredientNotExist
						}
					}
				}
				break
			} else if j == len(listOfMenuItems)-1 {
				return fmt.Errorf("There is no Menu Item with such ProductID: %s", orderItem.ProductID)
				// return internal.ErrMenuItemNotExist
			}
		}
	}

	rgx := regexp.MustCompile(`\d+`)
	var id int
	var highestID int

	for _, order := range listOfOrders {
		highestID, _ = strconv.Atoi(rgx.FindString(order.ID))

		if highestID > id {
			id = highestID
		}
	}

	newOrder.ID = "order" + strconv.Itoa(id+1)
	newOrder.Status = "open"
	newOrder.CreatedAt = time.Now().UTC().Format("2006-01-02T15:04:05Z")

	listOfOrders = append(listOfOrders, newOrder)
	s.orderRepo.WriteInventory(listOfInventoryItems)
	s.orderRepo.WriteOrders(listOfOrders)
	return nil
}

func (s *orderService) GetAll() ([]byte, error) {
	listOfOrders, err := s.orderRepo.ReadOrders()
	if err != nil {
		return []byte{}, err
	}
	if len(listOfOrders) < 1 {
		return []byte{}, internal.ErrOrdersIsEmpty
	}
	jsonContent, err := s.orderRepo.GetAll()
	return jsonContent, err
}

func (s *orderService) GetOne(r *http.Request) ([]byte, error) {
	splitedURL := strings.Split(r.URL.Path, "/")
	id := splitedURL[len(splitedURL)-1]

	listOfOrders, err := s.orderRepo.ReadOrders()
	if err != nil {
		return []byte{}, err
	}

	if len(listOfOrders) < 1 {
		return []byte{}, internal.ErrOrdersIsEmpty
	}
	for i, order := range listOfOrders {
		if order.ID == id {
			return json.MarshalIndent(listOfOrders[i], "", " ")
		}
	}

	return []byte{}, internal.ErrOrderNotExist
}

func (s *orderService) Update(r *http.Request) error {
	splitedURL := strings.Split(r.URL.Path, "/")
	id := splitedURL[len(splitedURL)-1]

	var UpdatedOrder models.Order

	err := json.NewDecoder(r.Body).Decode(&UpdatedOrder)
	if err != nil {
		return err
	}

	listOfOrders, err := s.orderRepo.ReadOrders()
	if err != nil {
		return err
	}

	if len(listOfOrders) < 1 {
		return internal.ErrOrdersIsEmpty
	}

	for i, order := range listOfOrders {
		if order.ID == id {
			if order.Status == "open" {
				listOfOrders[i].Items = UpdatedOrder.Items
				listOfOrders[i].CustomerName = UpdatedOrder.CustomerName
				s.orderRepo.WriteOrders(listOfOrders)
				return nil
			} else {
				return internal.ErrOrderClosed
			}
		}
	}

	return internal.ErrOrderNotExist
}

func (s *orderService) Delete(r *http.Request) error {
	splitedURL := strings.Split(r.URL.Path, "/")
	id := splitedURL[len(splitedURL)-1]

	listOfOrders, err := s.orderRepo.ReadOrders()
	if err != nil {
		return err
	}

	if len(listOfOrders) < 1 {
		return internal.ErrOrdersIsEmpty
	}

	for i, order := range listOfOrders {
		if order.ID == id {
			if order.Status == "closed" {
				listOfOrders = append(listOfOrders[:i], listOfOrders[i+1:]...)
				s.orderRepo.WriteOrders(listOfOrders)
				return nil
			} else {

				listOfMenuItems, err := s.orderRepo.ReadMenu()
				if err != nil {
					return err
				}

				listOfInventoryItems, err := s.orderRepo.ReadInventory()
				if err != nil {
					return err
				}
				for _, orderItem := range order.Items {
					for _, menuItem := range listOfMenuItems {
						if orderItem.ProductID == menuItem.ID {
							for _, ingredientItem := range menuItem.Ingredients {
								for i, ingredientItemInv := range listOfInventoryItems {
									if ingredientItem.IngredientID == ingredientItemInv.IngredientID {
										if ingredientItem.Quantity*float64(
											orderItem.Quantity,
										) < ingredientItemInv.Quantity {
											listOfInventoryItems[i].Quantity += ingredientItem.Quantity * float64(
												orderItem.Quantity,
											)
										}
									}
								}
							}
						}
					}
				}

				listOfOrders = append(listOfOrders[:i], listOfOrders[i+1:]...)
				s.orderRepo.WriteInventory(listOfInventoryItems)
				s.orderRepo.WriteOrders(listOfOrders)
				return nil
			}
		}
	}
	return internal.ErrOrderNotExist
}

func (s *orderService) Close(r *http.Request) error {
	splitedURL := strings.Split(r.URL.Path, "/")
	id := splitedURL[len(splitedURL)-2]

	listOfOrders, err := s.orderRepo.ReadOrders()
	if err != nil {
		return err
	}

	if len(listOfOrders) < 1 {
		return internal.ErrOrdersIsEmpty
	}

	for i, order := range listOfOrders {
		if order.ID == id {
			if order.Status == "open" {
				listOfOrders[i].Status = "closed"
				s.orderRepo.WriteOrders(listOfOrders)
				return nil
			} else {
				return internal.ErrOrderClosed
			}
		}
	}
	return internal.ErrOrderNotExist
}

func (s *orderService) Test() {
	a, _ := s.orderRepo.ReadOrders()
	s.orderRepo.WriteOrders(a)
}
