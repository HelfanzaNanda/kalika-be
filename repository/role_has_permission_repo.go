package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	RoleHasPermissionRepository interface{
		Create(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.RoleHasPermission) (domain.RoleHasPermission, error)
		Update(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.RoleHasPermission) (domain.RoleHasPermission, error)
		Delete(ctx echo.Context, db *gorm.DB, roleHasPermission *domain.RoleHasPermission) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.RoleHasPermission, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.RoleHasPermission, error)
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

func (repository RoleHasPermissionRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (roleHasPermissionRes []domain.RoleHasPermission, err error) {
	db.Find(&roleHasPermissionRes)
	return roleHasPermissionRes, nil
}

