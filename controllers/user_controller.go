package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type UserController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type UserControllerImpl struct {

}

func NewUserController() UserController {
	return &UserControllerImpl{}
}

func (controller *UserControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *UserControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *UserControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *UserControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *UserControllerImpl) Delete(ctx echo.Context) error {
	return nil
}