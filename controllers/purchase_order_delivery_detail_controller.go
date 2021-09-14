package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type PurchaseOrderDeliveryDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type PurchaseOrderDeliveryDetailControllerImpl struct {

}

func NewPurchaseOrderDeliveryDetailController() PurchaseOrderDeliveryDetailController {
	return &PurchaseOrderDeliveryDetailControllerImpl{}
}

func (controller *PurchaseOrderDeliveryDetailControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderDeliveryDetailControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *PurchaseOrderDeliveryDetailControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderDeliveryDetailControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderDeliveryDetailControllerImpl) Delete(ctx echo.Context) error {
	return nil
}