package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type SalesReturnController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type SalesReturnControllerImpl struct {

}

func NewSalesReturnController() SalesReturnController {
	return &SalesReturnControllerImpl{}
}

func (controller *SalesReturnControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *SalesReturnControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *SalesReturnControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *SalesReturnControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *SalesReturnControllerImpl) Delete(ctx echo.Context) error {
	return nil
}