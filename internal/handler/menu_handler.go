package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"hot-coffee/internal"
	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
)

type menuHandler struct {
	menuService service.MenuServiceInterface // Dependency injection
}

func NewMenuHandler(menuService service.MenuServiceInterface) *menuHandler {
	return &menuHandler{menuService: menuService}
}

func (h *menuHandler) AddNewMenu(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("AddNewMenu called", "method", "AddNewMenu")

	if r.Header.Get("Content-type") != "application/json" {
		internal.Logger.Warn("Unsupported Media Type", "expected", "application/json", "received", r.Header.Get("Content-type"))
		utils.ResponseErrorJson(internal.ErrUnsupportedMediaType, w)
		return
	}

	err := h.menuService.AddMenu(r)
	if err != nil {
		internal.Logger.Error("Failed to add new menu item", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}
	// w.Write([]byte(err.Error()))

	w.WriteHeader(http.StatusCreated)
	internal.Logger.Info("Menu item added successfully", "status", http.StatusCreated)
}

func (h *menuHandler) RetrieveAllMenu(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("RetrieveAllMenu called", "method", "RetrieveAllMenu")

	w.Header().Set("Content-type", "application/json")
	allMenuItems, err := h.menuService.GetAll()
	if err != nil {
		internal.Logger.Error("Failed to retrieve all menu items", "error", err)

		http.Error(w, "Unable to show the menu", http.StatusInternalServerError)
		return
	}
	internal.Logger.Info("Successfully retrieved menu items", "count", len(allMenuItems))

	jsonData, err := json.MarshalIndent(allMenuItems, "", " ")
	if err != nil {
		internal.Logger.Error("Failed to marshal menu items to JSON", "error", err)
		http.Error(w, "Unable to marshal to show the menu", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
	internal.Logger.Info("Menu items sent as JSON response", "status", http.StatusOK)
}

func (h *menuHandler) RetrieveSpecificMenu(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("RetrieveSpecificMenu called", "method", "RetrieveSpecificMenu")

	w.Header().Set("Content-Type", "application/json")
	id := strings.Split(r.URL.Path[1:], "/")[1]
	internal.Logger.Info("Retrieving specific menu item", "menuItemID", id)

	menuItem, err := h.menuService.GetOne(id)
	if err != nil {
		internal.Logger.Error("Failed to retrieve specific menu item", "menuItemID", id, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	menuItemJson, err := json.MarshalIndent(menuItem, "", " ")
	if err != nil {
		internal.Logger.Error("Failed to marshal specific menu item to JSON", "menuItemID", id, "error", err)
		http.Error(w, "Failed to marshal specific menu item", http.StatusInternalServerError)
		return
	}

	w.Write(menuItemJson)
	internal.Logger.Info("Specific menu item sent as JSON response", "menuItemID", id, "status", http.StatusOK)
}

func (h *menuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("UpdateMenu called", "method", "UpdateMenu")

	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	internal.Logger.Info("Updating menu item", "menuItemID", id)

	err := h.menuService.Update(r, id)
	if err != nil {
		internal.Logger.Error("Failed to update menu item", "menuItemID", id, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	internal.Logger.Info("Menu item updated successfully", "menuItemID", id, "status", http.StatusOK)
}

func (h *menuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("DeleteMenu called", "method", "DeleteMenu")

	id := strings.Split(r.URL.Path[1:], "/")[1]
	internal.Logger.Info("Deleting menu item", "menuItemID", id)

	err := h.menuService.Delete(id)
	if err != nil {
		internal.Logger.Error("Failed to delete menu item", "menuItemID", id, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	internal.Logger.Info("Menu item deleted successfully", "menuItemID", id, "status", http.StatusNoContent)
}
