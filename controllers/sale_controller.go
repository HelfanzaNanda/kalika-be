package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type SaleController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
	ReportDatatable(ctx echo.Context) error
}
type SaleControllerImpl struct {
		SalesService services.SalesService
}

func NewSaleController(salesService services.SalesService) SaleController {
	return &SaleControllerImpl{
		SalesService: salesService,
	}
}

func (dc *SaleControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	salesResponse, _ := dc.SalesService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesResponse.Code, salesResponse)
}

func (dc *SaleControllerImpl) FindAll(ctx echo.Context) error {
	salesResponse, _ := dc.SalesService.FindAll(ctx)

	return ctx.JSON(salesResponse.Code, salesResponse)
}

func (dc *SaleControllerImpl) Create(ctx echo.Context) error {
	salesResponse, _ := dc.SalesService.Create(ctx)

	return ctx.JSON(salesResponse.Code, salesResponse)
}

func (dc *SaleControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	salesResponse, _ := dc.SalesService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesResponse.Code, salesResponse)
}

func (dc *SaleControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	salesResponse, _ := dc.SalesService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(salesResponse.Code, salesResponse)
}

func (dc *SaleControllerImpl) Datatable(ctx echo.Context) error {
	salesResponse, _ := dc.SalesService.Datatable(ctx)
	return ctx.JSON(202, salesResponse)
}

func (dc *SaleControllerImpl) ReportDatatable(ctx echo.Context) error {
	salesResponse, _ := dc.SalesService.ReportDatatable(ctx)
	return ctx.JSON(202, salesResponse)
}