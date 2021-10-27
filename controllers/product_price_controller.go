package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type ProductPriceController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type ProductPriceControllerImpl struct {
		ProductPriceService services.ProductPriceService
}

func NewProductPriceController(productService services.ProductPriceService) ProductPriceController {
	return &ProductPriceControllerImpl{
		ProductPriceService: productService,
	}
}

func (dc *ProductPriceControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	productResponse, _ := dc.ProductPriceService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductPriceControllerImpl) FindAll(ctx echo.Context) error {
	productResponse, _ := dc.ProductPriceService.FindAll(ctx)

	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductPriceControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	productResponse, _ := dc.ProductPriceService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductPriceControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	productResponse, _ := dc.ProductPriceService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(productResponse.Code, productResponse)
}