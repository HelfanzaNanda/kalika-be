package web

type DivisionDatatable struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Name string `json:"name"`
	Type string `json:"type"`
	Active bool `json:"active"`
	Action string `json:"action"`
}