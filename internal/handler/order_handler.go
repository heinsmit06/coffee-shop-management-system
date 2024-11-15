package handler

import (
	"net/http"

	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
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
		utils.ResponseErrorJson(err, w)
		return
	}
	w.WriteHeader(201)
}

func (h *orderHandler) RetrieveAllOrders(w http.ResponseWriter, r *http.Request) {
	// h.orderService.Test()
	orderData, err := h.orderService.GetAll()
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	w.Write(orderData)
}

func (h *orderHandler) RetrieveSpecificOrder(w http.ResponseWriter, r *http.Request) {
	orderData, err := h.orderService.GetOne(r)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	w.Write(orderData)
}

func (h *orderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	err := h.orderService.Update(r)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	w.WriteHeader(201)
}

func (h *orderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	err := h.orderService.Delete(r)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	w.WriteHeader(204)
}

func (h *orderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
	err := h.orderService.Close(r)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	w.WriteHeader(201)
}
