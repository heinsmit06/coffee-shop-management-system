package handler

import (
	"net/http"

	"hot-coffee/internal/dal"
	"hot-coffee/internal/service"
)

func SetupServer(dirPath string) *http.ServeMux {
	mux := http.NewServeMux()

	menuRepo := dal.NewMenuRepo(dirPath + "/menu_item.json")
	menuService := service.NewMenuService(menuRepo)
	menuHandler := NewMenuHandler(menuService)

	orderRepo := dal.NewOrderRepo(dirPath + "/orders.json")
	orderService := service.NewOrderService(orderRepo)
	orderHandler := NewOrderHandler(orderService)

	// ORDER handling
	mux.HandleFunc("POST /order", orderHandler.CreateNewOrder)
	mux.HandleFunc("GET /order", orderHandler.RetrieveAllOrders)
	mux.HandleFunc("GET /order/{id}", orderHandler.RetrieveSpecificOrder)
	mux.HandleFunc("PUT /order/{id}", orderHandler.UpdateOrder)
	mux.HandleFunc("DELETE /order/{id}", orderHandler.DeleteOrder)
	mux.HandleFunc("POST /order/{id}/close", orderHandler.CloseOrder)

	// MENU Items handling
	mux.HandleFunc("POST /menu", menuHandler.AddNewMenu)
	mux.HandleFunc("GET /menu", menuHandler.RetrieveAllMenu)
	mux.HandleFunc("GET /menu/{id}", menuHandler.RetrieveSpecificMenu)
	mux.HandleFunc("PUT /menu/{id}", menuHandler.UpdateMenu)
	mux.HandleFunc("DELETE /menu", menuHandler.DeleteMenu)

	// INVENTORY handling
	mux.HandleFunc("POST /inventory", AddNewInventory)
	mux.HandleFunc("GET /inventory", RetrieveAllInventory)
	mux.HandleFunc("GET /inventory/{id}", RetrieveSpecificInventory)
	mux.HandleFunc("PUT /inventory/{id}", UpdateInventory)
	mux.HandleFunc("DELETE /inventory/{id}", DeleteInventory)

	// AGGREGATIONS handling

	return mux
}
