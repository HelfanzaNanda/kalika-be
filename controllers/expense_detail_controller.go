package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type ExpenseDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type ExpenseDetailControllerImpl struct {

}

func NewExpenseDetailController() ExpenseDetailController {
	return &ExpenseDetailControllerImpl{}
}

func (controller *ExpenseDetailControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *ExpenseDetailControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *ExpenseDetailControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *ExpenseDetailControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *ExpenseDetailControllerImpl) Delete(ctx echo.Context) error {
	return nil
}