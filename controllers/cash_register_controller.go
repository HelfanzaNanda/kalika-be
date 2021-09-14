package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type CashRegisterController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type CashRegisterControllerImpl struct {

}

func NewCashRegisterController() CashRegisterController {
	return &CashRegisterControllerImpl{}
}

func (controller *CashRegisterControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *CashRegisterControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *CashRegisterControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *CashRegisterControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *CashRegisterControllerImpl) Delete(ctx echo.Context) error {
	return nil
}