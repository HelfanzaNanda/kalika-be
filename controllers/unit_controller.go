package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type UnitController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type UnitControllerImpl struct {

}

func NewUnitController() UnitController {
	return &UnitControllerImpl{}
}

func (controller *UnitControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *UnitControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *UnitControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *UnitControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *UnitControllerImpl) Delete(ctx echo.Context) error {
	return nil
}