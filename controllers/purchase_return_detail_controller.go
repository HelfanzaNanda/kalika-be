package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type PurchaseReturnDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type PurchaseReturnDetailControllerImpl struct {
		PurchaseReturnDetailService services.PurchaseReturnDetailService
}

func NewPurchaseReturnDetailController(purchaseReturnDetailService services.PurchaseReturnDetailService) PurchaseReturnDetailController {
	return &PurchaseReturnDetailControllerImpl{
		PurchaseReturnDetailService: purchaseReturnDetailService,
	}
}

func (dc *PurchaseReturnDetailControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseReturnDetailResponse, _ := dc.PurchaseReturnDetailService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseReturnDetailResponse.Code, purchaseReturnDetailResponse)
}

func (dc *PurchaseReturnDetailControllerImpl) FindAll(ctx echo.Context) error {
	purchaseReturnDetailResponse, _ := dc.PurchaseReturnDetailService.FindAll(ctx)

	return ctx.JSON(purchaseReturnDetailResponse.Code, purchaseReturnDetailResponse)
}

func (dc *PurchaseReturnDetailControllerImpl) Create(ctx echo.Context) error {
	purchaseReturnDetailResponse, _ := dc.PurchaseReturnDetailService.Create(ctx)

	return ctx.JSON(purchaseReturnDetailResponse.Code, purchaseReturnDetailResponse)
}

func (dc *PurchaseReturnDetailControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	purchaseReturnDetailResponse, _ := dc.PurchaseReturnDetailService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseReturnDetailResponse.Code, purchaseReturnDetailResponse)
}

func (dc *PurchaseReturnDetailControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseReturnDetailResponse, _ := dc.PurchaseReturnDetailService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(purchaseReturnDetailResponse.Code, purchaseReturnDetailResponse)
}