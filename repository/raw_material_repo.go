package repository

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
)

type (
	RawMaterialRepository interface {
		Create(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (domain.RawMaterial, error)
		Update(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (domain.RawMaterial, error)
		Delete(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.RawMaterial, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]web.RawMaterialGet, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.RawMaterialDatatable, int64, int64, error)
	}

	RawMaterialRepositoryImpl struct {
	}
)

func NewRawMaterialRepository() RawMaterialRepository {
	return &RawMaterialRepositoryImpl{}
}

func (repository RawMaterialRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (domain.RawMaterial, error) {
	db.Create(&rawMaterial)
	rawMaterialRes, _ := repository.FindById(ctx, db, "id", helpers.IntToString(rawMaterial.Id))
	return rawMaterialRes, nil
}

func (repository RawMaterialRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (domain.RawMaterial, error) {
	db.Where("id = ?", rawMaterial.Id).Updates(&rawMaterial)
	rawMaterialRes, _ := repository.FindById(ctx, db, "id", helpers.IntToString(rawMaterial.Id))
	return rawMaterialRes, nil
}

func (repository RawMaterialRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, rawMaterial *domain.RawMaterial) (bool, error) {
	results := db.Where("id = ?", rawMaterial.Id).Delete(&rawMaterial)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|rawMaterial tidak ditemukan")
	}
	return true, nil
}

func (repository RawMaterialRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (rawMaterialRes domain.RawMaterial, err error) {
	results := db.Where(key+" = ?", value).First(&rawMaterialRes)
	if results.RowsAffected < 1 {
		return rawMaterialRes, errors.New("NOT_FOUND|rawMaterial tidak ditemukan")
	}
	return rawMaterialRes, nil
}

func (repository RawMaterialRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (rawMaterialRes []web.RawMaterialGet, err error) {
	qry := db.Table("raw_materials").
		Select(`
		raw_materials.id, raw_materials.name, raw_materials.price, raw_materials.stock,
		suppliers.id supplier_id, suppliers.name supplier_name,
		units.id unit_id, units.name unit_name,
		units.id smallest_unit_id, units.name smallest_unit_name,
		stores.id store_id, stores.name store_name
	`).
		Joins(`
		left join suppliers on suppliers.id = raw_materials.supplier_id
		left join units on units.id = raw_materials.unit_id
		left join stores on stores.id = raw_materials.store_id
	`)
	for k, v := range ctx.QueryParams() {
		fmt.Println(v)
		switch val := k; val {
		case "name":
			k = "raw_materials.name"
		case "supplier_name":
			k = "suppliers.name"
		case "store_name":
			k = "stores.name"
		default:
			fmt.Printf("%s.\n", k)
		}

		if k == "active" {
			qry = qry.Where(k+" = ?", v[0])
		} else if v[0] != "" && k != "id" {
			qry = qry.Where(k+" LIKE ?", "%"+v[0]+"%")
		}
	}
	qry.Order("id desc")
	qry.Find(&rawMaterialRes)

	return rawMaterialRes, nil
}

func (repository RawMaterialRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.RawMaterialDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("raw_materials").
		Select(`
		raw_materials.id, raw_materials.name, raw_materials.price, raw_materials.stock,
		suppliers.id supplier_id, suppliers.name supplier_name,
		units.id unit_id, units.name unit_name,
		units.id smallest_unit_id, units.name smallest_unit_name,
		stores.id store_id, stores.name store_name,
		divisions.name division_name
	`).
		Joins(`
		left join suppliers on suppliers.id = raw_materials.supplier_id
		left join units on units.id = raw_materials.unit_id
		left join stores on stores.id = raw_materials.store_id
		left join divisions on divisions.id = raw_materials.division_id
	`)
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(raw_materials.id = ? OR raw_materials.name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}
