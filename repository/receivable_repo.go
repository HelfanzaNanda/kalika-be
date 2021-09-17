package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	ReceivableRepository interface{
		Create(ctx echo.Context, db *gorm.DB, receivable *domain.Receivable) (domain.Receivable, error)
		Update(ctx echo.Context, db *gorm.DB, receivable *domain.Receivable) (domain.Receivable, error)
		Delete(ctx echo.Context, db *gorm.DB, receivable *domain.Receivable) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Receivable, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Receivable, error)
	}

	ReceivableRepositoryImpl struct {

	}
)

func NewReceivableRepository() ReceivableRepository {
	return &ReceivableRepositoryImpl{}
}

func (repository ReceivableRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, receivable *domain.Receivable) (domain.Receivable, error) {
	db.Create(&receivable)
	receivableRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(receivable.Id))
	return receivableRes, nil
}

func (repository ReceivableRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, receivable *domain.Receivable) (domain.Receivable, error) {
	db.Where("id = ?", receivable.Id).Updates(&receivable)
	receivableRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(receivable.Id))
	return receivableRes, nil
}

func (repository ReceivableRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, receivable *domain.Receivable) (bool, error) {
	results := db.Where("id = ?", receivable.Id).Delete(&receivable)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|receivable tidak ditemukan")
	}
	return true, nil
}

func (repository ReceivableRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (receivableRes domain.Receivable, err error) {
	results := db.Where(key+" = ?", value).First(&receivableRes)
	if results.RowsAffected < 1 {
		return receivableRes, errors.New("NOT_FOUND|receivable tidak ditemukan")
	}
	return receivableRes, nil
}

func (repository ReceivableRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (receivableRes []domain.Receivable, err error) {
	db.Find(&receivableRes)
	return receivableRes, nil
}

