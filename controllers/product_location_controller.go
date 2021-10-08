package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type ProductLocationController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type ProductLocationControllerImpl struct {
		ProductLocationService services.ProductLocationService
}

func NewProductLocationController(productService services.ProductLocationService) ProductLocationController {
	return &ProductLocationControllerImpl{
		ProductLocationService: productService,
	}
}

func (dc *ProductLocationControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	productResponse, _ := dc.ProductLocationService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductLocationControllerImpl) FindAll(ctx echo.Context) error {
	productResponse, _ := dc.ProductLocationService.FindAll(ctx)

	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductLocationControllerImpl) Create(ctx echo.Context) error {
	productResponse, _ := dc.ProductLocationService.Create(ctx)

	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductLocationControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	productResponse, _ := dc.ProductLocationService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductLocationControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	productResponse, _ := dc.ProductLocationService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductLocationControllerImpl) Datatable(ctx echo.Context) error {
	productResponse, _ := dc.ProductLocationService.Datatable(ctx)
	return ctx.JSON(202, productResponse)
}