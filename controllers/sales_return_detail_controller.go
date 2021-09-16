package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type SalesReturnDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type SalesReturnDetailControllerImpl struct {
		SalesReturnDetailService services.SalesReturnDetailService
}

func NewSalesReturnDetailController(salesReturnDetailService services.SalesReturnDetailService) SalesReturnDetailController {
	return &SalesReturnDetailControllerImpl{
		SalesReturnDetailService: salesReturnDetailService,
	}
}

func (dc *SalesReturnDetailControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	salesReturnDetailResponse, _ := dc.SalesReturnDetailService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesReturnDetailResponse.Code, salesReturnDetailResponse)
}

func (dc *SalesReturnDetailControllerImpl) FindAll(ctx echo.Context) error {
	salesReturnDetailResponse, _ := dc.SalesReturnDetailService.FindAll(ctx)

	return ctx.JSON(salesReturnDetailResponse.Code, salesReturnDetailResponse)
}

func (dc *SalesReturnDetailControllerImpl) Create(ctx echo.Context) error {
	salesReturnDetailResponse, _ := dc.SalesReturnDetailService.Create(ctx)

	return ctx.JSON(salesReturnDetailResponse.Code, salesReturnDetailResponse)
}

func (dc *SalesReturnDetailControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	salesReturnDetailResponse, _ := dc.SalesReturnDetailService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(salesReturnDetailResponse.Code, salesReturnDetailResponse)
}

func (dc *SalesReturnDetailControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	salesReturnDetailResponse, _ := dc.SalesReturnDetailService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(salesReturnDetailResponse.Code, salesReturnDetailResponse)
}