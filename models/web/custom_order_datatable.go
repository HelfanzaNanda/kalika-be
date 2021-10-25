package web

import "kalika-be/models/domain"

type CustomOrderDatatable struct {
	domain.CustomOrder
	PaymentMethodName string `json:"payment_method_name"`
	ProductName       string `json:"product_name"`
	CreatedByName     string `json:"created_by_name"`
	Action            string `json:"action"`
}
