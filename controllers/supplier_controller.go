package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type SupplierController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type SupplierControllerImpl struct {
		SupplierService services.SupplierService
}

func NewSupplierController(supplierService services.SupplierService) SupplierController {
	return &SupplierControllerImpl{
		SupplierService: supplierService,
	}
}

func (dc *SupplierControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	supplierResponse, _ := dc.SupplierService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(supplierResponse.Code, supplierResponse)
}

func (dc *SupplierControllerImpl) FindAll(ctx echo.Context) error {
	supplierResponse, _ := dc.SupplierService.FindAll(ctx)

	return ctx.JSON(supplierResponse.Code, supplierResponse)
}

func (dc *SupplierControllerImpl) Create(ctx echo.Context) error {
	supplierResponse, _ := dc.SupplierService.Create(ctx)

	return ctx.JSON(supplierResponse.Code, supplierResponse)
}

func (dc *SupplierControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	supplierResponse, _ := dc.SupplierService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(supplierResponse.Code, supplierResponse)
}

func (dc *SupplierControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	supplierResponse, _ := dc.SupplierService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(supplierResponse.Code, supplierResponse)
}