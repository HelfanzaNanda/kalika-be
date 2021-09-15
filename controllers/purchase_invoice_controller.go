package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type PurchaseInvoiceController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type PurchaseInvoiceControllerImpl struct {

}

func NewPurchaseInvoiceController() PurchaseInvoiceController {
	return &PurchaseInvoiceControllerImpl{}
}

func (controller *PurchaseInvoiceControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseInvoiceControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *PurchaseInvoiceControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseInvoiceControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseInvoiceControllerImpl) Delete(ctx echo.Context) error {
	return nil
}