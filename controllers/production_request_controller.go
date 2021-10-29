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
	Approve(ctx echo.Context) error
}
type ProductionRequestControllerImpl struct {
		ProductionRequestService services.ProductionRequestService
}

func NewProductionRequestController(productionRequestService services.ProductionRequestService) ProductionRequestController {
	return &ProductionRequestControllerImpl{
		ProductionRequestService: productionRequestService,
	}
}

func (o *ProductionRequestControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	productionRequestResponse, _ := o.ProductionRequestService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(productionRequestResponse.Code, productionRequestResponse)
}

func (o *ProductionRequestControllerImpl) FindAll(ctx echo.Context) error {
	productionRequestResponse, _ := o.ProductionRequestService.FindAll(ctx)

	return ctx.JSON(productionRequestResponse.Code, productionRequestResponse)
}

func (o *ProductionRequestControllerImpl) Create(ctx echo.Context) error {
	productionRequestResponse, _ := o.ProductionRequestService.Create(ctx)

	return ctx.JSON(productionRequestResponse.Code, productionRequestResponse)
}

func (o *ProductionRequestControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	productionRequestResponse, _ := o.ProductionRequestService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(productionRequestResponse.Code, productionRequestResponse)
}

func (o *ProductionRequestControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	productionRequestResponse, _ := o.ProductionRequestService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(productionRequestResponse.Code, productionRequestResponse)
}

func (o *ProductionRequestControllerImpl) Datatable(ctx echo.Context) error {
	productionRequestResponse, _ := o.ProductionRequestService.Datatable(ctx)

	return ctx.JSON(202, productionRequestResponse)
}

func (o *ProductionRequestControllerImpl) GeneratePdf(ctx echo.Context) error {
	productionRequestId := ctx.Param("id")
	productionRequestResponse, _ := o.ProductionRequestService.GeneratePdf(ctx, helpers.StringToInt(productionRequestId))
	return ctx.JSON(202, productionRequestResponse)
}

func (o *ProductionRequestControllerImpl) Approve(ctx echo.Context) error {
	productionRequestId := ctx.Param("id")
	productionRequestResponse, _ := o.ProductionRequestService.Approve(ctx, helpers.StringToInt(productionRequestId))

	return ctx.JSON(productionRequestResponse.Code, productionRequestResponse)
}