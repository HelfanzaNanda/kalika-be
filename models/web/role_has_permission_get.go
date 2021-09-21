package web


type RoleHasPermissionGet struct {
	Id int `json:"id" gorm:"AUTO_INCREMENT"`
	RoleId string `json:"role_id"`
	PermissionId string `json:"permission_id"`
	PermissionName string `json:"permission_name"`
}