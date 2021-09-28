package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type UnitController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type UnitControllerImpl struct {
		UnitService services.UnitService
}

func NewUnitController(unitService services.UnitService) UnitController {
	return &UnitControllerImpl{
		UnitService: unitService,
	}
}

func (dc *UnitControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	unitResponse, _ := dc.UnitService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(unitResponse.Code, unitResponse)
}

func (dc *UnitControllerImpl) FindAll(ctx echo.Context) error {
	unitResponse, _ := dc.UnitService.FindAll(ctx)

	return ctx.JSON(unitResponse.Code, unitResponse)
}

func (dc *UnitControllerImpl) Create(ctx echo.Context) error {
	unitResponse, _ := dc.UnitService.Create(ctx)

	return ctx.JSON(unitResponse.Code, unitResponse)
}

func (dc *UnitControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	unitResponse, _ := dc.UnitService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(unitResponse.Code, unitResponse)
}

func (dc *UnitControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	unitResponse, _ := dc.UnitService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(unitResponse.Code, unitResponse)
}

func (dc *UnitControllerImpl) Datatable(ctx echo.Context) error {
	unitResponse, _ := dc.UnitService.Datatable(ctx)
	return ctx.JSON(202, unitResponse)
}