package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type PurchaseOrderController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Datatable(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	ReportDatatable(ctx echo.Context) error
	GeneratePdf(ctx echo.Context) error
}
type PurchaseOrderControllerImpl struct {
		PurchaseOrderService services.PurchaseOrderService
}

func NewPurchaseOrderController(purchaseOrderService services.PurchaseOrderService) PurchaseOrderController {
	return &PurchaseOrderControllerImpl{
		PurchaseOrderService: purchaseOrderService,
	}
}

func (dc *PurchaseOrderControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseOrderResponse, _ := dc.PurchaseOrderService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseOrderResponse.Code, purchaseOrderResponse)
}

func (dc *PurchaseOrderControllerImpl) FindAll(ctx echo.Context) error {
	purchaseOrderResponse, _ := dc.PurchaseOrderService.FindAll(ctx)

	return ctx.JSON(purchaseOrderResponse.Code, purchaseOrderResponse)
}

func (dc *PurchaseOrderControllerImpl) Create(ctx echo.Context) error {
	purchaseOrderResponse, _ := dc.PurchaseOrderService.Create(ctx)

	return ctx.JSON(purchaseOrderResponse.Code, purchaseOrderResponse)
}

func (dc *PurchaseOrderControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	purchaseOrderResponse, _ := dc.PurchaseOrderService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseOrderResponse.Code, purchaseOrderResponse)
}

func (dc *PurchaseOrderControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseOrderResponse, _ := dc.PurchaseOrderService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(purchaseOrderResponse.Code, purchaseOrderResponse)
}

func (dc *PurchaseOrderControllerImpl) Datatable(ctx echo.Context) error {
	purchaseOrderResponse, _ := dc.PurchaseOrderService.Datatable(ctx)
	return ctx.JSON(202, purchaseOrderResponse)
}

func (dc *PurchaseOrderControllerImpl) ReportDatatable(ctx echo.Context) error {
	purchaseOrderResponse, _ := dc.PurchaseOrderService.ReportDatatable(ctx)
	return ctx.JSON(202, purchaseOrderResponse)
}

func (dc *PurchaseOrderControllerImpl) GeneratePdf(ctx echo.Context) error {
	purchaseOrderResponse, _ := dc.PurchaseOrderService.GeneratePdf(ctx)
	return ctx.JSON(202, purchaseOrderResponse)
}