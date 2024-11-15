package handler

import (
	"fmt"
	"net/http"

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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// w.Write([]byte(err.Error()))

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Menu item added successfully")
}

func (h *menuHandler) RetrieveAllMenu(w http.ResponseWriter, r *http.Request) {
}

func (h *menuHandler) RetrieveSpecificMenu(w http.ResponseWriter, r *http.Request) {
}

func (h *menuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
}

func (h *menuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
}
