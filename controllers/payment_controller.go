package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type PaymentController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type PaymentControllerImpl struct {

}

func NewPaymentController() PaymentController {
	return &PaymentControllerImpl{}
}

func (controller *PaymentControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *PaymentControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *PaymentControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *PaymentControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *PaymentControllerImpl) Delete(ctx echo.Context) error {
	return nil
}