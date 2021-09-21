package repository

import (
	"errors"
	//"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
)

type (
	RoleHasPermissionRepository interface{
		Create(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.RoleHasPermission) (domain.RoleHasPermission, error)
		Update(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.RoleHasPermission) (domain.RoleHasPermission, error)
		Delete(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.RoleHasPermission) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.RoleHasPermission, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]web.RoleHasPermissionGet, error)
	}

	RoleHasPermissionRepositoryImpl struct {

	}
)

func NewRoleHasPermissionRepository() RoleHasPermissionRepository {
	return &RoleHasPermissionRepositoryImpl{}
}

func (repository RoleHasPermissionRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.RoleHasPermission) (domain.RoleHasPermission, error) {
	db.Create(&roleHasPermission)
	roleHasPermissionRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(roleHasPermission.Id))
	return roleHasPermissionRes, nil
}

func (repository RoleHasPermissionRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.RoleHasPermission) (domain.RoleHasPermission, error) {
	db.Where("id = ?", roleHasPermission.Id).Updates(&roleHasPermission)
	roleHasPermissionRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(roleHasPermission.Id))
	return roleHasPermissionRes, nil
}

func (repository RoleHasPermissionRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.RoleHasPermission) (bool, error) {
	results := db.Where("id = ?", roleHasPermission.Id).Delete(&roleHasPermission)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|roleHasPermission tidak ditemukan")
	}
	return true, nil
}

func (repository RoleHasPermissionRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (roleHasPermissionRes domain.RoleHasPermission, err error) {
	results := db.Where(key+" = ?", value).First(&roleHasPermissionRes)
	if results.RowsAffected < 1 {
		return roleHasPermissionRes, errors.New("NOT_FOUND|roleHasPermission tidak ditemukan")
	}
	return roleHasPermissionRes, nil
}

func (repository RoleHasPermissionRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (roleHasPermissionRes []web.RoleHasPermissionGet, err error) {
	qry := db.Table("role_has_permissions").Select("role_has_permissions.*", "permissions.name as permission_name")
	for k, v := range ctx.QueryParams() {
		if v[0] != "" {
			qry = qry.Where(k+" = ?", v[0])
		}
	}
	qry.Joins("JOIN permissions on permissions.id = role_has_permissions.permission_id").Scan(&roleHasPermissionRes)
	return roleHasPermissionRes, nil
}

