package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type CakeVariantController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type CakeVariantControllerImpl struct {

}

func NewCakeVariantController() CakeVariantController {
	return &CakeVariantControllerImpl{}
}

func (controller *CakeVariantControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *CakeVariantControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *CakeVariantControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *CakeVariantControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *CakeVariantControllerImpl) Delete(ctx echo.Context) error {
	return nil
}