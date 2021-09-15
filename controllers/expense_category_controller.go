package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type ExpenseCategoryController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type ExpenseCategoryControllerImpl struct {

}

func NewExpenseCategoryController() ExpenseCategoryController {
	return &ExpenseCategoryControllerImpl{}
}

func (controller *ExpenseCategoryControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *ExpenseCategoryControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *ExpenseCategoryControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *ExpenseCategoryControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *ExpenseCategoryControllerImpl) Delete(ctx echo.Context) error {
	return nil
}