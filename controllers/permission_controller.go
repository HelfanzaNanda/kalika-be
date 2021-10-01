package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type PermissionController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type PermissionControllerImpl struct {
		Permission services.PermissionService
}

func NewPermissionController(permissionService services.PermissionService) PermissionController {
	return &PermissionControllerImpl{
		Permission: permissionService,
	}
}

func (dc *PermissionControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	roleHasPermissionResponse, _ := dc.Permission.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(roleHasPermissionResponse.Code, roleHasPermissionResponse)
}

func (dc *PermissionControllerImpl) FindAll(ctx echo.Context) error {
	roleHasPermissionResponse, _ := dc.Permission.FindAll(ctx)

	return ctx.JSON(roleHasPermissionResponse.Code, roleHasPermissionResponse)
}

func (dc *PermissionControllerImpl) Create(ctx echo.Context) error {
	roleHasPermissionResponse, _ := dc.Permission.Create(ctx)

	return ctx.JSON(roleHasPermissionResponse.Code, roleHasPermissionResponse)
}

func (dc *PermissionControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	roleHasPermissionResponse, _ := dc.Permission.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(roleHasPermissionResponse.Code, roleHasPermissionResponse)
}

func (dc *PermissionControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	roleHasPermissionResponse, _ := dc.Permission.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(roleHasPermissionResponse.Code, roleHasPermissionResponse)
}