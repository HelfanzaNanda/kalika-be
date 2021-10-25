package repository

import (
	"kalika-be/models/domain"
	"kalika-be/models/web"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	ReportRepository interface{
		ProfitLoss(ctx echo.Context, db *gorm.DB) (web.ProfitLossReport, error)
		ReceivableLedger(ctx echo.Context, db *gorm.DB) ([]web.ReportLedgerReceivable, error)
		DebtLedger(ctx echo.Context, db *gorm.DB) ([]web.ReportLedgerDebt, error)
	}

	ReportRepositoryImpl struct {

	}
)

func NewReportRepository() ReportRepository {
	return &ReportRepositoryImpl{}
}

func (repository *ReportRepositoryImpl) ProfitLoss(ctx echo.Context, db *gorm.DB) (res web.ProfitLossReport, err error) {
	var cogsSales float64
	var cogsSalesConsignment float64
	db.Model(&domain.SalesDetail{}).Select("sum(recipes.total_cogs) as total_cogs").Joins("JOIN recipes ON sales_details.product_id = recipes.product_id").First(&cogsSales)
	db.Model(&domain.SalesConsignmentDetail{}).Select("sum(recipes.total_cogs) as total_cogs").Joins("JOIN recipes ON sales_details.product_id = recipes.product_id").First(&cogsSalesConsignment)
	res.TotalCogs = cogsSales + cogsSalesConsignment
	db.Model(&domain.Sale{}).Select("sum(customer_pay) as sales").First(&res)
	db.Model(&domain.CustomOrder{}).Select("sum(price) as custom_order").First(&res)
	db.Model(&domain.SalesConsignment{}).Select("sum(total) as sales_consignment").First(&res)
	db.Model(&domain.ReceivableDetail{}).Select("sum(total) as receivable_payment").First(&res)
	db.Model(&domain.DebtDetail{}).Select("sum(total) as debt_payment").First(&res)
	db.Model(&domain.SalesReturn{}).Select("sum(total) as sales_return").First(&res)
	db.Model(&domain.Expense{}).Select("sum(total) as total_cost").First(&res)
	db.Model(&domain.ExpenseDetail{}).Select("expense_categories.name name, amount total").Joins("JOIN expense_categories ON expense_categories.id = expense_details.expense_category_id").Find(&res.Cost)
	return res, nil
}

func (repository *ReportRepositoryImpl) ReceivableLedger(ctx echo.Context, db *gorm.DB) (res []web.ReportLedgerReceivable, err error) {
	data := []web.ReportLedgerReceivable{}
	db.Model(&domain.Receivable{}).
		Select("sales_consignments.date as date, sales_consignments.number as number, store_consignments.store_name as customer, receivables.total as debit, (SELECT COALESCE(sum(receivable_details.total), 0) FROM receivable_details WHERE receivable_details.receivable_id = receivables.id) AS credit").
		Where("receivables.model = ?", "SalesConsignment").
		Joins("JOIN sales_consignments ON sales_consignments.id = receivables.model_id").
		Joins("JOIN store_consignments ON store_consignments.id = sales_consignments.store_consignment_id").
		Order("receivables.date asc").
		Find(&data)

	tempBalance := 0.0

	for key, val := range data {
		data[key].Balance = tempBalance + val.Debit - val.Credit
		tempBalance = data[key].Balance
	}

	return data, nil
}

func (repository *ReportRepositoryImpl) DebtLedger(ctx echo.Context, db *gorm.DB) (res []web.ReportLedgerDebt, err error) {
	data := []web.ReportLedgerDebt{}
	db.Model(&domain.Debt{}).
		Select("purchase_orders.date as date, purchase_orders.number as number, suppliers.name as supplier, debts.total as credit, (SELECT COALESCE(sum(debt_details.total), 0) FROM debt_details WHERE debt_details.debt_id = debts.id) AS debit").
		Where("debts.model = ?", "PurchaseOrder").
		Joins("JOIN purchase_orders ON purchase_orders.id = debts.model_id").
		Joins("JOIN suppliers ON suppliers.id = debts.supplier_id").
		Order("debts.date asc").
		Find(&data)

	tempBalance := 0.0

	for key, val := range data {
		data[key].Balance = tempBalance + val.Credit - val.Debit
		tempBalance = data[key].Balance
	}

	return data, nil
}