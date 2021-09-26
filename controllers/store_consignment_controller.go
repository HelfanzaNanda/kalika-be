package controllers

import (
	"kalika-be/helpers"
	"kalika-be/services"
	"github.com/labstack/echo"

)

type StoreConsignmentController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type StoreConsignmentControllerImpl struct {
		StoreConsignmentService services.StoreConsignmentService
}

func NewStoreConsignmentController(storeConsignmentService services.StoreConsignmentService) StoreConsignmentController {
	return &StoreConsignmentControllerImpl{
		StoreConsignmentService: storeConsignmentService,
	}
}

func (dc *StoreConsignmentControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	storeConsignmentResponse, _ := dc.StoreConsignmentService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(storeConsignmentResponse.Code, storeConsignmentResponse)
}

func (dc *StoreConsignmentControllerImpl) FindAll(ctx echo.Context) error {
	storeConsignmentResponse, _ := dc.StoreConsignmentService.FindAll(ctx)

	return ctx.JSON(storeConsignmentResponse.Code, storeConsignmentResponse)
}

func (dc *StoreConsignmentControllerImpl) Create(ctx echo.Context) error {

	storeConsignmentResponse, _ := dc.StoreConsignmentService.Create(ctx)

	return ctx.JSON(storeConsignmentResponse.Code, storeConsignmentResponse)
}

func (dc *StoreConsignmentControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	storeConsignmentResponse, _ := dc.StoreConsignmentService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(storeConsignmentResponse.Code, storeConsignmentResponse)
}

func (dc *StoreConsignmentControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	storeConsignmentResponse, _ := dc.StoreConsignmentService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(storeConsignmentResponse.Code, storeConsignmentResponse)
}

func (dc *StoreConsignmentControllerImpl) Datatable(ctx echo.Context) error {
	storeConsignmentResponse, _ := dc.StoreConsignmentService.Datatable(ctx)
	return ctx.JSON(202, storeConsignmentResponse)
}