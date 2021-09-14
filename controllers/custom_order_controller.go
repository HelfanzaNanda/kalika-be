package controllers

import (
	"fmt"
	"github.com/labstack/echo"
)

type CustomOrderController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type CustomOrderControllerImpl struct {

}

func NewCustomOrderController() CustomOrderController {
	return &CustomOrderControllerImpl{}
}

func (controller *CustomOrderControllerImpl) FindById(ctx echo.Context) error {
	return nil
}

func (controller *CustomOrderControllerImpl) FindAll(ctx echo.Context) error {
	fmt.Println("HALO GUYS")
	fmt.Println(ctx.QueryParams())
	return nil
}

func (controller *CustomOrderControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (controller *CustomOrderControllerImpl) Update(ctx echo.Context) error {
	return nil
}

func (controller *CustomOrderControllerImpl) Delete(ctx echo.Context) error {
	return nil
}