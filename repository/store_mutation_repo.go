package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
)

type (
	StoreMutationRepository interface{
		Create(echo.Context, *gorm.DB, *domain.StoreMutation) (web.StoreMutationGet, error)
		Update(echo.Context, *gorm.DB, *domain.StoreMutation) (web.StoreMutationGet, error)
		Delete(echo.Context, *gorm.DB, *domain.StoreMutation) (bool, error)
		FindById(echo.Context, *gorm.DB, string, string) (web.StoreMutationGet, error)
		FindAll(echo.Context, *gorm.DB) ([]domain.StoreMutation, error)
		Datatable(echo.Context, *gorm.DB, string, string, string, string) ([]web.StoreMutationDatatable, int64, int64, error)
		
	}

	StoreMutationRepositoryImpl struct {

	}
)

func NewStoreMutationRepository() StoreMutationRepository {
	return &StoreMutationRepositoryImpl{}
}

func (repository *StoreMutationRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, storeMutation *domain.StoreMutation) (web.StoreMutationGet, error) {
	db.Create(&storeMutation)
	storeMutationRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(storeMutation.Id))
	return storeMutationRes, nil
}

func (repository *StoreMutationRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, storeMutation *domain.StoreMutation) (web.StoreMutationGet, error) {
	db.Where("id = ?", storeMutation.Id).Updates(&storeMutation)
	storeMutationRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(storeMutation.Id))
	return storeMutationRes, nil
}

func (repository *StoreMutationRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, storeMutation *domain.StoreMutation) (bool, error) {
	results := db.Where("id = ?", storeMutation.Id).Delete(&storeMutation)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|storeMutation tidak ditemukan")
	}
	return true, nil
}

func (repository *StoreMutationRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (storeMutationRes web.StoreMutationGet, err error) {
	results := db.Where(key+" = ?", value).First(&storeMutationRes.StoreMutation)
	if results.RowsAffected < 1 {
		return storeMutationRes, errors.New("NOT_FOUND|storeMutation tidak ditemukan")
	}
	return storeMutationRes, nil
}

func (repository *StoreMutationRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (storeMutationRes []domain.StoreMutation, err error) {
	db.Find(&storeMutationRes)
	return storeMutationRes, nil
}

func (repository *StoreMutationRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (recipeRes []web.StoreMutationDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("stock_opnames").Select("stock_opnames.*, users.name as created_by_name")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(stock_opnames.number LIKE ?)", "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Joins("JOIN users ON users.id = stock_opnames.created_by")
	qry.Order("stock_opnames.id desc")
	qry.Preload("Store").Find(&recipeRes)
	return recipeRes, totalData, totalFiltered, nil
}