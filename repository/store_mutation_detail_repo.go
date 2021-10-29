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
	StoreMutationDetailRepository interface{
		Create(echo.Context, *gorm.DB, []domain.StoreMutationDetail) ([]web.StoreMutationDetailGet, error)
		Update(echo.Context, *gorm.DB, *domain.StoreMutationDetail) (domain.StoreMutationDetail, error)
		Delete(echo.Context, *gorm.DB, *domain.StoreMutationDetail) (bool, error)
		DeleteByStoreMutation(echo.Context, *gorm.DB, int) (bool, error)
		FindById(echo.Context, *gorm.DB, string, string) (domain.StoreMutationDetail, error)
		FindAll(echo.Context, *gorm.DB, map[string][]string) ([]web.StoreMutationDetailGet, error)
	}

	StoreMutationDetailRepositoryImpl struct {

	}
)

func NewStoreMutationDetailRepository() StoreMutationDetailRepository {
	return &StoreMutationDetailRepositoryImpl{}
}

func (repository *StoreMutationDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, stockOpnameDetail []domain.StoreMutationDetail) (res []web.StoreMutationDetailGet, err error) {
	for _, val := range stockOpnameDetail {
		db.Create(&val)
	}

	detailSearch := make(map[string][]string)
	detailSearch["store_mutation_id"] = append(detailSearch["store_mutation_id"], helpers.IntToString(stockOpnameDetail[0].StoreMutationId))

	res, _ = repository.FindAll(ctx, db, detailSearch)

	return res, nil
}

func (repository *StoreMutationDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, stockOpnameDetail *domain.StoreMutationDetail) (domain.StoreMutationDetail, error) {
	db.Where("id = ?", stockOpnameDetail.Id).Updates(&stockOpnameDetail)
	stockOpnameDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(stockOpnameDetail.Id))
	return stockOpnameDetailRes, nil
}

func (repository *StoreMutationDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, stockOpnameDetail *domain.StoreMutationDetail) (bool, error) {
	results := db.Where("id = ?", stockOpnameDetail.Id).Delete(&stockOpnameDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|stockOpnameDetail tidak ditemukan")
	}
	return true, nil
}

func (repository *StoreMutationDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (stockOpnameDetailRes domain.StoreMutationDetail, err error) {
	results := db.Where(key+" = ?", value).First(&stockOpnameDetailRes)
	if results.RowsAffected < 1 {
		return stockOpnameDetailRes, errors.New("NOT_FOUND|stockOpnameDetail tidak ditemukan")
	}
	return stockOpnameDetailRes, nil
}

func (repository *StoreMutationDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (stockOpnameDetailRes []web.StoreMutationDetailGet, err error) {
	results := db.Table("store_mutation_details").Preload("Product").Preload("Store")
	for k, v := range params {
		if v[0] != "" && k != "id" {
			results = results.Where(k+" = ?", v[0])
		}
	}
	results.Find(&stockOpnameDetailRes)
	return stockOpnameDetailRes, nil
}

func (repository *StoreMutationDetailRepositoryImpl) DeleteByStoreMutation(ctx echo.Context, db *gorm.DB, stockOpnameId int) (bool, error) {
	stockOpnameDetail := domain.StoreMutationDetail{}
	results := db.Where("store_mutation_id = ?", stockOpnameId).Delete(&stockOpnameDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|stockOpnameDetail tidak ditemukan")
	}
	return true, nil
}