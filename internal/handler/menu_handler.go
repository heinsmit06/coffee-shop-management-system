package handler

import (
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
}

func (h *menuHandler) RetrieveAllMenu(w http.ResponseWriter, r *http.Request) {
}

func (h *menuHandler) RetrieveSpecificMenu(w http.ResponseWriter, r *http.Request) {
}

func (h *menuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
}

func (h *menuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
}
