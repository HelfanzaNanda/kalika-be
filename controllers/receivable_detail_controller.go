package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type ReceivableDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type ReceivableDetailControllerImpl struct {
		ReceivableDetailService services.ReceivableDetailService
}

func NewReceivableDetailController(receivableDetailService services.ReceivableDetailService) ReceivableDetailController {
	return &ReceivableDetailControllerImpl{
		ReceivableDetailService: receivableDetailService,
	}
}

func (dc *ReceivableDetailControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	receivableDetailResponse, _ := dc.ReceivableDetailService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(receivableDetailResponse.Code, receivableDetailResponse)
}

func (dc *ReceivableDetailControllerImpl) FindAll(ctx echo.Context) error {
	receivableDetailResponse, _ := dc.ReceivableDetailService.FindAll(ctx)

	return ctx.JSON(receivableDetailResponse.Code, receivableDetailResponse)
}

func (dc *ReceivableDetailControllerImpl) Create(ctx echo.Context) error {
	receivableDetailResponse, _ := dc.ReceivableDetailService.Create(ctx)

	return ctx.JSON(receivableDetailResponse.Code, receivableDetailResponse)
}

func (dc *ReceivableDetailControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	receivableDetailResponse, _ := dc.ReceivableDetailService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(receivableDetailResponse.Code, receivableDetailResponse)
}

func (dc *ReceivableDetailControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	receivableDetailResponse, _ := dc.ReceivableDetailService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(receivableDetailResponse.Code, receivableDetailResponse)
}