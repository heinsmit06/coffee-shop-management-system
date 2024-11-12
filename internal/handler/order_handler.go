package handler

import (
	"net/http"

	"hot-coffee/internal/service"
)

type orderHandler struct {
	orderService service.OrderServiceInterface
}

func NewOrderHandler(orderService service.OrderServiceInterface) *orderHandler {
	return &orderHandler{orderService: orderService}
}

func CreateNewOrder(w http.ResponseWriter, r *http.Request) {
}

func RetrieveAllOrders(w http.ResponseWriter, r *http.Request) {
}

func RetrieveSpecificOrder(w http.ResponseWriter, r *http.Request) {
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
}

func CloseOrder(w http.ResponseWriter, r *http.Request) {
}
