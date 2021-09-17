package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type ExpenseDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type ExpenseDetailControllerImpl struct {
		ExpenseDetailService services.ExpenseDetailService
}

func NewExpenseDetailController(expenseDetailService services.ExpenseDetailService) ExpenseDetailController {
	return &ExpenseDetailControllerImpl{
		ExpenseDetailService: expenseDetailService,
	}
}

func (dc *ExpenseDetailControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	expenseDetailResponse, _ := dc.ExpenseDetailService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(expenseDetailResponse.Code, expenseDetailResponse)
}

func (dc *ExpenseDetailControllerImpl) FindAll(ctx echo.Context) error {
	expenseDetailResponse, _ := dc.ExpenseDetailService.FindAll(ctx)

	return ctx.JSON(expenseDetailResponse.Code, expenseDetailResponse)
}

func (dc *ExpenseDetailControllerImpl) Create(ctx echo.Context) error {
	expenseDetailResponse, _ := dc.ExpenseDetailService.Create(ctx)

	return ctx.JSON(expenseDetailResponse.Code, expenseDetailResponse)
}

func (dc *ExpenseDetailControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	expenseDetailResponse, _ := dc.ExpenseDetailService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(expenseDetailResponse.Code, expenseDetailResponse)
}

func (dc *ExpenseDetailControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	expenseDetailResponse, _ := dc.ExpenseDetailService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(expenseDetailResponse.Code, expenseDetailResponse)
}