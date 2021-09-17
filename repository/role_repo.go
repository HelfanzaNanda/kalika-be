package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	RoleRepository interface{
		Create(ctx echo.Context, db *gorm.DB, role *domain.Role) (domain.Role, error)
		Update(ctx echo.Context, db *gorm.DB, role *domain.Role) (domain.Role, error)
		Delete(ctx echo.Context, db *gorm.DB, role *domain.Role) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Role, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Role, error)
	}

	RoleRepositoryImpl struct {

	}
)

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository RoleRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, role *domain.Role) (domain.Role, error) {
	db.Create(&role)
	roleRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(role.Id))
	return roleRes, nil
}

func (repository RoleRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, role *domain.Role) (domain.Role, error) {
	db.Where("id = ?", role.Id).Updates(&role)
	roleRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(role.Id))
	return roleRes, nil
}

func (repository RoleRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, role *domain.Role) (bool, error) {
	results := db.Where("id = ?", role.Id).Delete(&role)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|role tidak ditemukan")
	}
	return true, nil
}

func (repository RoleRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (roleRes domain.Role, err error) {
	results := db.Where(key+" = ?", value).First(&roleRes)
	if results.RowsAffected < 1 {
		return roleRes, errors.New("NOT_FOUND|role tidak ditemukan")
	}
	return roleRes, nil
}

func (repository RoleRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (roleRes []domain.Role, err error) {
	db.Find(&roleRes)
	return roleRes, nil
}

