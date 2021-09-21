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
	DivisionRepository interface{
		Create(ctx echo.Context, db *gorm.DB, division *domain.Division) (domain.Division, error)
		Update(ctx echo.Context, db *gorm.DB, division *domain.Division) (domain.Division, error)
		Delete(ctx echo.Context, db *gorm.DB, division *domain.Division) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Division, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Division, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.DivisionDatatable, int64, int64, error)
	}

	DivisionRepositoryImpl struct {

	}
)

func NewDivisionRepository() DivisionRepository {
	return &DivisionRepositoryImpl{}
}

func (repository DivisionRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, division *domain.Division) (domain.Division, error) {
	db.Create(&division)
	divisionRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(division.Id))
	return divisionRes, nil
}

func (repository DivisionRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, division *domain.Division) (domain.Division, error) {
	db.Where("id = ?", division.Id).Updates(&division)
	divisionRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(division.Id))
	return divisionRes, nil
}

func (repository DivisionRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, division *domain.Division) (bool, error) {
	results := db.Where("id = ?", division.Id).Delete(&division)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|division tidak ditemukan")
	}
	return true, nil
}

func (repository DivisionRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (divisionRes domain.Division, err error) {
	results := db.Where(key+" = ?", value).First(&divisionRes)
	if results.RowsAffected < 1 {
		return divisionRes, errors.New("NOT_FOUND|division tidak ditemukan")
	}
	return divisionRes, nil
}

func (repository DivisionRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (divisionRes []domain.Division, err error) {
	db.Find(&divisionRes)
	return divisionRes, nil
}

func (repository DivisionRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (divisionRes []web.DivisionDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("divisions").Select("id, name, active, created_at, updated_at")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(id = ? OR name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Find(&divisionRes)
	return divisionRes, totalData, totalFiltered, nil
}

