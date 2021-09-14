package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type PurchaseReturnController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type PurchaseReturnControllerImpl struct {

}

func NewPurchaseReturnController() PurchaseReturnController {
	return &PurchaseReturnControllerImpl{}
}

func (controller *PurchaseReturnControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseReturnControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *PurchaseReturnControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseReturnControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseReturnControllerImpl) Delete(ctx echo.Context) error {
	return nil
}