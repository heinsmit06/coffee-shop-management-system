package handler

import (
	"net/http"

	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
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
		utils.ResponseErrorJson(err, w)
		return
	}
	w.WriteHeader(201)
	// fmt.Println("Ingredient successfully added")
	return
}

func (h *inventoryHandler) RetrieveAllInventory(w http.ResponseWriter, r *http.Request) {
	bodyData, err := h.inventoryService.GetAll()
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	w.Write(bodyData)
	return
}

func (h *inventoryHandler) RetrieveSpecificInventory(w http.ResponseWriter, r *http.Request) {
	bodyData, err := h.inventoryService.GetOne(r)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	w.Write(bodyData)
	return
}

func (h *inventoryHandler) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	err := h.inventoryService.Update(r)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	w.WriteHeader(200)
	return
}

func (h *inventoryHandler) DeleteInventory(w http.ResponseWriter, r *http.Request) {
	err := h.inventoryService.Delete(r)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	w.WriteHeader(204)
	return
}
