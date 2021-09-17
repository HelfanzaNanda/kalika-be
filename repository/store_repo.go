package repository

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
)

type (
	StoreRepository interface{
		Create(ctx echo.Context, db *gorm.DB, store *domain.Store) (domain.Store, error)
		Update(ctx echo.Context, db *gorm.DB, store *domain.Store) (domain.Store, error)
		Delete(ctx echo.Context, db *gorm.DB, store *domain.Store) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Store, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Store, error)
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

