package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type ExpenseController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type ExpenseControllerImpl struct {

}

func NewExpenseController() ExpenseController {
	return &ExpenseControllerImpl{}
}

func (controller *ExpenseControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *ExpenseControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *ExpenseControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *ExpenseControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *ExpenseControllerImpl) Delete(ctx echo.Context) error {
	return nil
}