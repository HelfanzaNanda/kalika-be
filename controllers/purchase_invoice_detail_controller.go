package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type PurchaseInvoiceDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type PurchaseInvoiceDetailControllerImpl struct {

}

func NewPurchaseInvoiceDetailController() PurchaseInvoiceDetailController {
	return &PurchaseInvoiceDetailControllerImpl{}
}

func (controller *PurchaseInvoiceDetailControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseInvoiceDetailControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *PurchaseInvoiceDetailControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseInvoiceDetailControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseInvoiceDetailControllerImpl) Delete(ctx echo.Context) error {
	return nil
}