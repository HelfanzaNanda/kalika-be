package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type StoreController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type StoreControllerImpl struct {
		StoreService services.StoreService
}

func NewStoreController(storeService services.StoreService) StoreController {
	return &StoreControllerImpl{
		StoreService: storeService,
	}
}

func (dc *StoreControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	storeResponse, _ := dc.StoreService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(storeResponse.Code, storeResponse)
}

func (dc *StoreControllerImpl) FindAll(ctx echo.Context) error {
	storeResponse, _ := dc.StoreService.FindAll(ctx)

	return ctx.JSON(storeResponse.Code, storeResponse)
}

func (dc *StoreControllerImpl) Create(ctx echo.Context) error {
	storeResponse, _ := dc.StoreService.Create(ctx)

	return ctx.JSON(storeResponse.Code, storeResponse)
}

func (dc *StoreControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	storeResponse, _ := dc.StoreService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(storeResponse.Code, storeResponse)
}

func (dc *StoreControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	storeResponse, _ := dc.StoreService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(storeResponse.Code, storeResponse)
}

func (dc *StoreControllerImpl) Datatable(ctx echo.Context) error {
	storeResponse, _ := dc.StoreService.Datatable(ctx)
	return ctx.JSON(202, storeResponse)
}