package web

import "kalika-be/models/domain"

type SalesDetailPosGet struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	SalesId int `json:"sales_id"`
	ProductId int `json:"product_id"`
	Product domain.Product `json:"product"`
	Qty int `json:"qty"`
	DiscountPercentage int `json:"discount_percentage"`
	DiscountValue int `json:"discount_value"`
	Total float64 `json:"total"`
	UnitPrice int `json:"unit_price"`
}