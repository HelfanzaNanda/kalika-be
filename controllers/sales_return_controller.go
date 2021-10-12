package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type SalesReturnController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
	ReportDatatable(ctx echo.Context) error
	GeneratePdf(ctx echo.Context) error
}
type SalesReturnControllerImpl struct {
		SalesReturnService services.SalesReturnService
}

func NewSalesReturnController(salesReturnService services.SalesReturnService) SalesReturnController {
	return &SalesReturnControllerImpl{
		SalesReturnService: salesReturnService,
	}
}

func (dc *SalesReturnControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	salesReturnResponse, _ := dc.SalesReturnService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesReturnResponse.Code, salesReturnResponse)
}

func (dc *SalesReturnControllerImpl) FindAll(ctx echo.Context) error {
	salesReturnResponse, _ := dc.SalesReturnService.FindAll(ctx)

	return ctx.JSON(salesReturnResponse.Code, salesReturnResponse)
}

func (dc *SalesReturnControllerImpl) Create(ctx echo.Context) error {
	salesReturnResponse, _ := dc.SalesReturnService.Create(ctx)

	return ctx.JSON(salesReturnResponse.Code, salesReturnResponse)
}

func (dc *SalesReturnControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	salesReturnResponse, _ := dc.SalesReturnService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesReturnResponse.Code, salesReturnResponse)
}

func (dc *SalesReturnControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	salesReturnResponse, _ := dc.SalesReturnService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(salesReturnResponse.Code, salesReturnResponse)
}

func (dc *SalesReturnControllerImpl) ReportDatatable(ctx echo.Context) error {
	salesReturnResponse, _ := dc.SalesReturnService.ReportDatatable(ctx)
	return ctx.JSON(202, salesReturnResponse)
}
func (dc *SalesReturnControllerImpl) Datatable(ctx echo.Context) error {
	salesReturnResponse, _ := dc.SalesReturnService.Datatable(ctx)
	return ctx.JSON(202, salesReturnResponse)
}

func (dc *SalesReturnControllerImpl) GeneratePdf(ctx echo.Context) error {
	saleReturnResponse, _ := dc.SalesReturnService.GeneratePdf(ctx)
	return ctx.JSON(202, saleReturnResponse)
}