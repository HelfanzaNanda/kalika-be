package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type CustomerController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type CustomerControllerImpl struct {

}

func NewCustomerController() CustomerController {
	return &CustomerControllerImpl{}
}

func (controller *CustomerControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *CustomerControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *CustomerControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *CustomerControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *CustomerControllerImpl) Delete(ctx echo.Context) error {
	return nil
}