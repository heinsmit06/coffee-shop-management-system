package handler

import (
	"net/http"

	"hot-coffee/internal/dal"
	"hot-coffee/internal/service"
)

func SetupServer() *http.ServeMux {
	menuRepo := dal.NewMenuRepo("data/json")
	menuService := service.NewMenuService(menuRepo)
	menuHandler := NewMenuHandler(menuService)

	mux := http.NewServeMux()
	// ORDER handling
	mux.HandleFunc("POST /order", CreateNewOrder)
	mux.HandleFunc("GET /order", RetrieveAllOrders)
	mux.HandleFunc("GET /order/{id}", RetrieveSpecificOrder)
	mux.HandleFunc("PUT /order/{id}", UpdateOrder)
	mux.HandleFunc("DELETE /order/{id}", DeleteOrder)
	mux.HandleFunc("POST /order/{id}/close", CloseOrder)

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
