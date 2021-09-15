package domain

import (
	"time"
)

type Store struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Code string `json:"code"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	PicName string `json:"pic_name"`
	PicPhone string `json:"pic_phone"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}