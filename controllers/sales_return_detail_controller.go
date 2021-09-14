package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type SalesReturnDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type SalesReturnDetailControllerImpl struct {

}

func NewSalesReturnDetailController() SalesReturnDetailController {
	return &SalesReturnDetailControllerImpl{}
}

func (controller *SalesReturnDetailControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *SalesReturnDetailControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *SalesReturnDetailControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *SalesReturnDetailControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *SalesReturnDetailControllerImpl) Delete(ctx echo.Context) error {
	return nil
}