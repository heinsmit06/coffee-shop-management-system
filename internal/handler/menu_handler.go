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

func RetrieveAllMenu(w http.ResponseWriter, r *http.Request) {
}

func RetrieveSpecificMenu(w http.ResponseWriter, r *http.Request) {
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
}
