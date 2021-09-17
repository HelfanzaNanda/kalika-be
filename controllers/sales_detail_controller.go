package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type SalesDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type SalesDetailControllerImpl struct {
		SalesDetailService services.SalesDetailService
}

func NewSalesDetailController(salesDetailService services.SalesDetailService) SalesDetailController {
	return &SalesDetailControllerImpl{
		SalesDetailService: salesDetailService,
	}
}

func (dc *SalesDetailControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	salesDetailResponse, _ := dc.SalesDetailService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesDetailResponse.Code, salesDetailResponse)
}

func (dc *SalesDetailControllerImpl) FindAll(ctx echo.Context) error {
	salesDetailResponse, _ := dc.SalesDetailService.FindAll(ctx)

	return ctx.JSON(salesDetailResponse.Code, salesDetailResponse)
}

func (dc *SalesDetailControllerImpl) Create(ctx echo.Context) error {
	salesDetailResponse, _ := dc.SalesDetailService.Create(ctx)

	return ctx.JSON(salesDetailResponse.Code, salesDetailResponse)
}

func (dc *SalesDetailControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	salesDetailResponse, _ := dc.SalesDetailService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesDetailResponse.Code, salesDetailResponse)
}

func (dc *SalesDetailControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	salesDetailResponse, _ := dc.SalesDetailService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(salesDetailResponse.Code, salesDetailResponse)
}