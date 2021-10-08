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
	PurchaseReturnRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (domain.PurchaseReturn, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (domain.PurchaseReturn, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseReturn, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PurchaseReturn, error)
		ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.PurchaseReturnDatatable, int64, int64, error)
	}

	PurchaseReturnRepositoryImpl struct {

	}
)

func NewPurchaseReturnRepository() PurchaseReturnRepository {
	return &PurchaseReturnRepositoryImpl{}
}

func (repository PurchaseReturnRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (domain.PurchaseReturn, error) {
	db.Create(&purchaseReturn)
	purchaseReturnRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseReturn.Id))
	return purchaseReturnRes, nil
}

func (repository PurchaseReturnRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (domain.PurchaseReturn, error) {
	db.Where("id = ?", purchaseReturn.Id).Updates(&purchaseReturn)
	purchaseReturnRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(purchaseReturn.Id))
	return purchaseReturnRes, nil
}

func (repository PurchaseReturnRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (bool, error) {
	results := db.Where("id = ?", purchaseReturn.Id).Delete(&purchaseReturn)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|purchaseReturn tidak ditemukan")
	}
	return true, nil
}

func (repository PurchaseReturnRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (purchaseReturnRes domain.PurchaseReturn, err error) {
	results := db.Where(key+" = ?", value).First(&purchaseReturnRes)
	if results.RowsAffected < 1 {
		return purchaseReturnRes, errors.New("NOT_FOUND|purchaseReturn tidak ditemukan")
	}
	return purchaseReturnRes, nil
}

func (repository PurchaseReturnRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (purchaseReturnRes []domain.PurchaseReturn, err error) {
	db.Find(&purchaseReturnRes)
	return purchaseReturnRes, nil
}

func (repository PurchaseReturnRepositoryImpl) ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.PurchaseReturnDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("purchase_returns")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(id = ? OR date LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}