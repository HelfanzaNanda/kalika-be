package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type PurchaseOrderDeliveryController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type PurchaseOrderDeliveryControllerImpl struct {

}

func NewPurchaseOrderDeliveryController() PurchaseOrderDeliveryController {
	return &PurchaseOrderDeliveryControllerImpl{}
}

func (controller *PurchaseOrderDeliveryControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderDeliveryControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *PurchaseOrderDeliveryControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderDeliveryControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderDeliveryControllerImpl) Delete(ctx echo.Context) error {
	return nil
}