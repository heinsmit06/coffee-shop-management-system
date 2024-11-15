package service

import (
	"encoding/json"

	"hot-coffee/internal"
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type ReportsServerInterface interface {
	GetTotalSales() ([]byte, error)
	GetMostPopular() ([]byte, error)
}

type reportsServer struct {
	reportsRepo dal.ReportsRepoInterface
}

func NewReportsServer(reportsRepo dal.ReportsRepoInterface) *reportsServer {
	return &reportsServer{reportsRepo: reportsRepo}
}

func (s *reportsServer) GetTotalSales() ([]byte, error) {
	listOfOrders, err := s.reportsRepo.ReadOrders()
	if err != nil {
		return []byte{}, err
	}

	listOfMenuItems, err := s.reportsRepo.ReadMenu()
	if err != nil {
		return []byte{}, err
	}

	if len(listOfOrders) < 1 {
		return []byte{}, internal.ErrOrdersIsEmpty
	}

	if len(listOfMenuItems) < 1 {
		return []byte{}, internal.ErrMenuIsEmpty
	}

	menuItemPriceMap := make(map[string]float64)
	var totalSales float64

	for _, menuItem := range listOfMenuItems {
		menuItemPriceMap[menuItem.ID] = menuItem.Price
	}

	for _, order := range listOfOrders {
		if order.Status == "closed" {
			for _, orderItem := range order.Items {
				totalSales += menuItemPriceMap[orderItem.ProductID] * float64(orderItem.Quantity)
			}
		}
	}

	jsonData, err := json.MarshalIndent(models.TotalSales{Amount: totalSales}, "", " ")

	return jsonData, nil
}

func (s *reportsServer) GetMostPopular() ([]byte, error) {
	listOfOrders, err := s.reportsRepo.ReadOrders()
	if err != nil {
		return []byte{}, err
	}

	if len(listOfOrders) < 1 {
		return []byte{}, internal.ErrOrdersIsEmpty
	}

	mapItemsOrderedN := make(map[string]int)

	for _, order := range listOfOrders {
		for _, orderItems := range order.Items {
			mapItemsOrderedN[orderItems.ProductID] += orderItems.Quantity
		}
	}

	var mostPopularCounter int

	for _, v := range mapItemsOrderedN {
		if v > mostPopularCounter {
			mostPopularCounter = v
		}
	}

	mostPopularList := []models.MostPopularItem{}

	for key, value := range mapItemsOrderedN {
		if value == mostPopularCounter {
			mostPopularList = append(
				mostPopularList,
				models.MostPopularItem{Item: key, Quantity: value},
			)
		}
	}

	jsonData, err := json.MarshalIndent(mostPopularList, "", " ")
	if err != nil {
		return []byte{}, err
	}

	return jsonData, nil
}
