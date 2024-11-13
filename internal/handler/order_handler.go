package handler

import (
	"fmt"
	"net/http"

	"hot-coffee/internal/service"
)

type orderHandler struct {
	orderService service.OrderServiceInterface
}

func NewOrderHandler(orderService service.OrderServiceInterface) *orderHandler {
	return &orderHandler{orderService: orderService}
}

func (h *orderHandler) CreateNewOrder(w http.ResponseWriter, r *http.Request) {
	err := h.orderService.Create(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("createHandlerWorks")
}

func (h *orderHandler) RetrieveAllOrders(w http.ResponseWriter, r *http.Request) {
}

func (h *orderHandler) RetrieveSpecificOrder(w http.ResponseWriter, r *http.Request) {
}

func (h *orderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
}

func (h *orderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
}

func (h *orderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
}
