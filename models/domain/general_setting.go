package domain

import (
	"time"
)

type GeneralSetting struct {
	Id        int       `json:"id" gorm:"AUTO_INCREMENT"`
	Item      string    `json:"item"`
	Value     string    `json:"value"`
	Filepath  string    `json:"filepath"`
	Filename  string    `json:"filename"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2020-03-16T13:55:09.598136+07:00"`
}
