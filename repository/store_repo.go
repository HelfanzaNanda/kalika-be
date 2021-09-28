package repository

import (
	"errors"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	StoreRepository interface{
		Create(ctx echo.Context, db *gorm.DB, store *domain.Store) (domain.Store, error)
		Update(ctx echo.Context, db *gorm.DB, store *domain.Store) (domain.Store, error)
		Delete(ctx echo.Context, db *gorm.DB, store *domain.Store) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Store, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Store, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.StoreDatatable, int64, int64, error)
	}

	StoreRepositoryImpl struct {

	}
)

func NewStoreRepository() StoreRepository {
	return &StoreRepositoryImpl{}
}

func (repository StoreRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, store *domain.Store) (domain.Store, error) {
	db.Create(&store)
	storeRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(store.Id))
	return storeRes, nil
}

func (repository StoreRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, store *domain.Store) (domain.Store, error) {
	db.Where("id = ?", store.Id).Updates(&store)
	storeRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(store.Id))
	return storeRes, nil
}

func (repository StoreRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, store *domain.Store) (bool, error) {
	results := db.Where("id = ?", store.Id).Delete(&store)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|store tidak ditemukan")
	}
	return true, nil
}

func (repository StoreRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (storeRes domain.Store, err error) {
	results := db.Where(key+" = ?", value).First(&storeRes)
	if results.RowsAffected < 1 {
		return storeRes, errors.New("NOT_FOUND|store tidak ditemukan")
	}
	return storeRes, nil
}

func (repository StoreRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (storeRes []domain.Store, err error) {
	db.Find(&storeRes)
	return storeRes, nil
}

func (repository StoreRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.StoreDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("stores")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(id = ? OR name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}