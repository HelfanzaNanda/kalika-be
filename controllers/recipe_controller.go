package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type RecipeController interface {
	FindById(ctx echo.Context) error
	FindByProductId(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type RecipeControllerImpl struct {
		RecipeService services.RecipeService
}

func NewRecipeController(roleService services.RecipeService) RecipeController {
	return &RecipeControllerImpl{
		RecipeService: roleService,
	}
}

func (dc *RecipeControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	roleResponse, _ := dc.RecipeService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(roleResponse.Code, roleResponse)
}

func (dc *RecipeControllerImpl) FindByProductId(ctx echo.Context) error {
	id := ctx.Param("id")
	roleResponse, _ := dc.RecipeService.FindByProductId(ctx, helpers.StringToInt(id))

	return ctx.JSON(roleResponse.Code, roleResponse)
}

func (dc *RecipeControllerImpl) FindAll(ctx echo.Context) error {
	roleResponse, _ := dc.RecipeService.FindAll(ctx)

	return ctx.JSON(roleResponse.Code, roleResponse)
}

func (dc *RecipeControllerImpl) Create(ctx echo.Context) error {
	roleResponse, _ := dc.RecipeService.Create(ctx)

	return ctx.JSON(roleResponse.Code, roleResponse)
}

func (dc *RecipeControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	roleResponse, _ := dc.RecipeService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(roleResponse.Code, roleResponse)
}

func (dc *RecipeControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	roleResponse, _ := dc.RecipeService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(roleResponse.Code, roleResponse)
}

func (dc *RecipeControllerImpl) Datatable(ctx echo.Context) error {
	roleResponse, _ := dc.RecipeService.Datatable(ctx)

	return ctx.JSON(202, roleResponse)
}