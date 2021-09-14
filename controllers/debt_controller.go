package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type DebtController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type DebtControllerImpl struct {

}

func NewDebtController() DebtController {
	return &DebtControllerImpl{}
}

func (controller *DebtControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *DebtControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *DebtControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *DebtControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *DebtControllerImpl) Delete(ctx echo.Context) error {
	return nil
}