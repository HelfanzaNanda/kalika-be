package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type StoreMutationController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
	GeneratePdf(ctx echo.Context) error
}
type StoreMutationControllerImpl struct {
	StoreMutationService services.StoreMutationService
}

func NewStoreMutationController(stockOpnameService services.StoreMutationService) StoreMutationController {
	return &StoreMutationControllerImpl{
		StoreMutationService: stockOpnameService,
	}
}

func (o *StoreMutationControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	stockOpnameResponse, _ := o.StoreMutationService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *StoreMutationControllerImpl) FindAll(ctx echo.Context) error {
	stockOpnameResponse, _ := o.StoreMutationService.FindAll(ctx)

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *StoreMutationControllerImpl) Create(ctx echo.Context) error {
	stockOpnameResponse, _ := o.StoreMutationService.Create(ctx)

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *StoreMutationControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	stockOpnameResponse, _ := o.StoreMutationService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *StoreMutationControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	stockOpnameResponse, _ := o.StoreMutationService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(stockOpnameResponse.Code, stockOpnameResponse)
}

func (o *StoreMutationControllerImpl) Datatable(ctx echo.Context) error {
	stockOpnameResponse, _ := o.StoreMutationService.Datatable(ctx)

	return ctx.JSON(202, stockOpnameResponse)
}

func (o *StoreMutationControllerImpl) GeneratePdf(ctx echo.Context) error {
	stockOpnameId := ctx.Param("id")
	stockOpnameResponse, _ := o.StoreMutationService.GeneratePdf(ctx, helpers.StringToInt(stockOpnameId))
	return ctx.JSON(202, stockOpnameResponse)
}