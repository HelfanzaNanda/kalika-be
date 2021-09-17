package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	ReceivableDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (domain.ReceivableDetail, error)
		Update(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (domain.ReceivableDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.ReceivableDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.ReceivableDetail, error)
	}

	ReceivableDetailRepositoryImpl struct {

	}
)

func NewReceivableDetailRepository() ReceivableDetailRepository {
	return &ReceivableDetailRepositoryImpl{}
}

func (repository ReceivableDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (domain.ReceivableDetail, error) {
	db.Create(&receivableDetail)
	receivableDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(receivableDetail.Id))
	return receivableDetailRes, nil
}

func (repository ReceivableDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (domain.ReceivableDetail, error) {
	db.Where("id = ?", receivableDetail.Id).Updates(&receivableDetail)
	receivableDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(receivableDetail.Id))
	return receivableDetailRes, nil
}

func (repository ReceivableDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (bool, error) {
	results := db.Where("id = ?", receivableDetail.Id).Delete(&receivableDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|receivableDetail tidak ditemukan")
	}
	return true, nil
}

func (repository ReceivableDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (receivableDetailRes domain.ReceivableDetail, err error) {
	results := db.Where(key+" = ?", value).First(&receivableDetailRes)
	if results.RowsAffected < 1 {
		return receivableDetailRes, errors.New("NOT_FOUND|receivableDetail tidak ditemukan")
	}
	return receivableDetailRes, nil
}

func (repository ReceivableDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (receivableDetailRes []domain.ReceivableDetail, err error) {
	db.Find(&receivableDetailRes)
	return receivableDetailRes, nil
}

