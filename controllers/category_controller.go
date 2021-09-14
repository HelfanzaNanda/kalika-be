package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type CategoryController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type CategoryControllerImpl struct {

}

func NewCategoryController() CategoryController {
	return &CategoryControllerImpl{}
}

func (controller *CategoryControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *CategoryControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *CategoryControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *CategoryControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *CategoryControllerImpl) Delete(ctx echo.Context) error {
	return nil
}