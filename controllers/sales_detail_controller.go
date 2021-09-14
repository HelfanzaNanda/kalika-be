package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type SalesDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type SalesDetailControllerImpl struct {

}

func NewSalesDetailController() SalesDetailController {
	return &SalesDetailControllerImpl{}
}

func (controller *SalesDetailControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *SalesDetailControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *SalesDetailControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *SalesDetailControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *SalesDetailControllerImpl) Delete(ctx echo.Context) error {
	return nil
}