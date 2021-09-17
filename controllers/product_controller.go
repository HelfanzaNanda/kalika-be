package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type ProductController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type ProductControllerImpl struct {
		ProductService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (dc *ProductControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	productResponse, _ := dc.ProductService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductControllerImpl) FindAll(ctx echo.Context) error {
	productResponse, _ := dc.ProductService.FindAll(ctx)

	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductControllerImpl) Create(ctx echo.Context) error {
	productResponse, _ := dc.ProductService.Create(ctx)

	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	productResponse, _ := dc.ProductService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(productResponse.Code, productResponse)
}

func (dc *ProductControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	productResponse, _ := dc.ProductService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(productResponse.Code, productResponse)
}