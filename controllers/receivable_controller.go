package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type ReceivableController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
	ReportDatatable(ctx echo.Context) error
}

type ReceivableControllerImpl struct {
		ReceivableService services.ReceivableService
}

func NewReceivableController(receivableService services.ReceivableService) ReceivableController {
	return &ReceivableControllerImpl{
		ReceivableService: receivableService,
	}
}

func (dc *ReceivableControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	receivableResponse, _ := dc.ReceivableService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(receivableResponse.Code, receivableResponse)
}

func (dc *ReceivableControllerImpl) FindAll(ctx echo.Context) error {
	receivableResponse, _ := dc.ReceivableService.FindAll(ctx)

	return ctx.JSON(receivableResponse.Code, receivableResponse)
}

func (dc *ReceivableControllerImpl) Create(ctx echo.Context) error {
	receivableResponse, _ := dc.ReceivableService.Create(ctx)

	return ctx.JSON(receivableResponse.Code, receivableResponse)
}

func (dc *ReceivableControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	receivableResponse, _ := dc.ReceivableService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(receivableResponse.Code, receivableResponse)
}

func (dc *ReceivableControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	receivableResponse, _ := dc.ReceivableService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(receivableResponse.Code, receivableResponse)
}

func (dc *ReceivableControllerImpl) Datatable(ctx echo.Context) error {
	receivableResponse, _ := dc.ReceivableService.Datatable(ctx)
	return ctx.JSON(202, receivableResponse)
}

func (dc *ReceivableControllerImpl) ReportDatatable(ctx echo.Context) error {
	receivableResponse, _ := dc.ReceivableService.ReportDatatable(ctx)
	return ctx.JSON(202, receivableResponse)
}