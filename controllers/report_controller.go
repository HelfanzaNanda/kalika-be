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
	GeneratePdfProfitLoss(ctx echo.Context) error
	GeneratePdfReceivableLedger(ctx echo.Context) error
	GeneratePdfDebtLedger(ctx echo.Context) error
	GeneratePdfCashBankLedger(ctx echo.Context) error
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

func (dc *ReportControllerImpl) GeneratePdfProfitLoss(ctx echo.Context) error {
	reportResponse, _ := dc.ReportService.GeneratePdfProfitLoss(ctx)
	return ctx.JSON(202, reportResponse)
}

func (dc *ReportControllerImpl) GeneratePdfReceivableLedger(ctx echo.Context) error {
	reportResponse, _ := dc.ReportService.GeneratePdfReceivableLedger(ctx)
	return ctx.JSON(202, reportResponse)
}
func (dc *ReportControllerImpl) GeneratePdfDebtLedger(ctx echo.Context) error {
	reportResponse, _ := dc.ReportService.GeneratePdfDebtLedger(ctx)
	return ctx.JSON(202, reportResponse)
}
func (dc *ReportControllerImpl) GeneratePdfCashBankLedger(ctx echo.Context) error {
	reportResponse, _ := dc.ReportService.GeneratePdfCashBankLedger(ctx)
	return ctx.JSON(202, reportResponse)
}