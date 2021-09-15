package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
	"kalika-be/controllers"
	"kalika-be/middlewares"
	"kalika-be/repository"
	"kalika-be/services"
)

func Routes(db *gorm.DB) *echo.Echo {
	//validate := validator.New()

	//USER THINGS
	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository, db)
	userController := controllers.NewUserController(userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	api := e.Group("/api")
	api.POST("/login", userController.Login)

	api.Use(middlewares.Auth)

	api.GET("/users", userController.FindAll)
	api.GET("/users/:id", userController.FindById)
	api.POST("/users", userController.Create)
	api.PUT("/users/:id", userController.Update)
	api.DELETE("/users/:id", userController.Delete)

	return e
}