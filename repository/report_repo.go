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
		ProfitLoss(ctx echo.Context, db *gorm.DB, filter *web.ProfitLossFilter) (web.ProfitLossReport, error)
		ReceivableLedger(ctx echo.Context, db *gorm.DB, filter *web.DateRange) ([]web.ReportLedgerReceivable, error)
		DebtLedger(ctx echo.Context, db *gorm.DB, filter *web.DateRange) ([]web.ReportLedgerDebt, error)
		CashBankLedger(ctx echo.Context, db *gorm.DB, filter *web.DateRange) ([]web.ReportLedgerCashBank, error)
	}

	ReportRepositoryImpl struct {

	}
)

func NewReportRepository() ReportRepository {
	return &ReportRepositoryImpl{}
}

func (repository *ReportRepositoryImpl) ProfitLoss(ctx echo.Context, db *gorm.DB, filter * web.ProfitLossFilter) (res web.ProfitLossReport, err error) {
	var cogsSales float64 = 0
	var cogsSalesConsignment float64 = 0

	qrySales := db.Model(&domain.SalesDetail{})
	qrySales.Select("COALESCE(SUM(recipes.total_cogs), 0) as total_cogs")
	qrySales.Joins("JOIN recipes ON sales_details.product_id = recipes.product_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qrySales.Where("(DATE(sales_details.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qrySales.First(&cogsSales)

	qrySalesConsignment := db.Model(&domain.SalesConsignmentDetail{})
	qrySalesConsignment.Select("COALESCE(SUM(recipes.total_cogs), 0) as total_cogs")
	qrySalesConsignment.Joins("JOIN recipes ON sales_consignment_details.product_id = recipes.product_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qrySalesConsignment.Where("(DATE(sales_consignment_details.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qrySalesConsignment.First(&cogsSalesConsignment)
	res.TotalCogs = cogsSales + cogsSalesConsignment

	qrySale := db.Model(&domain.Sale{})
	qrySale.Select("sum(customer_pay) as sales")
	if filter.StartDate != "" && filter.EndDate != "" {
		qrySale.Where("(DATE(sales.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	if filter.StoreId != 0 {
		qrySale.Where("(sales.store_id = ?)", filter.StoreId)
	}
	if filter.CreatedBy != 0 {
		qrySale.Where("(sales.created_by = ?)", filter.CreatedBy)
	}
	qrySale.First(&res.Sales)

	qryCustomOrder := db.Model(&domain.CustomOrder{})
	qryCustomOrder.Select("sum(price) as custom_order")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryCustomOrder.Where("(DATE(custom_orders.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	if filter.StoreId != 0 {
		qrySale.Where("(sales.store_id = ?)", filter.StoreId)
	}
	if filter.CreatedBy != 0 {
		qrySale.Where("(sales.created_by = ?)", filter.CreatedBy)
	}
	qryCustomOrder.First(&res.CustomOrder)

	qryTotalSalesConsignment := db.Model(&domain.SalesConsignment{})
	qryTotalSalesConsignment.Select("sum(total) as sales_consignment")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryTotalSalesConsignment.Where("(DATE(sales_consignments.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryTotalSalesConsignment.First(&res.SalesConsignment)

	qryTotalReceivablePayment := db.Model(&domain.ReceivableDetail{})
	qryTotalReceivablePayment.Select("sum(total) as receivable_payment")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryTotalReceivablePayment.Where("(DATE(receivable_details.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryTotalReceivablePayment.First(&res.ReceivablePayment)

	qryTotalDebtPayment := db.Model(&domain.DebtDetail{})
	qryTotalDebtPayment.Select("sum(total) as debt_payment")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryTotalDebtPayment.Where("(DATE(debt_details.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryTotalDebtPayment.First(&res.DebtPayment)
	
	qryTotalSalesReturn := db.Model(&domain.SalesReturn{})
	qryTotalSalesReturn.Select("sum(total) as sales_return")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryTotalSalesReturn.Where("(DATE(sales_returns.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryTotalSalesReturn.First(&res.SalesReturn)
	
	qryTotalCost := db.Model(&domain.Expense{})
	qryTotalCost.Select("sum(total) as total_cost")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryTotalCost.Where("(DATE(expenses.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryTotalCost.First(&res.TotalCost)
	
	qryExpense := db.Model(&domain.ExpenseDetail{})
	qryExpense.Select("expense_categories.name name, amount total")
	qryExpense.Joins("JOIN expense_categories ON expense_categories.id = expense_details.expense_category_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryExpense.Where("(DATE(expense_details.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryExpense.Find(&res.Cost)
	
	return res, nil
}

func (repository *ReportRepositoryImpl) ReceivableLedger(ctx echo.Context, db *gorm.DB, filter * web.DateRange) (res []web.ReportLedgerReceivable, err error) {
	data := []web.ReportLedgerReceivable{}
	salesConsignmentData := []web.ReportLedgerReceivable{}
	returnSalesConsignmentData := []web.ReportLedgerReceivable{}
	qryReceivable := db.Model(&domain.Receivable{})
	qryReceivable.Select("sales_consignments.date as date, sales_consignments.number as number, store_consignments.store_name as customer, receivables.total as debit, (SELECT COALESCE(sum(receivable_details.total), 0) FROM receivable_details WHERE receivable_details.receivable_id = receivables.id AND payment_method_id > 0) AS credit")
	qryReceivable.Where("receivables.model = ?", "SalesConsignment")
	qryReceivable.Joins("JOIN sales_consignments ON sales_consignments.id = receivables.model_id")
	qryReceivable.Joins("JOIN store_consignments ON store_consignments.id = sales_consignments.store_consignment_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryReceivable.Where("(DATE(receivables.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryReceivable.Order("receivables.date asc")
	qryReceivable.Find(&salesConsignmentData)

	qrySalesReturn := db.Model(&domain.SalesReturn{})
	qrySalesReturn.Select("sales_returns.created_at as date, sales_returns.number as number, store_consignments.store_name as customer, 0 as debit, sales_returns.total AS credit")
	qrySalesReturn.Where("sales_returns.model = ?", "SalesConsignment")
	qrySalesReturn.Joins("JOIN sales_consignments ON sales_consignments.id = sales_returns.model_id")
	qrySalesReturn.Joins("JOIN store_consignments ON store_consignments.id = sales_returns.store_consignment_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qrySalesReturn.Where("(DATE(sales_returns.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qrySalesReturn.Order("sales_returns.created_at asc")
	qrySalesReturn.Find(&returnSalesConsignmentData)

	data = append(data, salesConsignmentData...)
	data = append(data, returnSalesConsignmentData...)

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

func (repository *ReportRepositoryImpl) DebtLedger(ctx echo.Context, db *gorm.DB, filter * web.DateRange) (res []web.ReportLedgerDebt, err error) {
	data := []web.ReportLedgerDebt{}
	qry := db.Model(&domain.Debt{})
	qry.Select("purchase_orders.date as date, purchase_orders.number as number, suppliers.name as supplier, debts.total as credit, (SELECT COALESCE(sum(debt_details.total), 0) FROM debt_details WHERE debt_details.debt_id = debts.id) AS debit")
	qry.Where("debts.model = ?", "PurchaseOrder")
	qry.Joins("JOIN purchase_orders ON purchase_orders.id = debts.model_id")
	qry.Joins("JOIN suppliers ON suppliers.id = debts.supplier_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qry.Where("(DATE(debts.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qry.Order("debts.date asc")
	qry.Find(&data)

	tempBalance := 0.0

	for key, val := range data {
		data[key].Balance = tempBalance + val.Credit - val.Debit
		tempBalance = data[key].Balance
	}

	return data, nil
}

func (repository *ReportRepositoryImpl) CashBankLedger(ctx echo.Context, db *gorm.DB, filter * web.DateRange) (res []web.ReportLedgerCashBank, err error) {
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

	qrySales := db.Model(&domain.Payment{})
	qrySales.Select("payments.date as date, sales.number as number, payments.total as debit, 0 AS credit, payment_methods.name as payment_method, 'Penjualan' as type")
	qrySales.Where("payments.model = ?", "Sales")
	qrySales.Where("payment_methods.category = ?", "cash")
	qrySales.Joins("JOIN sales ON sales.id = payments.model_id")
	qrySales.Joins("JOIN payment_methods ON payment_methods.id = payments.payment_method_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qrySales.Where("(DATE(payments.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qrySales.Find(&dataSales)

	qrySalesConsignment := db.Model(&domain.Payment{})
	qrySalesConsignment.Select("payments.date as date, sales_consignments.number as number, payments.total as debit, 0 AS credit, payment_methods.name as payment_method, 'Penjualan Konsinyasi' as type")
	qrySalesConsignment.Where("payments.model = ?", "SalesConsignment")
	qrySalesConsignment.Where("payment_methods.category = ?", "cash")
	qrySalesConsignment.Joins("JOIN sales_consignments ON sales_consignments.id = payments.model_id")
	qrySalesConsignment.Joins("JOIN payment_methods ON payment_methods.id = payments.payment_method_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qrySalesConsignment.Where("(DATE(payments.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qrySalesConsignment.Find(&dataSalesConsignments)

	qryCustomOrder := db.Model(&domain.CustomOrder{})
	qryCustomOrder.Select("custom_orders.created_at as date, custom_orders.number as number, custom_orders.total as debit, 0 AS credit, payment_methods.name as payment_method, 'Penjualan Pesanan' as type")
	qryCustomOrder.Where("payment_methods.category = ?", "cash")
	qryCustomOrder.Joins("JOIN payment_methods ON payment_methods.id = custom_orders.payment_method_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryCustomOrder.Where("(DATE(custom_orders.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryCustomOrder.Find(&dataCustomOrders)

	qryReceivable := db.Model(&domain.ReceivableDetail{})
	qryReceivable.Select("receivable_details.date_pay as date, receivables.note as number, receivables.total as debit, 0 AS credit, payment_methods.name as payment_method, 'Pembayaran Piutang' as type")
	qryReceivable.Where("payment_methods.category = ?", "cash")
	qryReceivable.Joins("JOIN receivables ON receivables.id = receivable_details.receivable_id")
	qryReceivable.Joins("JOIN payment_methods ON payment_methods.id = receivable_details.payment_method_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryReceivable.Where("(DATE(receivable_details.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryReceivable.Find(&dataReceivablePayments)

	qryPurchaseOrder := db.Model(&domain.Payment{})
	qryPurchaseOrder.Select("payments.date as date, purchase_orders.number as number, payments.total as credit, 0 AS debit, payment_methods.name as payment_method, 'Order Pembelian' as type")
	qryPurchaseOrder.Where("payments.model = ?", "PurchaseOrder")
	qryPurchaseOrder.Where("payment_methods.category = ?", "cash")
	qryPurchaseOrder.Joins("JOIN purchase_orders ON purchase_orders.id = payments.model_id")
	qryPurchaseOrder.Joins("JOIN payment_methods ON payment_methods.id = payments.payment_method_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryPurchaseOrder.Where("(DATE(payments.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryPurchaseOrder.Find(&dataPurchaseOrders)

	qryExpense := db.Model(&domain.Expense{})
	qryExpense.Select("expenses.date as date, expenses.number as number, expenses.total as credit, 0 AS debit, payment_methods.name as payment_method, 'Biaya' as type")
	qryExpense.Joins("JOIN payment_methods ON payment_methods.id = expenses.payment_method_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryExpense.Where("(DATE(expenses.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryExpense.Find(&dataExpenses)

	qryDebtPayment := db.Model(&domain.DebtDetail{})
	qryDebtPayment.Select("debt_details.date_pay as date, debts.note as number, debts.total as credit, 0 AS debit, payment_methods.name as payment_method, 'Pembayaran Hutang' as type")
	qryDebtPayment.Joins("JOIN debts ON debts.id = debt_details.debt_id")
	qryDebtPayment.Joins("JOIN payment_methods ON payment_methods.id = debt_details.payment_method_id")
	if filter.StartDate != "" && filter.EndDate != "" {
		qryDebtPayment.Where("(DATE(debt_details.created_at) BETWEEN ? AND ?)", filter.StartDate, filter.EndDate)
	}
	qryDebtPayment.Find(&dataDebtPayments)

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