package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type RawMaterialController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type RawMaterialControllerImpl struct {

}

func NewRawMaterialController() RawMaterialController {
	return &RawMaterialControllerImpl{}
}

func (controller *RawMaterialControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *RawMaterialControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *RawMaterialControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *RawMaterialControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *RawMaterialControllerImpl) Delete(ctx echo.Context) error {
	return nil
}