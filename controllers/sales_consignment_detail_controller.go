package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type SalesConsignmentDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type SalesConsignmentDetailControllerImpl struct {
		SalesConsignmentDetailService services.SalesConsignmentDetailService
}

func NewSalesConsignmentDetailController(salesConsignmentDetailService services.SalesConsignmentDetailService) SalesConsignmentDetailController {
	return &SalesConsignmentDetailControllerImpl{
		SalesConsignmentDetailService: salesConsignmentDetailService,
	}
}

func (dc *SalesConsignmentDetailControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	salesConsignmentDetailResponse, _ := dc.SalesConsignmentDetailService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesConsignmentDetailResponse.Code, salesConsignmentDetailResponse)
}

func (dc *SalesConsignmentDetailControllerImpl) FindAll(ctx echo.Context) error {
	salesConsignmentDetailResponse, _ := dc.SalesConsignmentDetailService.FindAll(ctx)

	return ctx.JSON(salesConsignmentDetailResponse.Code, salesConsignmentDetailResponse)
}

func (dc *SalesConsignmentDetailControllerImpl) Create(ctx echo.Context) error {
	salesConsignmentDetailResponse, _ := dc.SalesConsignmentDetailService.Create(ctx)

	return ctx.JSON(salesConsignmentDetailResponse.Code, salesConsignmentDetailResponse)
}

func (dc *SalesConsignmentDetailControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	salesConsignmentDetailResponse, _ := dc.SalesConsignmentDetailService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesConsignmentDetailResponse.Code, salesConsignmentDetailResponse)
}

func (dc *SalesConsignmentDetailControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	salesConsignmentDetailResponse, _ := dc.SalesConsignmentDetailService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(salesConsignmentDetailResponse.Code, salesConsignmentDetailResponse)
}