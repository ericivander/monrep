package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/ericivander/monrep/flow"
	"github.com/labstack/echo/v4"
)

// ReportHandler struct holds flow needed for report handler
type ReportHandler struct {
	reportFlow flow.ReportFlow
}

// NewReportHandler construct ReportHandler
func NewReportHandler(reportFlow flow.ReportFlow) *ReportHandler {
	return &ReportHandler{reportFlow: reportFlow}
}

// GetReports handle request report handler
func (r *ReportHandler) GetReports(c echo.Context) error {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()

	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil {
		year = currentYear
	}
	if year < 1 || year > 9999 {
		return errors.New("invalid year")
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil {
		month = int(currentMonth)
	}
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}

	report, err := r.reportFlow.GetMonthlyReport(month, year)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, report)
}
