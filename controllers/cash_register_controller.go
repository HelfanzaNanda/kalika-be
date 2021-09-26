package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type CashRegisterController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type CashRegisterControllerImpl struct {
		CashRegisterService services.CashRegisterService
}

func NewCashRegisterController(cashRegisterService services.CashRegisterService) CashRegisterController {
	return &CashRegisterControllerImpl{
		CashRegisterService: cashRegisterService,
	}
}

func (dc *CashRegisterControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	cashRegisterResponse, _ := dc.CashRegisterService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(cashRegisterResponse.Code, cashRegisterResponse)
}

func (dc *CashRegisterControllerImpl) FindAll(ctx echo.Context) error {
	cashRegisterResponse, _ := dc.CashRegisterService.FindAll(ctx)

	return ctx.JSON(cashRegisterResponse.Code, cashRegisterResponse)
}

func (dc *CashRegisterControllerImpl) Create(ctx echo.Context) error {
	cashRegisterResponse, _ := dc.CashRegisterService.Create(ctx)

	return ctx.JSON(cashRegisterResponse.Code, cashRegisterResponse)
}

func (dc *CashRegisterControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	cashRegisterResponse, _ := dc.CashRegisterService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(cashRegisterResponse.Code, cashRegisterResponse)
}

func (dc *CashRegisterControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	cashRegisterResponse, _ := dc.CashRegisterService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(cashRegisterResponse.Code, cashRegisterResponse)
}

func (dc *CashRegisterControllerImpl) Datatable(ctx echo.Context) error {
	cashRegisterResponse, _ := dc.CashRegisterService.Datatable(ctx)
	return ctx.JSON(202, cashRegisterResponse)
}