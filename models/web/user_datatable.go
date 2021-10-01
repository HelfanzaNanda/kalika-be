package web

type UserDatatable struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	Name string `json:"name"`
	Username string `json:"username"`
	RoleId string `json:"role_id"`
	RoleName string `json:"role_name"`
	StoreId string `json:"store_id"`
	StoreName string `json:"store_name"`
	Action string `json:"action"`
}