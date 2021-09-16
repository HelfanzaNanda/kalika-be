package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type CakeTypeController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type CakeTypeControllerImpl struct {
		CakeTypeService services.CakeTypeService
}

func NewCakeTypeController(cakeTypeService services.CakeTypeService) CakeTypeController {
	return &CakeTypeControllerImpl{
		CakeTypeService: cakeTypeService,
	}
}

func (dc *CakeTypeControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	cakeTypeResponse, _ := dc.CakeTypeService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(cakeTypeResponse.Code, cakeTypeResponse)
}

func (dc *CakeTypeControllerImpl) FindAll(ctx echo.Context) error {
	cakeTypeResponse, _ := dc.CakeTypeService.FindAll(ctx)

	return ctx.JSON(cakeTypeResponse.Code, cakeTypeResponse)
}

func (dc *CakeTypeControllerImpl) Create(ctx echo.Context) error {
	cakeTypeResponse, _ := dc.CakeTypeService.Create(ctx)

	return ctx.JSON(cakeTypeResponse.Code, cakeTypeResponse)
}

func (dc *CakeTypeControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	cakeTypeResponse, _ := dc.CakeTypeService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(cakeTypeResponse.Code, cakeTypeResponse)
}

func (dc *CakeTypeControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	cakeTypeResponse, _ := dc.CakeTypeService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(cakeTypeResponse.Code, cakeTypeResponse)
}