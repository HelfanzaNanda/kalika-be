package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type ProductionRequestController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
	GeneratePdf(ctx echo.Context) error
}
type ProductionRequestControllerImpl struct {
		ProductionRequestService services.ProductionRequestService
}

func NewProductionRequestController(stockOpnameService services.ProductionRequestService) ProductionRequestController {
	return &ProductionRequestControllerImpl{
		ProductionRequestService: stockOpnameService,
	}
}

func (o *ProductionRequestControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	stockOpnameResponse, _ := o.ProductionRequestService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *ProductionRequestControllerImpl) FindAll(ctx echo.Context) error {
	stockOpnameResponse, _ := o.ProductionRequestService.FindAll(ctx)

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *ProductionRequestControllerImpl) Create(ctx echo.Context) error {
	stockOpnameResponse, _ := o.ProductionRequestService.Create(ctx)

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *ProductionRequestControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	stockOpnameResponse, _ := o.ProductionRequestService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *ProductionRequestControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	stockOpnameResponse, _ := o.ProductionRequestService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *ProductionRequestControllerImpl) Datatable(ctx echo.Context) error {
	stockOpnameResponse, _ := o.ProductionRequestService.Datatable(ctx)

	return ctx.JSON(202, stockOpnameResponse)
}

func (o *ProductionRequestControllerImpl) GeneratePdf(ctx echo.Context) error {
	stockOpnameId := ctx.Param("id")
	stockOpnameResponse, _ := o.ProductionRequestService.GeneratePdf(ctx, helpers.StringToInt(stockOpnameId))
	return ctx.JSON(202, stockOpnameResponse)
}