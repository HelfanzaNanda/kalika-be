package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type SalesConsignmentDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type SalesConsignmentDetailControllerImpl struct {

}

func NewSalesConsignmentDetailController() SalesConsignmentDetailController {
	return &SalesConsignmentDetailControllerImpl{}
}

func (controller *SalesConsignmentDetailControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *SalesConsignmentDetailControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *SalesConsignmentDetailControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *SalesConsignmentDetailControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *SalesConsignmentDetailControllerImpl) Delete(ctx echo.Context) error {
	return nil
}