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
	CashRegisterRepository interface{
		Create(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (domain.CashRegister, error)
		Update(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (domain.CashRegister, error)
		Delete(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.CashRegister, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.CashRegister, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.CashRegisterDatatable, int64, int64, error)
	}

	CashRegisterRepositoryImpl struct {

	}
)

func NewCashRegisterRepository() CashRegisterRepository {
	return &CashRegisterRepositoryImpl{}
}

func (repository CashRegisterRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (domain.CashRegister, error) {
	db.Create(&cashRegister)
	cashRegisterRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(cashRegister.Id))
	return cashRegisterRes, nil
}

func (repository CashRegisterRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (domain.CashRegister, error) {
	db.Where("id = ?", cashRegister.Id).Updates(&cashRegister)
	cashRegisterRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(cashRegister.Id))
	return cashRegisterRes, nil
}

func (repository CashRegisterRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, cashRegister *domain.CashRegister) (bool, error) {
	results := db.Where("id = ?", cashRegister.Id).Delete(&cashRegister)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|cashRegister tidak ditemukan")
	}
	return true, nil
}

func (repository CashRegisterRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (cashRegisterRes domain.CashRegister, err error) {
	results := db.Where(key+" = ?", value).First(&cashRegisterRes)
	if results.RowsAffected < 1 {
		return cashRegisterRes, errors.New("NOT_FOUND|cashRegister tidak ditemukan")
	}
	return cashRegisterRes, nil
}

func (repository CashRegisterRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (cashRegisterRes []domain.CashRegister, err error) {
	qry := db.Table("cash_registers").Select("cash_registers.*")
	for k, v := range ctx.QueryParams() {
		if v[0] != "" && k != "id" {
			qry = qry.Where(k+" = ?", v[0])
		}
	}
	qry.Scan(&cashRegisterRes)
	return cashRegisterRes, nil
}

func (repository CashRegisterRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.CashRegisterDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("cash_registers").
		Select(` stores.id store_id, stores.name store_name, 
			cash_registers.id, cash_registers.cash_in_hand, cash_registers.status, cash_registers.created_by, cash_registers.number, cash_registers.closed_at
		`).
		Joins("left join stores on stores.id = cash_registers.store_id")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(cash_registers.id = ? OR stores.name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("cash_registers.id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}