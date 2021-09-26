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
	UnitRepository interface{
		Create(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (domain.Unit, error)
		Update(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (domain.Unit, error)
		Delete(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Unit, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Unit, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.UnitDatatable, int64, int64, error)
	}

	UnitRepositoryImpl struct {

	}
)

func NewUnitRepository() UnitRepository {
	return &UnitRepositoryImpl{}
}

func (repository UnitRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (domain.Unit, error) {
	db.Create(&unit)
	unitRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(unit.Id))
	return unitRes, nil
}

func (repository UnitRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (domain.Unit, error) {
	db.Where("id = ?", unit.Id).Updates(&unit)
	unitRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(unit.Id))
	return unitRes, nil
}

func (repository UnitRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, unit *domain.Unit) (bool, error) {
	results := db.Where("id = ?", unit.Id).Delete(&unit)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|unit tidak ditemukan")
	}
	return true, nil
}

func (repository UnitRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (unitRes domain.Unit, err error) {
	results := db.Where(key+" = ?", value).First(&unitRes)
	if results.RowsAffected < 1 {
		return unitRes, errors.New("NOT_FOUND|unit tidak ditemukan")
	}
	return unitRes, nil
}

func (repository UnitRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (unitRes []domain.Unit, err error) {
	db.Find(&unitRes)
	return unitRes, nil
}


func (repository UnitRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.UnitDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("units")
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