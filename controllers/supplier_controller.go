package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type SupplierController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type SupplierControllerImpl struct {

}

func NewSupplierController() SupplierController {
	return &SupplierControllerImpl{}
}

func (controller *SupplierControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *SupplierControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *SupplierControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *SupplierControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *SupplierControllerImpl) Delete(ctx echo.Context) error {
	return nil
}