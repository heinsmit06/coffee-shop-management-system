package handler

import (
	"net/http"

	"hot-coffee/internal/service"
)

type reportsHandler struct {
	reportsService service.ReportsServerInterface
}

func NewReportHandler(reportsService service.ReportsServerInterface) *reportsHandler {
	return &reportsHandler{reportsService: reportsService}
}

func (h *reportsHandler) GetTotalSales(w http.ResponseWriter, r *http.Request) {
}

func (h *reportsHandler) GetPopularItems(w http.ResponseWriter, r *http.Request) {
}
