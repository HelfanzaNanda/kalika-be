package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type CakeVariantController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type CakeVariantControllerImpl struct {
		CakeVariantService services.CakeVariantService
}

func NewCakeVariantController(cakeVariantService services.CakeVariantService) CakeVariantController {
	return &CakeVariantControllerImpl{
		CakeVariantService: cakeVariantService,
	}
}

func (dc *CakeVariantControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	cakeVariantResponse, _ := dc.CakeVariantService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(cakeVariantResponse.Code, cakeVariantResponse)
}

func (dc *CakeVariantControllerImpl) FindAll(ctx echo.Context) error {
	cakeVariantResponse, _ := dc.CakeVariantService.FindAll(ctx)

	return ctx.JSON(cakeVariantResponse.Code, cakeVariantResponse)
}

func (dc *CakeVariantControllerImpl) Create(ctx echo.Context) error {
	cakeVariantResponse, _ := dc.CakeVariantService.Create(ctx)

	return ctx.JSON(cakeVariantResponse.Code, cakeVariantResponse)
}

func (dc *CakeVariantControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	cakeVariantResponse, _ := dc.CakeVariantService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(cakeVariantResponse.Code, cakeVariantResponse)
}

func (dc *CakeVariantControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	cakeVariantResponse, _ := dc.CakeVariantService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(cakeVariantResponse.Code, cakeVariantResponse)
}

func (dc *CakeVariantControllerImpl) Datatable(ctx echo.Context) error {
	cakeVariantResponse, _ := dc.CakeVariantService.Datatable(ctx)
	return ctx.JSON(202, cakeVariantResponse)
}