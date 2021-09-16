package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type CustomOrderController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type CustomOrderControllerImpl struct {
		CustomOrderService services.CustomOrderService
}

func NewCustomOrderController(customOrderService services.CustomOrderService) CustomOrderController {
	return &CustomOrderControllerImpl{
		CustomOrderService: customOrderService,
	}
}

func (dc *CustomOrderControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	customOrderResponse, _ := dc.CustomOrderService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(customOrderResponse.Code, customOrderResponse)
}

func (dc *CustomOrderControllerImpl) FindAll(ctx echo.Context) error {
	customOrderResponse, _ := dc.CustomOrderService.FindAll(ctx)

	return ctx.JSON(customOrderResponse.Code, customOrderResponse)
}

func (dc *CustomOrderControllerImpl) Create(ctx echo.Context) error {
	customOrderResponse, _ := dc.CustomOrderService.Create(ctx)

	return ctx.JSON(customOrderResponse.Code, customOrderResponse)
}

func (dc *CustomOrderControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	customOrderResponse, _ := dc.CustomOrderService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(customOrderResponse.Code, customOrderResponse)
}

func (dc *CustomOrderControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	customOrderResponse, _ := dc.CustomOrderService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(customOrderResponse.Code, customOrderResponse)
}