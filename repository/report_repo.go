package repository

import (
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"sort"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	ReportRepository interface{
		ProfitLoss(ctx echo.Context, db *gorm.DB) (web.ProfitLossReport, error)
		ReceivableLedger(ctx echo.Context, db *gorm.DB) ([]web.ReportLedgerReceivable, error)
		DebtLedger(ctx echo.Context, db *gorm.DB) ([]web.ReportLedgerDebt, error)
		CashBankLedger(ctx echo.Context, db *gorm.DB) ([]web.ReportLedgerCashBank, error)
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

func (repository *ReportRepositoryImpl) CashBankLedger(ctx echo.Context, db *gorm.DB) (res []web.ReportLedgerCashBank, err error) {
	data := []web.ReportLedgerCashBank{}
	// INCOME
	dataSales := []web.ReportLedgerCashBank{}
	dataSalesConsignments := []web.ReportLedgerCashBank{}
	dataCustomOrders := []web.ReportLedgerCashBank{}
	dataReceivablePayments := []web.ReportLedgerCashBank{}
	// EXPENSE
	dataPurchaseOrders := []web.ReportLedgerCashBank{}
	dataExpenses := []web.ReportLedgerCashBank{}
	dataDebtPayments := []web.ReportLedgerCashBank{}

	db.Model(&domain.Payment{}).
		Select("payments.date as date, sales.number as number, payments.total as debit, 0 AS credit, payment_methods.name as payment_method, 'Penjualan' as type").
		Where("payments.model = ?", "Sales").
		Where("payment_methods.category = ?", "cash").
		Joins("JOIN sales ON sales.id = payments.model_id").
		Joins("JOIN payment_methods ON payment_methods.id = payments.payment_method_id").
		Find(&dataSales)

	db.Model(&domain.Payment{}).
		Select("payments.date as date, sales_consignments.number as number, payments.total as debit, 0 AS credit, payment_methods.name as payment_method, 'Penjualan Konsinyasi' as type").
		Where("payments.model = ?", "SalesConsignment").
		Where("payment_methods.category = ?", "cash").
		Joins("JOIN sales_consignments ON sales_consignments.id = payments.model_id").
		Joins("JOIN payment_methods ON payment_methods.id = payments.payment_method_id").
		Find(&dataSalesConsignments)

	db.Model(&domain.CustomOrder{}).
		Select("custom_orders.created_at as date, custom_orders.number as number, custom_orders.total as debit, 0 AS credit, payment_methods.name as payment_method, 'Penjualan Pesanan' as type").
		Where("payment_methods.category = ?", "cash").
		Joins("JOIN payment_methods ON payment_methods.id = custom_orders.payment_method_id").
		Find(&dataCustomOrders)

	db.Model(&domain.ReceivableDetail{}).
		Select("receivable_details.date_pay as date, receivables.note as number, receivables.total as debit, 0 AS credit, payment_methods.name as payment_method, 'Pembayaran Piutang' as type").
		Where("payment_methods.category = ?", "cash").
		Joins("JOIN receivables ON receivables.id = receivable_details.receivable_id").
		Joins("JOIN payment_methods ON payment_methods.id = receivable_details.payment_method_id").
		Find(&dataReceivablePayments)

	db.Model(&domain.Payment{}).
		Select("payments.date as date, purchase_orders.number as number, payments.total as credit, 0 AS debit, payment_methods.name as payment_method, 'Order Pembelian' as type").
		Where("payments.model = ?", "PurchaseOrder").
		Where("payment_methods.category = ?", "cash").
		Joins("JOIN purchase_orders ON purchase_orders.id = payments.model_id").
		Joins("JOIN payment_methods ON payment_methods.id = payments.payment_method_id").
		Find(&dataPurchaseOrders)

	db.Model(&domain.Expense{}).
		Select("expenses.date as date, expenses.number as number, expenses.total as credit, 0 AS debit, payment_methods.name as payment_method, 'Biaya' as type").
		Joins("JOIN payment_methods ON payment_methods.id = expenses.payment_method_id").
		Find(&dataExpenses)

	db.Model(&domain.DebtDetail{}).
		Select("debt_details.date_pay as date, debts.note as number, debts.total as credit, 0 AS debit, payment_methods.name as payment_method, 'Pembayaran Hutang' as type").
		Joins("JOIN debts ON debts.id = debt_details.debt_id").
		Joins("JOIN payment_methods ON payment_methods.id = debt_details.payment_method_id").
		Find(&dataDebtPayments)

	data = append(data, dataSales...)
	data = append(data, dataSalesConsignments...)
	data = append(data, dataCustomOrders...)
	data = append(data, dataReceivablePayments...)
	data = append(data, dataPurchaseOrders...)
	data = append(data, dataExpenses...)
	data = append(data, dataDebtPayments...)

	sort.Slice(data, func(i, j int) bool {
		return data[i].Date.Before(data[j].Date)
	})

	tempBalance := 0.0

	for key, val := range data {
		data[key].Balance = tempBalance + val.Debit - val.Credit
		tempBalance = data[key].Balance
	}

	return data, nil
}