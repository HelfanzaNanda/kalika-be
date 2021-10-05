package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type PurchaseOrderDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	//Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type PurchaseOrderDetailControllerImpl struct {
		PurchaseOrderDetailService services.PurchaseOrderDetailService
}

func NewPurchaseOrderDetailController(purchaseOrderDetailService services.PurchaseOrderDetailService) PurchaseOrderDetailController {
	return &PurchaseOrderDetailControllerImpl{
		PurchaseOrderDetailService: purchaseOrderDetailService,
	}
}

func (dc *PurchaseOrderDetailControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseOrderDetailResponse, _ := dc.PurchaseOrderDetailService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseOrderDetailResponse.Code, purchaseOrderDetailResponse)
}

func (dc *PurchaseOrderDetailControllerImpl) FindAll(ctx echo.Context) error {
	purchaseOrderDetailResponse, _ := dc.PurchaseOrderDetailService.FindAll(ctx)

	return ctx.JSON(purchaseOrderDetailResponse.Code, purchaseOrderDetailResponse)
}

//func (dc *PurchaseOrderDetailControllerImpl) Create(ctx echo.Context) error {
//	purchaseOrderDetailResponse, _ := dc.PurchaseOrderDetailService.Create(ctx)
//
//	return ctx.JSON(purchaseOrderDetailResponse.Code, purchaseOrderDetailResponse)
//}

func (dc *PurchaseOrderDetailControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	purchaseOrderDetailResponse, _ := dc.PurchaseOrderDetailService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(purchaseOrderDetailResponse.Code, purchaseOrderDetailResponse)
}

func (dc *PurchaseOrderDetailControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	purchaseOrderDetailResponse, _ := dc.PurchaseOrderDetailService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(purchaseOrderDetailResponse.Code, purchaseOrderDetailResponse)
}