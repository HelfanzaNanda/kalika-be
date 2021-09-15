package domain
import (
	"time"
)

type PurchaseReturnDetail struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	PurchaseReturnId int `json:"purchase_return_id"`
	RawMaterialId int `json:"raw_material_id"`
	Qty int `json:"qty"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}