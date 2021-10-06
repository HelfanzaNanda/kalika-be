package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type SalesConsignmentController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type SalesConsignmentControllerImpl struct {
		SalesConsignmentService services.SalesConsignmentService
}

func NewSalesConsignmentController(salesConsignmentService services.SalesConsignmentService) SalesConsignmentController {
	return &SalesConsignmentControllerImpl{
		SalesConsignmentService: salesConsignmentService,
	}
}

func (dc *SalesConsignmentControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	salesConsignmentResponse, _ := dc.SalesConsignmentService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesConsignmentResponse.Code, salesConsignmentResponse)
}

func (dc *SalesConsignmentControllerImpl) FindAll(ctx echo.Context) error {
	salesConsignmentResponse, _ := dc.SalesConsignmentService.FindAll(ctx)

	return ctx.JSON(salesConsignmentResponse.Code, salesConsignmentResponse)
}

func (dc *SalesConsignmentControllerImpl) Create(ctx echo.Context) error {
	salesConsignmentResponse, _ := dc.SalesConsignmentService.Create(ctx)

	return ctx.JSON(salesConsignmentResponse.Code, salesConsignmentResponse)
}

func (dc *SalesConsignmentControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	salesConsignmentResponse, _ := dc.SalesConsignmentService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesConsignmentResponse.Code, salesConsignmentResponse)
}

func (dc *SalesConsignmentControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	salesConsignmentResponse, _ := dc.SalesConsignmentService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(salesConsignmentResponse.Code, salesConsignmentResponse)
}

func (dc *SalesConsignmentControllerImpl) Datatable(ctx echo.Context) error {
	divisionResponse, _ := dc.SalesConsignmentService.Datatable(ctx)

	return ctx.JSON(202, divisionResponse)
}