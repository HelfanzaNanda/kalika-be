package repository

import (
	"errors"
	"kalika-be/models/domain"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	GeneralSettingRepository interface {
		CreateOrUpdate(ctx echo.Context, db *gorm.DB, item string, value string) (domain.GeneralSetting, error)
		Create(ctx echo.Context, db *gorm.DB, item string, value string) ([]domain.GeneralSetting, error)
		Update(ctx echo.Context, db *gorm.DB, item string, value string) ([]domain.GeneralSetting, error)
		FindById(ctx echo.Context, db *gorm.DB, item string) (domain.GeneralSetting, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.GeneralSetting, error)
	}

	GeneralSettingRepositoryImpl struct {

	}
)

func NewGeneralSettingRepository() GeneralSettingRepository {
	return &GeneralSettingRepositoryImpl{}
}

func (repository GeneralSettingRepositoryImpl) CreateOrUpdate(ctx echo.Context, db *gorm.DB, item string, value string) (domain.GeneralSetting, error) {

	
	generalSettingRes, _ := repository.FindById(ctx, db, item)
	if generalSettingRes.Id > 0 {
		repository.Update(ctx, db, item, value)
	}else {
		repository.Create(ctx, db, item, value)
	}
	return generalSettingRes, nil
}

func (repository GeneralSettingRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, item string, value string) ([]domain.GeneralSetting, error) {
	generalSetting := domain.GeneralSetting{}
	generalSetting.Item = item
	generalSetting.Value = value
	db.Create(&generalSetting)
	generalSettingRes, _ := repository.FindAll(ctx, db)
	return generalSettingRes, nil
}

func (repository GeneralSettingRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, item string, value string) ([]domain.GeneralSetting, error) {
	generalSetting := domain.GeneralSetting{}
	generalSetting.Item = item
	generalSetting.Value = value
	db.Where("item = ?", generalSetting.Item).Updates(&generalSetting)
	generalSettingRes, _ := repository.FindAll(ctx, db)
	return generalSettingRes, nil
}

func (repository GeneralSettingRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, item string) (res domain.GeneralSetting, err error) {
	results := db.Where("item = ?", item).First(&res)
	if results.RowsAffected < 1 {
		return res, errors.New("NOT_FOUND|cakeType tidak ditemukan")
	}
	return res, nil
}

func (repository GeneralSettingRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (res []domain.GeneralSetting, err error) {
	db.Find(&res)
	return res, nil
}