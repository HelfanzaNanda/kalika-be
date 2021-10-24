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
	DebtDetailRepository interface{
		Create(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (domain.DebtDetail, error)
		Update(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (domain.DebtDetail, error)
		Delete(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.DebtDetail, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.DebtDetail, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.DebtDetailDatatable, int64, int64, error)
	}

	DebtDetailRepositoryImpl struct {

	}
)

func NewDebtDetailRepository() DebtDetailRepository {
	return &DebtDetailRepositoryImpl{}
}

func (repository *DebtDetailRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (domain.DebtDetail, error) {
	debt := domain.Debt{}
	db.Where("id = ?", debtDetail.DebtId).First(&debt)
	if debt.Debts < 1 {
		return domain.DebtDetail{}, errors.New("TOTAL_BIGGER_THAN_DEBT|Hutang anda telah lunas")
	}
	if debtDetail.Total > debt.Debts {
		return domain.DebtDetail{}, errors.New("TOTAL_BIGGER_THAN_DEBT|Total yang dibayarkan lebih besar daripada sisa hutang")
	}
	db.Create(&debtDetail)
	db.Model(&domain.Debt{}).Where("model = ? AND model_id = ?", debt.Model, debt.ModelId).Update("debts", gorm.Expr("debts - ?", debtDetail.Total))
	debtDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(debtDetail.Id))
	return debtDetailRes, nil
}

func (repository *DebtDetailRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (domain.DebtDetail, error) {
	db.Where("id = ?", debtDetail.Id).Updates(&debtDetail)
	debtDetailRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(debtDetail.Id))
	return debtDetailRes, nil
}

func (repository *DebtDetailRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, debtDetail *domain.DebtDetail) (bool, error) {
	results := db.Where("id = ?", debtDetail.Id).Delete(&debtDetail)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|debtDetail tidak ditemukan")
	}
	return true, nil
}

func (repository *DebtDetailRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (debtDetailRes domain.DebtDetail, err error) {
	results := db.Where(key+" = ?", value).First(&debtDetailRes)
	if results.RowsAffected < 1 {
		return debtDetailRes, errors.New("NOT_FOUND|debtDetail tidak ditemukan")
	}
	return debtDetailRes, nil
}

func (repository *DebtDetailRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (debtDetailRes []domain.DebtDetail, err error) {
	db.Find(&debtDetailRes)
	return debtDetailRes, nil
}

func (repository *DebtDetailRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.DebtDetailDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("debt_details").
		Select(`
		debt_details.*,
		payment_methods.name payment_method
	`).
		Joins(`
		left join payment_methods on payment_methods.id = debt_details.payment_method_id
	`)
	qry.Count(&totalData)
	if search != "" {

	}

	if ctx.QueryParam("debt_id") != "" {
		qry.Where("debt_id = ?", ctx.QueryParam("debt_id"))
	}

	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("debt_details.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}

