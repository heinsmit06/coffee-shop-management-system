package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"hot-coffee/internal/service"
)

type menuHandler struct {
	menuService service.MenuServiceInterface // Dependency injection
}

func NewMenuHandler(menuService service.MenuServiceInterface) *menuHandler {
	return &menuHandler{menuService: menuService}
}

func (h *menuHandler) AddNewMenu(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	err := h.menuService.AddMenu(r)
	if err != nil {
		fmt.Println("abc")
		http.Error(w, "error:"+err.Error(), http.StatusBadRequest)
		return
	}
	// w.Write([]byte(err.Error()))

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Menu item added successfully")
}

func (h *menuHandler) RetrieveAllMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	allMenuItems, err := h.menuService.GetAll()
	if err != nil {
		http.Error(w, "Unable to show the menu", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(allMenuItems, "", " ")
	if err != nil {
		http.Error(w, "Unable to marshal to show the menu", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (h *menuHandler) RetrieveSpecificMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.Split(r.URL.Path[1:], "/")[1]
	menuItem, err := h.menuService.GetOne(id)
	if err != nil {
		http.Error(w, "Failed getting one menu item", http.StatusInternalServerError)
		return
	}

	menuItemJson, err := json.MarshalIndent(menuItem, "", " ")
	if err != nil {
		http.Error(w, "Failed to marshal one menu item", http.StatusInternalServerError)
		return
	}

	w.Write(menuItemJson)
}

func (h *menuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.Split(r.URL.Path[1:], "/")[1]

	err := h.menuService.Update(r, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *menuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path[1:], "/")[1]

	err := h.menuService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
