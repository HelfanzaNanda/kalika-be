package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type DivisionController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type DivisionControllerImpl struct {

}

func NewDivisionController() DivisionController {
	return &DivisionControllerImpl{}
}

func (controller *DivisionControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *DivisionControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *DivisionControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *DivisionControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *DivisionControllerImpl) Delete(ctx echo.Context) error {
	return nil
}