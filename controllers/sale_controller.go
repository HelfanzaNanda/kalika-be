package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type SaleController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type SaleControllerImpl struct {

}

func NewSaleController() SaleController {
	return &SaleControllerImpl{}
}

func (controller *SaleControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *SaleControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *SaleControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *SaleControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *SaleControllerImpl) Delete(ctx echo.Context) error {
	return nil
}