package handler

import (
	"net/http"

	"hot-coffee/internal"
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
	internal.Logger.Info("AddNewInventory called", "method", "AddNewInventory")

	err := h.inventoryService.AddInventory(r)
	if err != nil {
		internal.Logger.Error("Failed to add new inventory item", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}
	w.WriteHeader(201)
	// fmt.Println("Ingredient successfully added")
	internal.Logger.Info("Inventory item added successfully", "status", http.StatusCreated)
	return
}

func (h *inventoryHandler) RetrieveAllInventory(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("RetrieveAllInventory called", "method", "RetrieveAllInventory")

	w.Header().Set("Content-type", "application/json")
	bodyData, err := h.inventoryService.GetAll()
	if err != nil {
		internal.Logger.Error("Failed to retrieve all inventory items", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}

	w.Write(bodyData)
	internal.Logger.Info("All inventory items sent as JSON response", "status", http.StatusOK)
}

func (h *inventoryHandler) RetrieveSpecificInventory(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("RetrieveSpecificInventory called", "method", "RetrieveSpecificInventory")

	w.Header().Set("Content-type", "application/json")
	bodyData, err := h.inventoryService.GetOne(r)
	if err != nil {
		internal.Logger.Error("Failed to retrieve specific inventory item", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}

	w.Write(bodyData)
	internal.Logger.Info("Specific inventory item sent as JSON response", "status", http.StatusOK)
}

func (h *inventoryHandler) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("UpdateInventory called", "method", "UpdateInventory")

	err := h.inventoryService.Update(r)
	if err != nil {
		internal.Logger.Error("Failed to update inventory item", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	internal.Logger.Info("Inventory item updated successfully", "status", http.StatusOK)
}

func (h *inventoryHandler) DeleteInventory(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("DeleteInventory called", "method", "DeleteInventory")

	err := h.inventoryService.Delete(r)
	if err != nil {
		internal.Logger.Error("Failed to delete inventory item", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	internal.Logger.Info("Inventory item deleted successfully", "status", http.StatusNoContent)
}
