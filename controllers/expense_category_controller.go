package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type ExpenseCategoryController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type ExpenseCategoryControllerImpl struct {
		ExpenseCategoryService services.ExpenseCategoryService
}

func NewExpenseCategoryController(expenseCategoryService services.ExpenseCategoryService) ExpenseCategoryController {
	return &ExpenseCategoryControllerImpl{
		ExpenseCategoryService: expenseCategoryService,
	}
}

func (dc *ExpenseCategoryControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	expenseCategoryResponse, _ := dc.ExpenseCategoryService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(expenseCategoryResponse.Code, expenseCategoryResponse)
}

func (dc *ExpenseCategoryControllerImpl) FindAll(ctx echo.Context) error {
	expenseCategoryResponse, _ := dc.ExpenseCategoryService.FindAll(ctx)

	return ctx.JSON(expenseCategoryResponse.Code, expenseCategoryResponse)
}

func (dc *ExpenseCategoryControllerImpl) Create(ctx echo.Context) error {
	expenseCategoryResponse, _ := dc.ExpenseCategoryService.Create(ctx)

	return ctx.JSON(expenseCategoryResponse.Code, expenseCategoryResponse)
}

func (dc *ExpenseCategoryControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	expenseCategoryResponse, _ := dc.ExpenseCategoryService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(expenseCategoryResponse.Code, expenseCategoryResponse)
}

func (dc *ExpenseCategoryControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	expenseCategoryResponse, _ := dc.ExpenseCategoryService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(expenseCategoryResponse.Code, expenseCategoryResponse)
}