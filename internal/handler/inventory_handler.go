package handler

import (
	"fmt"
	"net/http"

	"hot-coffee/internal/service"
)

type inventoryHandler struct {
	inventoryService service.InventoryServiceInterface
}

func NewInventoryHandler(inventoryService service.InventoryServiceInterface) *inventoryHandler {
	return &inventoryHandler{inventoryService: inventoryService}
}

func (h *inventoryHandler) AddNewInventory(w http.ResponseWriter, r *http.Request) {
	err := h.inventoryService.AddInventory(r)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *inventoryHandler) RetrieveAllInventory(w http.ResponseWriter, r *http.Request) {
}

func (h *inventoryHandler) RetrieveSpecificInventory(w http.ResponseWriter, r *http.Request) {
}

func (h *inventoryHandler) UpdateInventory(w http.ResponseWriter, r *http.Request) {
}

func (h *inventoryHandler) DeleteInventory(w http.ResponseWriter, r *http.Request) {
}
