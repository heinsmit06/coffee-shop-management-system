package handler

import (
	"net/http"

	"hot-coffee/internal"
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
	internal.Logger.Info("GetTotalSales called", "method", "GetTotalSales")

	w.Header().Set("Content-type", "application/json")
	responseData, err := h.reportsService.GetTotalSales()
	if err != nil {
		internal.Logger.Error("Failed to retrieve total sales", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}

	w.Write(responseData)
	internal.Logger.Info("Total sales data sent as JSON response", "status", http.StatusOK)
}

func (h *reportsHandler) GetPopularItems(w http.ResponseWriter, r *http.Request) {
	internal.Logger.Info("GetPopularItems called", "method", "GetPopularItems")

	w.Header().Set("Content-type", "application/json")
	responseData, err := h.reportsService.GetMostPopular()
	if err != nil {
		internal.Logger.Error("Failed to retrieve popular items", "error", err)
		utils.ResponseErrorJson(err, w)
		return
	}

	w.Write(responseData)
	internal.Logger.Info("Popular items data sent as JSON response", "status", http.StatusOK)
}
