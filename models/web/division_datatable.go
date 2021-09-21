package web


type DivisionDatatable struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Name string `json:"name"`
	Active bool `json:"active"`
	Action string `json:"action"`
}