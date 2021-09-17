package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	StoreConsignmentRepository interface{
		Create(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (domain.StoreConsignment, error)
		Update(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (domain.StoreConsignment, error)
		Delete(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.StoreConsignment, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.StoreConsignment, error)
	}

	StoreConsignmentRepositoryImpl struct {

	}
)

func NewStoreConsignmentRepository() StoreConsignmentRepository {
	return &StoreConsignmentRepositoryImpl{}
}

func (repository StoreConsignmentRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (domain.StoreConsignment, error) {
	db.Create(&storeConsignment)
	storeConsignmentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(storeConsignment.Id))
	return storeConsignmentRes, nil
}

func (repository StoreConsignmentRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (domain.StoreConsignment, error) {
	db.Where("id = ?", storeConsignment.Id).Updates(&storeConsignment)
	storeConsignmentRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(storeConsignment.Id))
	return storeConsignmentRes, nil
}

func (repository StoreConsignmentRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, storeConsignment *domain.StoreConsignment) (bool, error) {
	results := db.Where("id = ?", storeConsignment.Id).Delete(&storeConsignment)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|storeConsignment tidak ditemukan")
	}
	return true, nil
}

func (repository StoreConsignmentRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (storeConsignmentRes domain.StoreConsignment, err error) {
	results := db.Where(key+" = ?", value).First(&storeConsignmentRes)
	if results.RowsAffected < 1 {
		return storeConsignmentRes, errors.New("NOT_FOUND|storeConsignment tidak ditemukan")
	}
	return storeConsignmentRes, nil
}

func (repository StoreConsignmentRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (storeConsignmentRes []domain.StoreConsignment, err error) {
	db.Find(&storeConsignmentRes)
	return storeConsignmentRes, nil
}

