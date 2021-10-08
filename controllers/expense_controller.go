package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type ExpenseController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
	ReportDatatable(ctx echo.Context) error
}
type ExpenseControllerImpl struct {
		ExpenseService services.ExpenseService
}

func NewExpenseController(expenseService services.ExpenseService) ExpenseController {
	return &ExpenseControllerImpl{
		ExpenseService: expenseService,
	}
}

func (dc *ExpenseControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	expenseResponse, _ := dc.ExpenseService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(expenseResponse.Code, expenseResponse)
}

func (dc *ExpenseControllerImpl) FindAll(ctx echo.Context) error {
	expenseResponse, _ := dc.ExpenseService.FindAll(ctx)

	return ctx.JSON(expenseResponse.Code, expenseResponse)
}

func (dc *ExpenseControllerImpl) Create(ctx echo.Context) error {
	expenseResponse, _ := dc.ExpenseService.Create(ctx)

	return ctx.JSON(expenseResponse.Code, expenseResponse)
}

func (dc *ExpenseControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	expenseResponse, _ := dc.ExpenseService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(expenseResponse.Code, expenseResponse)
}

func (dc *ExpenseControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	expenseResponse, _ := dc.ExpenseService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(expenseResponse.Code, expenseResponse)
}

func (dc *ExpenseControllerImpl) Datatable(ctx echo.Context) error {
	expenseResponse, _ := dc.ExpenseService.Datatable(ctx)
	return ctx.JSON(202, expenseResponse)
}

func (dc *ExpenseControllerImpl) ReportDatatable(ctx echo.Context) error {
	expenseResponse, _ := dc.ExpenseService.ReportDatatable(ctx)
	return ctx.JSON(202, expenseResponse)
}