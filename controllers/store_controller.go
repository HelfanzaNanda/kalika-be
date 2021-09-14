package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type StoreController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type StoreControllerImpl struct {

}

func NewStoreController() StoreController {
	return &StoreControllerImpl{}
}

func (controller *StoreControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *StoreControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *StoreControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *StoreControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *StoreControllerImpl) Delete(ctx echo.Context) error {
	return nil
}