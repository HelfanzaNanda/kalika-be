package repository

import (
	"errors"
	//"fmt"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	PurchaseReturnRepository interface{
		Create(ctx echo.Context, db *gorm.DB, purchaseReturn *web.PurchaseReturnPost) (domain.PurchaseReturn, error)
		Update(ctx echo.Context, db *gorm.DB, purchaseReturn *web.PurchaseReturnPost) (domain.PurchaseReturn, error)
		Delete(ctx echo.Context, db *gorm.DB, purchaseReturn *domain.PurchaseReturn) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.PurchaseReturn, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.PurchaseReturn, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.PurchaseReturnDatatable, int64, int64, error)
		ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) ([]web.PurchaseReturnDatatable, int64, int64, error)
		FindByCreatedAt(ctx echo.Context, db *gorm.DB, dateRange *web.DateRange) ([]web.PurchaseReturnGet, error)
	}

	PurchaseReturnRepositoryImpl struct {

	}
)

func NewPurchaseReturnRepository() PurchaseReturnRepository {
	return &PurchaseReturnRepositoryImpl{}
}

func (repository PurchaseReturnRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, purchaseReturn *web.PurchaseReturnPost) (domain.PurchaseReturn, error) {
	layoutFormat := "2006-01-02"
	date, _ := time.Parse(layoutFormat, purchaseReturn.Date)
	model := domain.PurchaseReturn{}
	model.Date = date
	model.Number = "PR"+helpers.IntToString(int(time.Now().Unix()))
	model.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	db.Create(&model)

	purchaseReturnRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(model.Id))
	return purchaseReturnRes, nil

}

func (repository PurchaseReturnRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, purchaseReturn *web.PurchaseReturnPost) (domain.PurchaseReturn, error) {
	layoutFormat := "2006-01-02"
	date, _ := time.Parse(layoutFormat, purchaseReturn.Date)
	model := domain.PurchaseReturn{}
	model.Date = date
	model.Number = "PR"+helpers.IntToString(int(time.Now().Unix()))
	model.CreatedBy = helpers.StringToInt(ctx.Get("userInfo").(map[string]interface{})["id"].(string))
	db.Where("id = ?", purchaseReturn.Id).Updates(&model)

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

func (repository PurchaseReturnRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.PurchaseReturnDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("purchase_returns")
	qry.Select(`
		purchase_returns.*,
		users.name created_by_name
	`)
	qry.Joins("left join users on users.id = purchase_returns.created_by")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(purchase_returns.id = ? OR purchase_returns.date LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("purchase_returns.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}

func (repository PurchaseReturnRepositoryImpl) ReportDatatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string, filter map[string]string) (datatableRes []web.PurchaseReturnDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("purchase_returns")
	qry.Select(`
		purchase_returns.*,
		users.name created_by_name
	`)
	qry.Joins("left join users on users.id = purchase_returns.created_by")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(purchase_returns.id = ? OR purchase_returns.date LIKE ?)", search, "%"+search+"%")
	}
	if filter["start_date"] != "" && filter["end_date"] != "" {
		qry.Where("(purchase_returns.created_at > ? AND purchase_returns.created_at < ?)", filter["start_date"], filter["end_date"])
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("purchase_returns.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}

func (repository PurchaseReturnRepositoryImpl) FindByCreatedAt(ctx echo.Context, db *gorm.DB, dateRange *web.DateRange) (purchaseReturnRes []web.PurchaseReturnGet, err error) {
	qry := db.Table("purchase_returns")
	qry.Select(`
		purchase_returns.*,
		users.name created_by_name
	`)
	qry.Joins("join users on users.id = purchase_returns.created_by")
	if dateRange.StartDate != "" && dateRange.EndDate != ""{
		qry.Where("(purchase_returns.created_at > ? AND purchase_returns.created_at < ?)", dateRange.StartDate, dateRange.EndDate)
	}
	qry.Find(&purchaseReturnRes)
	return purchaseReturnRes, nil
}