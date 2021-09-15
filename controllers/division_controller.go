package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type DivisionController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type DivisionControllerImpl struct {
	DivisionService services.DivisionService
}

func NewDivisionController(divisionService services.DivisionService) DivisionController {
	return &DivisionControllerImpl{
		DivisionService: divisionService,
	}
}

func (dc *DivisionControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	userResponse, _ := dc.DivisionService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(userResponse.Code, userResponse)
}

func (dc *DivisionControllerImpl) FindAll(ctx echo.Context) error {
	userResponse, _ := dc.DivisionService.FindAll(ctx)

	return ctx.JSON(userResponse.Code, userResponse)
}

func (dc *DivisionControllerImpl) Create(ctx echo.Context) error {
	userResponse, _ := dc.DivisionService.Create(ctx)

	return ctx.JSON(userResponse.Code, userResponse)
}

func (dc *DivisionControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	userResponse, _ := dc.DivisionService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(userResponse.Code, userResponse)
}

func (dc *DivisionControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	userResponse, _ := dc.DivisionService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(userResponse.Code, userResponse)
}