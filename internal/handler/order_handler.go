package handler

import (
	"net/http"

	"hot-coffee/internal"
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
	internal.Logger.Info("CreateNewOrder called", "method", "CreateNewOrder")

	err := h.orderService.Create(r)
	if err != nil {
		internal.Logger.Error("Failed to create new order", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}
	w.WriteHeader(201)
	internal.Logger.Info("Order created successfully", "status", http.StatusCreated)
}

func (h *orderHandler) RetrieveAllOrders(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("RetrieveAllOrders called", "method", "RetrieveAllOrders")

	w.Header().Set("Content-type", "application/json")
	// h.orderService.Test()
	orderData, err := h.orderService.GetAll()
	if err != nil {
		internal.Logger.Error("Failed to retrieve all orders", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}
	w.Write(orderData)
	internal.Logger.Info("All orders sent as JSON response", "status", http.StatusOK)
}

func (h *orderHandler) RetrieveSpecificOrder(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("RetrieveSpecificOrder called", "method", "RetrieveSpecificOrder")

	w.Header().Set("Content-type", "application/json")
	orderData, err := h.orderService.GetOne(r)
	if err != nil {
		internal.Logger.Error("Failed to retrieve specific order", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}

	w.Write(orderData)
	internal.Logger.Info("Specific order sent as JSON response", "status", http.StatusOK)
}

func (h *orderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("UpdateOrder called", "method", "UpdateOrder")

	err := h.orderService.Update(r)
	if err != nil {
		internal.Logger.Error("Failed to update order", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}

	w.WriteHeader(201)
	internal.Logger.Info("Order updated successfully", "status", http.StatusCreated)
}

func (h *orderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("DeleteOrder called", "method", "DeleteOrder")

	err := h.orderService.Delete(r)
	if err != nil {
		internal.Logger.Error("Failed to delete order", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}

	w.WriteHeader(204)
	internal.Logger.Info("Order deleted successfully", "status", http.StatusNoContent)
}

func (h *orderHandler) CloseOrder(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("CloseOrder called", "method", "CloseOrder")

	err := h.orderService.Close(r)
	if err != nil {
		internal.Logger.Error("Failed to close order", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}

	w.WriteHeader(201)
	internal.Logger.Info("Order closed successfully", "status", http.StatusCreated)
}
