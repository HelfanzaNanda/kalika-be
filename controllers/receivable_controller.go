package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type ReceivableController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type ReceivableControllerImpl struct {

}

func NewReceivableController() ReceivableController {
	return &ReceivableControllerImpl{}
}

func (controller *ReceivableControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *ReceivableControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *ReceivableControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *ReceivableControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *ReceivableControllerImpl) Delete(ctx echo.Context) error {
	return nil
}