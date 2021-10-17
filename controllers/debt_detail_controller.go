package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type DebtDetailController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type DebtDetailControllerImpl struct {
		DebtDetailService services.DebtDetailService
}

func NewDebtDetailController(debtDetailService services.DebtDetailService) DebtDetailController {
	return &DebtDetailControllerImpl{
		DebtDetailService: debtDetailService,
	}
}

func (dc *DebtDetailControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	debtDetailResponse, _ := dc.DebtDetailService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(debtDetailResponse.Code, debtDetailResponse)
}

func (dc *DebtDetailControllerImpl) FindAll(ctx echo.Context) error {
	debtDetailResponse, _ := dc.DebtDetailService.FindAll(ctx)

	return ctx.JSON(debtDetailResponse.Code, debtDetailResponse)
}

func (dc *DebtDetailControllerImpl) Create(ctx echo.Context) error {
	debtDetailResponse, _ := dc.DebtDetailService.Create(ctx)

	return ctx.JSON(debtDetailResponse.Code, debtDetailResponse)
}

func (dc *DebtDetailControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	debtDetailResponse, _ := dc.DebtDetailService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(debtDetailResponse.Code, debtDetailResponse)
}

func (dc *DebtDetailControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	debtDetailResponse, _ := dc.DebtDetailService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(debtDetailResponse.Code, debtDetailResponse)
}

func (dc *DebtDetailControllerImpl) Datatable(ctx echo.Context) error {
	debtDetailResponse, _ := dc.DebtDetailService.Datatable(ctx)
	return ctx.JSON(202, debtDetailResponse)
}