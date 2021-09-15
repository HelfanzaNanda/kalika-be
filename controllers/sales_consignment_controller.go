package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type SalesConsignmentController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type SalesConsignmentControllerImpl struct {

}

func NewSalesConsignmentController() SalesConsignmentController {
	return &SalesConsignmentControllerImpl{}
}

func (controller *SalesConsignmentControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *SalesConsignmentControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *SalesConsignmentControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *SalesConsignmentControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *SalesConsignmentControllerImpl) Delete(ctx echo.Context) error {
	return nil
}