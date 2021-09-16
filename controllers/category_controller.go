package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type CategoryController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type CategoryControllerImpl struct {
		CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (dc *CategoryControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	categoryResponse, _ := dc.CategoryService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(categoryResponse.Code, categoryResponse)
}

func (dc *CategoryControllerImpl) FindAll(ctx echo.Context) error {
	categoryResponse, _ := dc.CategoryService.FindAll(ctx)

	return ctx.JSON(categoryResponse.Code, categoryResponse)
}

func (dc *CategoryControllerImpl) Create(ctx echo.Context) error {
	categoryResponse, _ := dc.CategoryService.Create(ctx)

	return ctx.JSON(categoryResponse.Code, categoryResponse)
}

func (dc *CategoryControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	categoryResponse, _ := dc.CategoryService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(categoryResponse.Code, categoryResponse)
}

func (dc *CategoryControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	categoryResponse, _ := dc.CategoryService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(categoryResponse.Code, categoryResponse)
}