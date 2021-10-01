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
	PermissionRepository interface{
		Create(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.Permission) (domain.Permission, error)
		Update(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.Permission) (domain.Permission, error)
		Delete(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.Permission) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Permission, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]web.PermissionGet, error)
	}

	PermissionRepositoryImpl struct {

	}
)

func NewPermissionRepository() PermissionRepository {
	return &PermissionRepositoryImpl{}
}

func (repository PermissionRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.Permission) (domain.Permission, error) {
	db.Create(&roleHasPermission)
	permissionRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(roleHasPermission.Id))
	return permissionRes, nil
}

func (repository PermissionRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.Permission) (domain.Permission, error) {
	db.Where("id = ?", roleHasPermission.Id).Updates(&roleHasPermission)
	permissionRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(roleHasPermission.Id))
	return permissionRes, nil
}

func (repository PermissionRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.Permission) (bool, error) {
	results := db.Where("id = ?", roleHasPermission.Id).Delete(&roleHasPermission)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|roleHasPermission tidak ditemukan")
	}
	return true, nil
}

func (repository PermissionRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (permissionRes domain.Permission, err error) {
	results := db.Where(key+" = ?", value).First(&permissionRes)
	if results.RowsAffected < 1 {
		return permissionRes, errors.New("NOT_FOUND|roleHasPermission tidak ditemukan")
	}
	return permissionRes, nil
}

func (repository PermissionRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (permissionRes []web.PermissionGet, err error) {
	qry := db.Table("permissions").Select("id, name")
	for k, v := range ctx.QueryParams() {
		if v[0] != "" {
			qry = qry.Where(k+" = ?", v[0])
		}
	}
	qry.Find(&permissionRes)
	return permissionRes, nil
}

