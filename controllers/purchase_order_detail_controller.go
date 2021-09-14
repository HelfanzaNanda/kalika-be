package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type PurchaseOrderDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type PurchaseOrderDetailControllerImpl struct {

}

func NewPurchaseOrderDetailController() PurchaseOrderDetailController {
	return &PurchaseOrderDetailControllerImpl{}
}

func (controller *PurchaseOrderDetailControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderDetailControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *PurchaseOrderDetailControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderDetailControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderDetailControllerImpl) Delete(ctx echo.Context) error {
	return nil
}