package domain

import (
	"time"
)

type StockOpnameDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	StockOpnameId int `json:"stock_opname_id"`
	StoreId int `json:"store_id"`
	ProductId int `json:"product_id"`
	StockOnBook float64 `json:"stock_on_book"`
	StockOnPhysic float64 `json:"stock_on_physic"`
	Note string `json:"note"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}