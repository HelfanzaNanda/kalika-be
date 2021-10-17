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
	}

	ReportRepositoryImpl struct {

	}
)

func NewReportRepository() ReportRepository {
	return &ReportRepositoryImpl{}
}

func (repository ReportRepositoryImpl) ProfitLoss(ctx echo.Context, db *gorm.DB) (res web.ProfitLossReport, err error) {
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