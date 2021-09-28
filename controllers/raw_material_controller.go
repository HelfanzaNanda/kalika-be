package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type RawMaterialController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type RawMaterialControllerImpl struct {
		RawMaterialService services.RawMaterialService
}

func NewRawMaterialController(rawMaterialService services.RawMaterialService) RawMaterialController {
	return &RawMaterialControllerImpl{
		RawMaterialService: rawMaterialService,
	}
}

func (dc *RawMaterialControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	rawMaterialResponse, _ := dc.RawMaterialService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(rawMaterialResponse.Code, rawMaterialResponse)
}

func (dc *RawMaterialControllerImpl) FindAll(ctx echo.Context) error {
	rawMaterialResponse, _ := dc.RawMaterialService.FindAll(ctx)

	return ctx.JSON(rawMaterialResponse.Code, rawMaterialResponse)
}

func (dc *RawMaterialControllerImpl) Create(ctx echo.Context) error {
	rawMaterialResponse, _ := dc.RawMaterialService.Create(ctx)

	return ctx.JSON(rawMaterialResponse.Code, rawMaterialResponse)
}

func (dc *RawMaterialControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	rawMaterialResponse, _ := dc.RawMaterialService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(rawMaterialResponse.Code, rawMaterialResponse)
}

func (dc *RawMaterialControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	rawMaterialResponse, _ := dc.RawMaterialService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(rawMaterialResponse.Code, rawMaterialResponse)
}

func (dc *RawMaterialControllerImpl) Datatable(ctx echo.Context) error {
	rawMaterialResponse, _ := dc.RawMaterialService.Datatable(ctx)
	return ctx.JSON(202, rawMaterialResponse)
}