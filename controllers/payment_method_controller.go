package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type PaymentMethodController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type PaymentMethodControllerImpl struct {

}

func NewPaymentMethodController() PaymentMethodController {
	return &PaymentMethodControllerImpl{}
}

func (controller *PaymentMethodControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *PaymentMethodControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *PaymentMethodControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *PaymentMethodControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *PaymentMethodControllerImpl) Delete(ctx echo.Context) error {
	return nil
}