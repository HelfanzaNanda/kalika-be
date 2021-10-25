package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/services"
	//"net/http"
)

type ReportController interface {
	ProfitLoss(ctx echo.Context) error
	ReceivableLedger(ctx echo.Context) error
	DebtLedger(ctx echo.Context) error
	CashBankLedger(ctx echo.Context) error
}
type ReportControllerImpl struct {
	ReportService services.ReportService
}

func NewReportController(reportService services.ReportService) ReportController {
	return &ReportControllerImpl{
		ReportService: reportService,
	}
}

func (dc *ReportControllerImpl) ProfitLoss(ctx echo.Context) error {
	reportResponse, _ := dc.ReportService.ProfitLoss(ctx)

	return ctx.JSON(reportResponse.Code, reportResponse)
}

func (dc *ReportControllerImpl) ReceivableLedger(ctx echo.Context) error {
	reportResponse, _ := dc.ReportService.ReceivableLedger(ctx)

	return ctx.JSON(reportResponse.Code, reportResponse)
}

func (dc *ReportControllerImpl) DebtLedger(ctx echo.Context) error {
	reportResponse, _ := dc.ReportService.DebtLedger(ctx)

	return ctx.JSON(reportResponse.Code, reportResponse)
}

func (dc *ReportControllerImpl) CashBankLedger(ctx echo.Context) error {
	reportResponse, _ := dc.ReportService.CashBankLedger(ctx)

	return ctx.JSON(reportResponse.Code, reportResponse)
}