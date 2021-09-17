package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type RoleHasPermissionController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
type RoleHasPermissionControllerImpl struct {
		RoleHasPermissionService services.RoleHasPermissionService
}

func NewRoleHasPermissionController(roleHasPermissionService services.RoleHasPermissionService) RoleHasPermissionController {
	return &RoleHasPermissionControllerImpl{
		RoleHasPermissionService: roleHasPermissionService,
	}
}

func (dc *RoleHasPermissionControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	roleHasPermissionResponse, _ := dc.RoleHasPermissionService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(roleHasPermissionResponse.Code, roleHasPermissionResponse)
}

func (dc *RoleHasPermissionControllerImpl) FindAll(ctx echo.Context) error {
	roleHasPermissionResponse, _ := dc.RoleHasPermissionService.FindAll(ctx)

	return ctx.JSON(roleHasPermissionResponse.Code, roleHasPermissionResponse)
}

func (dc *RoleHasPermissionControllerImpl) Create(ctx echo.Context) error {
	roleHasPermissionResponse, _ := dc.RoleHasPermissionService.Create(ctx)

	return ctx.JSON(roleHasPermissionResponse.Code, roleHasPermissionResponse)
}

func (dc *RoleHasPermissionControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	roleHasPermissionResponse, _ := dc.RoleHasPermissionService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(roleHasPermissionResponse.Code, roleHasPermissionResponse)
}

func (dc *RoleHasPermissionControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	roleHasPermissionResponse, _ := dc.RoleHasPermissionService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(roleHasPermissionResponse.Code, roleHasPermissionResponse)
}