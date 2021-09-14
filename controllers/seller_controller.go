package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type SellerController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type SellerControllerImpl struct {

}

func NewSellerController() SellerController {
	return &SellerControllerImpl{}
}

func (controller *SellerControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *SellerControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *SellerControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *SellerControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *SellerControllerImpl) Delete(ctx echo.Context) error {
	return nil
}