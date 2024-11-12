package service

import (
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type OrderHandlerInterface interface {
	Create(models.Order) error
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

func (s *orderService) Create(order models.Order) error {
	listOfOrders, err := s.orderRepo.ReadOrders()
	if err != nil {
		return err
	}
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
