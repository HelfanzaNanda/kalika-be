package web

type RoleDatatable struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Name string `json:"name"`
	Action string `json:"action"`
}