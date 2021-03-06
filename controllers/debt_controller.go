package controllers

import (
	"kalika-be/helpers"
	"kalika-be/services"

	"github.com/labstack/echo"
	//"net/http"
)

type DebtController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
	ReportDatatable(ctx echo.Context) error
	GeneratePdf(ctx echo.Context) error
}
type DebtControllerImpl struct {
		DebtService services.DebtService
}

func NewDebtController(debtService services.DebtService) DebtController {
	return &DebtControllerImpl{
		DebtService: debtService,
	}
}

func (dc *DebtControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	debtResponse, _ := dc.DebtService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(debtResponse.Code, debtResponse)
}

func (dc *DebtControllerImpl) FindAll(ctx echo.Context) error {
	debtResponse, _ := dc.DebtService.FindAll(ctx)

	return ctx.JSON(debtResponse.Code, debtResponse)
}

func (dc *DebtControllerImpl) Create(ctx echo.Context) error {
	debtResponse, _ := dc.DebtService.Create(ctx)

	return ctx.JSON(debtResponse.Code, debtResponse)
}

func (dc *DebtControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	debtResponse, _ := dc.DebtService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(debtResponse.Code, debtResponse)
}

func (dc *DebtControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	debtResponse, _ := dc.DebtService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(debtResponse.Code, debtResponse)
}

func (dc *DebtControllerImpl) Datatable(ctx echo.Context) error {
	debtResponse, _ := dc.DebtService.Datatable(ctx)
	return ctx.JSON(202, debtResponse)
}
func (dc *DebtControllerImpl) ReportDatatable(ctx echo.Context) error {
	debtResponse, _ := dc.DebtService.ReportDatatable(ctx)
	return ctx.JSON(202, debtResponse)
}

func (dc *DebtControllerImpl) GeneratePdf(ctx echo.Context) error {
	debtResponse, _ := dc.DebtService.GeneratePdf(ctx)
	// filename := "laporan-hutang.pdf"
	// ctx.Response().Header().Set("Content-type", "application/octet-stream")
	// ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
	// return ctx.Attachment("exports/"+filename, filename)
	return ctx.JSON(202, debtResponse)
}