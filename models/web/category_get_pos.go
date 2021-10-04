package web

type CategoryGetPos struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	DivisionId int `json:"division_id"`
	Name string `json:"name"`
	Active bool `json:"active"`
	TotalProduct int `json:"total_product"`
}