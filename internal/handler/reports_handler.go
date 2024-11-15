package handler

import (
	"net/http"

	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
)

type reportsHandler struct {
	reportsService service.ReportsServerInterface
}

func NewReportHandler(reportsService service.ReportsServerInterface) *reportsHandler {
	return &reportsHandler{reportsService: reportsService}
}

func (h *reportsHandler) GetTotalSales(w http.ResponseWriter, r *http.Request) {
	responseData, err := h.reportsService.GetTotalSales()
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	w.Write(responseData)
}

func (h *reportsHandler) GetPopularItems(w http.ResponseWriter, r *http.Request) {
	responseData, err := h.reportsService.GetMostPopular()
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	w.Write(responseData)
}
