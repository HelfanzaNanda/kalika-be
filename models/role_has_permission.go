package models

import "time"

type RoleHasPermission struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	RoleId string `json:"role_id"`
	PermissionId string `json:"permission_id"`
	CreatedAt time.Time`json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}