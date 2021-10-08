package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type StockOpnameController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type StockOpnameControllerImpl struct {
		StockOpnameService services.StockOpnameService
}

func NewStockOpnameController(stockOpnameService services.StockOpnameService) StockOpnameController {
	return &StockOpnameControllerImpl{
		StockOpnameService: stockOpnameService,
	}
}

func (o *StockOpnameControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	stockOpnameResponse, _ := o.StockOpnameService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *StockOpnameControllerImpl) FindAll(ctx echo.Context) error {
	stockOpnameResponse, _ := o.StockOpnameService.FindAll(ctx)

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *StockOpnameControllerImpl) Create(ctx echo.Context) error {
	stockOpnameResponse, _ := o.StockOpnameService.Create(ctx)

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *StockOpnameControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	stockOpnameResponse, _ := o.StockOpnameService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *StockOpnameControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	stockOpnameResponse, _ := o.StockOpnameService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *StockOpnameControllerImpl) Datatable(ctx echo.Context) error {
	divisionResponse, _ := o.StockOpnameService.Datatable(ctx)

	return ctx.JSON(202, divisionResponse)
}