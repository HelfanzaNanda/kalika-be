package controllers

import (
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	"net/http"
)

type (
	UserController interface {
		FindById(ctx echo.Context) error
		FindAll(ctx echo.Context) error
		Create(ctx echo.Context) error
		Update(ctx echo.Context) error
		Delete(ctx echo.Context) error
		Login(ctx echo.Context) error
	}

	UserControllerImpl struct {
		UserService services.UserService
	}
)

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (uc *UserControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	userResponse, _ := uc.UserService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(userResponse.Code, userResponse)
}

func (uc *UserControllerImpl) FindAll(ctx echo.Context) error {
	userResponse, _ := uc.UserService.FindAll(ctx)

	return ctx.JSON(userResponse.Code, userResponse)
}

func (uc *UserControllerImpl) Create(ctx echo.Context) error {
	userResponse, _ := uc.UserService.Create(ctx)

	return ctx.JSON(userResponse.Code, userResponse)
}

func (uc *UserControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	userResponse, _ := uc.UserService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(userResponse.Code, userResponse)
}

func (uc *UserControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	userResponse, _ := uc.UserService.Delete(ctx, helpers.StringToInt(id))

	return ctx.JSON(userResponse.Code, userResponse)
}

func (uc *UserControllerImpl) Login(ctx echo.Context) error {
	userResponse, err := uc.UserService.Login(ctx)

	if err != nil {
		if err.Error() == "NOT_FOUND" {
			return ctx.JSON(http.StatusNotFound, userResponse)
		}
	}
	return ctx.JSON(http.StatusOK, userResponse)
}