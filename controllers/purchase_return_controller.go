package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type PurchaseReturnController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type PurchaseReturnControllerImpl struct {
		PurchaseReturnService services.PurchaseReturnService
}

func NewPurchaseReturnController(purchaseReturnService services.PurchaseReturnService) PurchaseReturnController {
	return &PurchaseReturnControllerImpl{
		PurchaseReturnService: purchaseReturnService,
	}
}

func (dc *PurchaseReturnControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseReturnResponse, _ := dc.PurchaseReturnService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseReturnResponse.Code, purchaseReturnResponse)
}

func (dc *PurchaseReturnControllerImpl) FindAll(ctx echo.Context) error {
	purchaseReturnResponse, _ := dc.PurchaseReturnService.FindAll(ctx)

	return ctx.JSON(purchaseReturnResponse.Code, purchaseReturnResponse)
}

func (dc *PurchaseReturnControllerImpl) Create(ctx echo.Context) error {
	purchaseReturnResponse, _ := dc.PurchaseReturnService.Create(ctx)

	return ctx.JSON(purchaseReturnResponse.Code, purchaseReturnResponse)
}

func (dc *PurchaseReturnControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	purchaseReturnResponse, _ := dc.PurchaseReturnService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseReturnResponse.Code, purchaseReturnResponse)
}

func (dc *PurchaseReturnControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseReturnResponse, _ := dc.PurchaseReturnService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(purchaseReturnResponse.Code, purchaseReturnResponse)
}