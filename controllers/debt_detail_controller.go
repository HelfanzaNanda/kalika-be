package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type DebtDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type DebtDetailControllerImpl struct {

}

func NewDebtDetailController() DebtDetailController {
	return &DebtDetailControllerImpl{}
}

func (controller *DebtDetailControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *DebtDetailControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *DebtDetailControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *DebtDetailControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *DebtDetailControllerImpl) Delete(ctx echo.Context) error {
	return nil
}