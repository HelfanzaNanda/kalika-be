package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type PurchaseInvoiceController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type PurchaseInvoiceControllerImpl struct {
		PurchaseInvoiceService services.PurchaseInvoiceService
}

func NewPurchaseInvoiceController(purchaseInvoiceService services.PurchaseInvoiceService) PurchaseInvoiceController {
	return &PurchaseInvoiceControllerImpl{
		PurchaseInvoiceService: purchaseInvoiceService,
	}
}

func (dc *PurchaseInvoiceControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	paymentInvoiceResponse, _ := dc.PurchaseInvoiceService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(paymentInvoiceResponse.Code, paymentInvoiceResponse)
}

func (dc *PurchaseInvoiceControllerImpl) FindAll(ctx echo.Context) error {
	paymentInvoiceResponse, _ := dc.PurchaseInvoiceService.FindAll(ctx)

	return ctx.JSON(paymentInvoiceResponse.Code, paymentInvoiceResponse)
}

func (dc *PurchaseInvoiceControllerImpl) Create(ctx echo.Context) error {
	paymentInvoiceResponse, _ := dc.PurchaseInvoiceService.Create(ctx)

	return ctx.JSON(paymentInvoiceResponse.Code, paymentInvoiceResponse)
}

func (dc *PurchaseInvoiceControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	paymentInvoiceResponse, _ := dc.PurchaseInvoiceService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(paymentInvoiceResponse.Code, paymentInvoiceResponse)
}

func (dc *PurchaseInvoiceControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	paymentInvoiceResponse, _ := dc.PurchaseInvoiceService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(paymentInvoiceResponse.Code, paymentInvoiceResponse)
}