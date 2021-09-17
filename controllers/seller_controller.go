package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type SellerController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type SellerControllerImpl struct {
		SellerService services.SellerService
}

func NewSellerController(sellerService services.SellerService) SellerController {
	return &SellerControllerImpl{
		SellerService: sellerService,
	}
}

func (dc *SellerControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	sellerResponse, _ := dc.SellerService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(sellerResponse.Code, sellerResponse)
}

func (dc *SellerControllerImpl) FindAll(ctx echo.Context) error {
	sellerResponse, _ := dc.SellerService.FindAll(ctx)

	return ctx.JSON(sellerResponse.Code, sellerResponse)
}

func (dc *SellerControllerImpl) Create(ctx echo.Context) error {
	sellerResponse, _ := dc.SellerService.Create(ctx)

	return ctx.JSON(sellerResponse.Code, sellerResponse)
}

func (dc *SellerControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	sellerResponse, _ := dc.SellerService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(sellerResponse.Code, sellerResponse)
}

func (dc *SellerControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	sellerResponse, _ := dc.SellerService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(sellerResponse.Code, sellerResponse)
}