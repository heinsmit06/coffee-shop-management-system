package service

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type OrderServiceInterface interface {
	Create(*http.Request) error
	GetAll() ([]models.Order, error)
	GetOne() (models.Order, error)
	Update(models.Order) error
	Delete(models.Order) error
	Close(models.Order) error
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

	var newOrder models.Order

	err = json.NewDecoder(r.Body).Decode(&newOrder)

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
	s.orderRepo.WriteOrders(listOfOrders)
	return nil
}

func (s *orderService) GetAll() ([]models.Order, error) {
	return s.orderRepo.ReadOrders()
}

func (s *orderService) GetOne() (models.Order, error) {
	var order models.Order
	return order, nil
}

func (s *orderService) Update(order models.Order) error {
	return nil
}

func (s *orderService) Delete(order models.Order) error {
	return nil
}

func (s *orderService) Close(order models.Order) error {
	return nil
}
