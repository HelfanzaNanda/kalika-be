package controllers

import (
	//"fmt"
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"kalika-be/services"
	//"net/http"
)

type PaymentController interface {
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Datatable(ctx echo.Context) error
}
type PaymentControllerImpl struct {
		PaymentService services.PaymentService
}

func NewPaymentController(paymentService services.PaymentService) PaymentController {
	return &PaymentControllerImpl{
		PaymentService: paymentService,
	}
}

func (dc *PaymentControllerImpl) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	paymentResponse, _ := dc.PaymentService.FindById(ctx, helpers.StringToInt(id))

	return ctx.JSON(paymentResponse.Code, paymentResponse)
}

func (dc *PaymentControllerImpl) FindAll(ctx echo.Context) error {
	paymentResponse, _ := dc.PaymentService.FindAll(ctx)

	return ctx.JSON(paymentResponse.Code, paymentResponse)
}

func (dc *PaymentControllerImpl) Create(ctx echo.Context) error {
	paymentResponse, _ := dc.PaymentService.Create(ctx)

	return ctx.JSON(paymentResponse.Code, paymentResponse)
}

func (dc *PaymentControllerImpl) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	paymentResponse, _ := dc.PaymentService.Update(ctx, helpers.StringToInt(id))

	return ctx.JSON(paymentResponse.Code, paymentResponse)
}

func (dc *PaymentControllerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	paymentResponse, _ := dc.PaymentService.Delete(ctx, helpers.StringToInt(id))
	return ctx.JSON(paymentResponse.Code, paymentResponse)
}

func (dc *PaymentControllerImpl) Datatable(ctx echo.Context) error {
	paymentResponse, _ := dc.PaymentService.Datatable(ctx)
	return ctx.JSON(202, paymentResponse)
}