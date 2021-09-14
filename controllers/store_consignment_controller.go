package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type StoreConsignmentController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type StoreConsignmentControllerImpl struct {

}

func NewStoreConsignmentController() StoreConsignmentController {
	return &StoreConsignmentControllerImpl{}
}

func (controller *StoreConsignmentControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *StoreConsignmentControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *StoreConsignmentControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *StoreConsignmentControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *StoreConsignmentControllerImpl) Delete(ctx echo.Context) error {
	return nil
}