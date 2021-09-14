package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type UnitConversionController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type UnitConversionControllerImpl struct {

}

func NewUnitConversionController() UnitConversionController {
	return &UnitConversionControllerImpl{}
}

func (controller *UnitConversionControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *UnitConversionControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *UnitConversionControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *UnitConversionControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *UnitConversionControllerImpl) Delete(ctx echo.Context) error {
	return nil
}