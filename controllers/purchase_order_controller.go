package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type PurchaseOrderController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type PurchaseOrderControllerImpl struct {

}

func NewPurchaseOrderController() PurchaseOrderController {
	return &PurchaseOrderControllerImpl{}
}

func (controller *PurchaseOrderControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *PurchaseOrderControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *PurchaseOrderControllerImpl) Delete(ctx echo.Context) error {
	return nil
}