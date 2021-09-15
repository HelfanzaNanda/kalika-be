package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type CakeTypeController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type CakeTypeControllerImpl struct {

}

func NewCakeTypeController() CakeTypeController {
	return &CakeTypeControllerImpl{}
}

func (controller *CakeTypeControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *CakeTypeControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *CakeTypeControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *CakeTypeControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *CakeTypeControllerImpl) Delete(ctx echo.Context) error {
	return nil
}