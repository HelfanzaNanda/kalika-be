package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type PurchaseInvoiceDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type PurchaseInvoiceDetailControllerImpl struct {
		PurchaseInvoiceDetailService services.PurchaseInvoiceDetailService
}

func NewPurchaseInvoiceDetailController(purchaseInvoiceDetailService services.PurchaseInvoiceDetailService) PurchaseInvoiceDetailController {
	return &PurchaseInvoiceDetailControllerImpl{
		PurchaseInvoiceDetailService: purchaseInvoiceDetailService,
	}
}

func (dc *PurchaseInvoiceDetailControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseInvoiceDetailResponse, _ := dc.PurchaseInvoiceDetailService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseInvoiceDetailResponse.Code, purchaseInvoiceDetailResponse)
}

func (dc *PurchaseInvoiceDetailControllerImpl) FindAll(ctx echo.Context) error {
	purchaseInvoiceDetailResponse, _ := dc.PurchaseInvoiceDetailService.FindAll(ctx)

	return ctx.JSON(purchaseInvoiceDetailResponse.Code, purchaseInvoiceDetailResponse)
}

func (dc *PurchaseInvoiceDetailControllerImpl) Create(ctx echo.Context) error {
	purchaseInvoiceDetailResponse, _ := dc.PurchaseInvoiceDetailService.Create(ctx)

	return ctx.JSON(purchaseInvoiceDetailResponse.Code, purchaseInvoiceDetailResponse)
}

func (dc *PurchaseInvoiceDetailControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	purchaseInvoiceDetailResponse, _ := dc.PurchaseInvoiceDetailService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseInvoiceDetailResponse.Code, purchaseInvoiceDetailResponse)
}

func (dc *PurchaseInvoiceDetailControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseInvoiceDetailResponse, _ := dc.PurchaseInvoiceDetailService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(purchaseInvoiceDetailResponse.Code, purchaseInvoiceDetailResponse)
}