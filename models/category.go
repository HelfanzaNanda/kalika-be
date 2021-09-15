package models
import (
	"time"
)

type Category struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	DivisionId int `json:"division_id"`
	Name string `json:"name"`
	Active bool `json:"active"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}