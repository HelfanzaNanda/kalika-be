package controllers

import (
	"kalika-be/services"

	"github.com/labstack/echo"
)

type CheckStockController interface {
	Datatable(ctx echo.Context) error
	GeneratePdf(ctx echo.Context) error
}
type CheckStockControllerImpl struct {
	CheckStockService services.CheckStockService
}

func NewCheckStockController(CheckStockService services.CheckStockService) CheckStockController {
	return &CheckStockControllerImpl{
		CheckStockService: CheckStockService,
	}
}

func (dc *CheckStockControllerImpl) Datatable(ctx echo.Context) error {
	checkstockResponse, _ := dc.CheckStockService.Datatable(ctx)
	return ctx.JSON(202, checkstockResponse)
}

func (dc *CheckStockControllerImpl) GeneratePdf(ctx echo.Context) error {
	checkstockResponse, _ := dc.CheckStockService.GeneratePdf(ctx)
	return ctx.JSON(202, checkstockResponse)
}
