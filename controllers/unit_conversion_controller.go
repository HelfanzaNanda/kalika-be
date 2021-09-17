package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type UnitConversionController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type UnitConversionControllerImpl struct {
		UnitConversionService services.UnitConversionService
}

func NewUnitConversionController(unitConversionService services.UnitConversionService) UnitConversionController {
	return &UnitConversionControllerImpl{
		UnitConversionService: unitConversionService,
	}
}

func (dc *UnitConversionControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	unitConversionResponse, _ := dc.UnitConversionService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(unitConversionResponse.Code, unitConversionResponse)
}

func (dc *UnitConversionControllerImpl) FindAll(ctx echo.Context) error {
	unitConversionResponse, _ := dc.UnitConversionService.FindAll(ctx)

	return ctx.JSON(unitConversionResponse.Code, unitConversionResponse)
}

func (dc *UnitConversionControllerImpl) Create(ctx echo.Context) error {
	unitConversionResponse, _ := dc.UnitConversionService.Create(ctx)

	return ctx.JSON(unitConversionResponse.Code, unitConversionResponse)
}

func (dc *UnitConversionControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	unitConversionResponse, _ := dc.UnitConversionService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(unitConversionResponse.Code, unitConversionResponse)
}

func (dc *UnitConversionControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	unitConversionResponse, _ := dc.UnitConversionService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(unitConversionResponse.Code, unitConversionResponse)
}