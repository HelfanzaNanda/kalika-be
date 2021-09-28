package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type CustomerController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type CustomerControllerImpl struct {
		CustomerService services.CustomerService
}

func NewCustomerController(customerService services.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

func (dc *CustomerControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	customerResponse, _ := dc.CustomerService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(customerResponse.Code, customerResponse)
}

func (dc *CustomerControllerImpl) FindAll(ctx echo.Context) error {
	customerResponse, _ := dc.CustomerService.FindAll(ctx)

	return ctx.JSON(customerResponse.Code, customerResponse)
}

func (dc *CustomerControllerImpl) Create(ctx echo.Context) error {
	customerResponse, _ := dc.CustomerService.Create(ctx)

	return ctx.JSON(customerResponse.Code, customerResponse)
}

func (dc *CustomerControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	customerResponse, _ := dc.CustomerService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(customerResponse.Code, customerResponse)
}

func (dc *CustomerControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	customerResponse, _ := dc.CustomerService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(customerResponse.Code, customerResponse)
}
func (dc *CustomerControllerImpl) Datatable(ctx echo.Context) error {
	customerResponse, _ := dc.CustomerService.Datatable(ctx)
	return ctx.JSON(202, customerResponse)
}