package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"kalika-be/controllers"
)

func Routes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	api := e.Group("/api")
	userController := controllers.NewUserController()

	api.GET("/users", userController.FindAll)
	api.GET("/users/:id", userController.FindById)
	api.POST("/users", userController.Create)
	api.PUT("/users/:id", userController.Update)
	api.DELETE("/users/:id", userController.Delete)

	return e
}