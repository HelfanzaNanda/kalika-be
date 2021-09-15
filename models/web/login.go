package web

type Login struct {
	Name string `json:"name"`
	Username string `json:"username"`
	RoleId int `json:"role_id""`
	StoreId int `json:"store_id""`
	Token string `json:"token"`
}
