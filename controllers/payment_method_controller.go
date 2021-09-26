package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type PaymentMethodController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type PaymentMethodControllerImpl struct {
		PaymentMethodService services.PaymentMethodService
}

func NewPaymentMethodController(paymentMethodService services.PaymentMethodService) PaymentMethodController {
	return &PaymentMethodControllerImpl{
		PaymentMethodService: paymentMethodService,
	}
}

func (dc *PaymentMethodControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	paymentMethodResponse, _ := dc.PaymentMethodService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(paymentMethodResponse.Code, paymentMethodResponse)
}

func (dc *PaymentMethodControllerImpl) FindAll(ctx echo.Context) error {
	paymentMethodResponse, _ := dc.PaymentMethodService.FindAll(ctx)

	return ctx.JSON(paymentMethodResponse.Code, paymentMethodResponse)
}

func (dc *PaymentMethodControllerImpl) Create(ctx echo.Context) error {
	paymentMethodResponse, _ := dc.PaymentMethodService.Create(ctx)

	return ctx.JSON(paymentMethodResponse.Code, paymentMethodResponse)
}

func (dc *PaymentMethodControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	paymentMethodResponse, _ := dc.PaymentMethodService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(paymentMethodResponse.Code, paymentMethodResponse)
}

func (dc *PaymentMethodControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	paymentMethodResponse, _ := dc.PaymentMethodService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(paymentMethodResponse.Code, paymentMethodResponse)
}

func (dc *PaymentMethodControllerImpl) Datatable(ctx echo.Context) error {
	paymentMethodResponse, _ := dc.PaymentMethodService.Datatable(ctx)
	return ctx.JSON(202, paymentMethodResponse)
}