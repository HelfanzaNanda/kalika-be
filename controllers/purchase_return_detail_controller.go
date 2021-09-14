package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type PurchaseReturnDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type PurchaseReturnDetailControllerImpl struct {

}

func NewPurchaseReturnDetailController() PurchaseReturnDetailController {
	return &PurchaseReturnDetailControllerImpl{}
}

func (controller *PurchaseReturnDetailControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseReturnDetailControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *PurchaseReturnDetailControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseReturnDetailControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseReturnDetailControllerImpl) Delete(ctx echo.Context) error {
	return nil
}