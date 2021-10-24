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
	ReceivableDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (domain.ReceivableDetail, error)
		Update(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (domain.ReceivableDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.ReceivableDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.ReceivableDetail, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.ReceivableDetailDatatable, int64, int64, error)
	}

	ReceivableDetailRepositoryImpl struct {

	}
)

func NewReceivableDetailRepository() ReceivableDetailRepository {
	return &ReceivableDetailRepositoryImpl{}
}

func (repository *ReceivableDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (domain.ReceivableDetail, error) {
	receivable := domain.Receivable{}
	db.Where("id = ?", receivableDetail.ReceivableId).First(&receivable)
	if receivable.Receivables < 1 {
		return domain.ReceivableDetail{}, errors.New("TOTAL_BIGGER_THAN_DEBT|Piutang anda telah lunas")
	}
	if receivableDetail.Total > receivable.Receivables {
		return domain.ReceivableDetail{}, errors.New("TOTAL_BIGGER_THAN_DEBT|Total yang dibayarkan lebih besar daripada sisa piutang")
	}
	db.Create(&receivableDetail)
	db.Model(&domain.Receivable{}).Where("model = ? AND model_id = ?", receivable.Model, receivable.ModelId).Update("receivables", gorm.Expr("receivables - ?", receivableDetail.Total))
	receivableDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(receivableDetail.Id))
	return receivableDetailRes, nil
}

func (repository *ReceivableDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (domain.ReceivableDetail, error) {
	db.Where("id = ?", receivableDetail.Id).Updates(&receivableDetail)
	receivableDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(receivableDetail.Id))
	return receivableDetailRes, nil
}

func (repository *ReceivableDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, receivableDetail *domain.ReceivableDetail) (bool, error) {
	results := db.Where("id = ?", receivableDetail.Id).Delete(&receivableDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|receivableDetail tidak ditemukan")
	}
	return true, nil
}

func (repository *ReceivableDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (receivableDetailRes domain.ReceivableDetail, err error) {
	results := db.Where(key+" = ?", value).First(&receivableDetailRes)
	if results.RowsAffected < 1 {
		return receivableDetailRes, errors.New("NOT_FOUND|receivableDetail tidak ditemukan")
	}
	return receivableDetailRes, nil
}

func (repository *ReceivableDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (receivableDetailRes []domain.ReceivableDetail, err error) {
	db.Find(&receivableDetailRes)
	return receivableDetailRes, nil
}

func (repository *ReceivableDetailRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.ReceivableDetailDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("receivable_details").
		Select(`
		receivable_details.*,
		payment_methods.name payment_method
	`).
		Joins(`
		left join payment_methods on payment_methods.id = receivable_details.payment_method_id
	`)
	qry.Count(&totalData)
	if search != "" {

	}

	if ctx.QueryParam("receivable_id") != "" {
		qry.Where("receivable_id = ?", ctx.QueryParam("receivable_id"))
	}

	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("receivable_details.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}