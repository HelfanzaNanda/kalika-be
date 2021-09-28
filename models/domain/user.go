package domain

import "time"

type User struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId int `json:"role_id"`
	StoreId int `json:"store_id"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}