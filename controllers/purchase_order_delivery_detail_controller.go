package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type PurchaseOrderDeliveryDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type PurchaseOrderDeliveryDetailControllerImpl struct {
		PurchaseOrderDeliveryDetailService services.PurchaseOrderDeliveryDetailService
}

func NewPurchaseOrderDeliveryDetailController(purchaseOrderDeliveryDetailService services.PurchaseOrderDeliveryDetailService) PurchaseOrderDeliveryDetailController {
	return &PurchaseOrderDeliveryDetailControllerImpl{
		PurchaseOrderDeliveryDetailService: purchaseOrderDeliveryDetailService,
	}
}

func (dc *PurchaseOrderDeliveryDetailControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseOrderDeliveryDetailResponse, _ := dc.PurchaseOrderDeliveryDetailService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseOrderDeliveryDetailResponse.Code, purchaseOrderDeliveryDetailResponse)
}

func (dc *PurchaseOrderDeliveryDetailControllerImpl) FindAll(ctx echo.Context) error {
	purchaseOrderDeliveryDetailResponse, _ := dc.PurchaseOrderDeliveryDetailService.FindAll(ctx)

	return ctx.JSON(purchaseOrderDeliveryDetailResponse.Code, purchaseOrderDeliveryDetailResponse)
}

func (dc *PurchaseOrderDeliveryDetailControllerImpl) Create(ctx echo.Context) error {
	purchaseOrderDeliveryDetailResponse, _ := dc.PurchaseOrderDeliveryDetailService.Create(ctx)

	return ctx.JSON(purchaseOrderDeliveryDetailResponse.Code, purchaseOrderDeliveryDetailResponse)
}

func (dc *PurchaseOrderDeliveryDetailControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	purchaseOrderDeliveryDetailResponse, _ := dc.PurchaseOrderDeliveryDetailService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseOrderDeliveryDetailResponse.Code, purchaseOrderDeliveryDetailResponse)
}

func (dc *PurchaseOrderDeliveryDetailControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseOrderDeliveryDetailResponse, _ := dc.PurchaseOrderDeliveryDetailService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(purchaseOrderDeliveryDetailResponse.Code, purchaseOrderDeliveryDetailResponse)
}