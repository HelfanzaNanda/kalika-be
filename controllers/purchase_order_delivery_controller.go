package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type PurchaseOrderDeliveryController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type PurchaseOrderDeliveryControllerImpl struct {
		PurchaseOrderDeliveryService services.PurchaseOrderDeliveryService
}

func NewPurchaseOrderDeliveryController(purchaseOrderDeliveryService services.PurchaseOrderDeliveryService) PurchaseOrderDeliveryController {
	return &PurchaseOrderDeliveryControllerImpl{
		PurchaseOrderDeliveryService: purchaseOrderDeliveryService,
	}
}

func (dc *PurchaseOrderDeliveryControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseOrderDeliveryResponse, _ := dc.PurchaseOrderDeliveryService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseOrderDeliveryResponse.Code, purchaseOrderDeliveryResponse)
}

func (dc *PurchaseOrderDeliveryControllerImpl) FindAll(ctx echo.Context) error {
	purchaseOrderDeliveryResponse, _ := dc.PurchaseOrderDeliveryService.FindAll(ctx)

	return ctx.JSON(purchaseOrderDeliveryResponse.Code, purchaseOrderDeliveryResponse)
}

func (dc *PurchaseOrderDeliveryControllerImpl) Create(ctx echo.Context) error {
	purchaseOrderDeliveryResponse, _ := dc.PurchaseOrderDeliveryService.Create(ctx)

	return ctx.JSON(purchaseOrderDeliveryResponse.Code, purchaseOrderDeliveryResponse)
}

func (dc *PurchaseOrderDeliveryControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	purchaseOrderDeliveryResponse, _ := dc.PurchaseOrderDeliveryService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseOrderDeliveryResponse.Code, purchaseOrderDeliveryResponse)
}

func (dc *PurchaseOrderDeliveryControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseOrderDeliveryResponse, _ := dc.PurchaseOrderDeliveryService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(purchaseOrderDeliveryResponse.Code, purchaseOrderDeliveryResponse)
}