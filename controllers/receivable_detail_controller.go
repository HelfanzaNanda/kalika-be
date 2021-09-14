package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type ReceivableDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type ReceivableDetailControllerImpl struct {

}

func NewReceivableDetailController() ReceivableDetailController {
	return &ReceivableDetailControllerImpl{}
}

func (controller *ReceivableDetailControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *ReceivableDetailControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *ReceivableDetailControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *ReceivableDetailControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *ReceivableDetailControllerImpl) Delete(ctx echo.Context) error {
	return nil
}