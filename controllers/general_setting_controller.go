package controllers

import (
	//"fmt"
	"kalika-be/services"

	"github.com/labstack/echo"
	//"net/http"
)

type GeneralSettingController interface {
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
}
type GeneralSettingControllerImpl struct {
	GeneralSettingService services.GeneralSettingService
}

func NewGeneralSettingController(GeneralSettingService services.GeneralSettingService) GeneralSettingController {
	return &GeneralSettingControllerImpl{
		GeneralSettingService: GeneralSettingService,
	}
}

func (dc *GeneralSettingControllerImpl) FindAll(ctx echo.Context) error {
	generalSettingRes, _ := dc.GeneralSettingService.FindAll(ctx)

	return ctx.JSON(generalSettingRes.Code, generalSettingRes)
}

func (dc *GeneralSettingControllerImpl) Create(ctx echo.Context) error {
	generalSettingRes, _ := dc.GeneralSettingService.Create(ctx)

	return ctx.JSON(generalSettingRes.Code, generalSettingRes)
}
