package models

import "time"

type Setting struct {
	Key       string     `json:"key,omitempty" gorm:"primary_key;type:varchar(100)"`
	Value     string     `json:"value,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	DeletedAt *time.Time `json:"-"`
}
