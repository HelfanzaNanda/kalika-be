package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type RoleController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type RoleControllerImpl struct {
		RoleService services.RoleService
}

func NewRoleController(roleService services.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

func (dc *RoleControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	roleResponse, _ := dc.RoleService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(roleResponse.Code, roleResponse)
}

func (dc *RoleControllerImpl) FindAll(ctx echo.Context) error {
	roleResponse, _ := dc.RoleService.FindAll(ctx)

	return ctx.JSON(roleResponse.Code, roleResponse)
}

func (dc *RoleControllerImpl) Create(ctx echo.Context) error {
	roleResponse, _ := dc.RoleService.Create(ctx)

	return ctx.JSON(roleResponse.Code, roleResponse)
}

func (dc *RoleControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	roleResponse, _ := dc.RoleService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(roleResponse.Code, roleResponse)
}

func (dc *RoleControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	roleResponse, _ := dc.RoleService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(roleResponse.Code, roleResponse)
}

func (dc *RoleControllerImpl) Datatable(ctx echo.Context) error {
	roleResponse, _ := dc.RoleService.Datatable(ctx)

	return ctx.JSON(202, roleResponse)
}