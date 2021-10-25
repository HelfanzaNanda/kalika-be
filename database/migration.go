package database

import (
	"kalika-be/models/domain"
	"sort"
)

func Migrate() {
	hasNewMigration := false
	setting := domain.Setting{Key: "db.migration.version"}
	db.AutoMigrate(&setting)
	db.Where(domain.Setting{Key: setting.Key}).FirstOrCreate(&setting)

	index := make([]string, 0)
	for i := range migration {
		index = append(index, i)
	}
	sort.Strings(index)
	for _, i := range index {
		if setting.Value == "" || setting.Value < i {
			migration[i]()
			setting.Value = i
			hasNewMigration = true
		}
	}
	if hasNewMigration {
		db.Where(domain.Setting{Key: setting.Key}).Assign(setting).FirstOrCreate(&setting)
	}
}

var migration = map[string]func(){
	"0001": func() { db.AutoMigrate(&domain.CakeType{}) },
	"0002": func() { db.AutoMigrate(&domain.CakeVariant{}) },
	"0003": func() { db.AutoMigrate(&domain.CashRegister{}) },
	"0004": func() { db.AutoMigrate(&domain.Category{}) },
	"0005": func() { db.AutoMigrate(&domain.CustomOrder{}) },
	"0006": func() { db.AutoMigrate(&domain.Customer{}) },
	"0007": func() { db.AutoMigrate(&domain.Debt{}) },
	"0008": func() { db.AutoMigrate(&domain.DebtDetail{}) },
	"0009": func() { db.AutoMigrate(&domain.Division{}) },
	"0010": func() { db.AutoMigrate(&domain.Expense{}) },
	"0011": func() { db.AutoMigrate(&domain.ExpenseCategory{}) },
	"0012": func() { db.AutoMigrate(&domain.ExpenseDetail{}) },
	"0013": func() { db.AutoMigrate(&domain.Payment{}) },
	"0014": func() { db.AutoMigrate(&domain.PaymentMethod{}) },
	"0015": func() { db.AutoMigrate(&domain.Permission{}) },
	"0016": func() { db.AutoMigrate(&domain.Product{}) },
	"0017": func() { db.AutoMigrate(&domain.PurchaseInvoice{}) },
	"0018": func() { db.AutoMigrate(&domain.PurchaseInvoiceDetail{}) },
	"0019": func() { db.AutoMigrate(&domain.PurchaseOrder{}) },
	"0020": func() { db.AutoMigrate(&domain.PurchaseOrderDelivery{}) },
	"0021": func() { db.AutoMigrate(&domain.PurchaseOrderDeliveryDetail{}) },
	"0022": func() { db.AutoMigrate(&domain.PurchaseOrderDetail{}) },
	"0023": func() { db.AutoMigrate(&domain.PurchaseReturn{}) },
	"0024": func() { db.AutoMigrate(&domain.PurchaseReturnDetail{}) },
	"0025": func() { db.AutoMigrate(&domain.RawMaterial{}) },
	"0026": func() { db.AutoMigrate(&domain.Receivable{}) },
	"0027": func() { db.AutoMigrate(&domain.ReceivableDetail{}) },
	"0028": func() { db.AutoMigrate(&domain.Role{}) },
	"0029": func() { db.AutoMigrate(&domain.RoleHasPermission{}) },
	"0030": func() { db.AutoMigrate(&domain.Sale{}) },
	"0031": func() { db.AutoMigrate(&domain.SalesConsignment{}) },
	"0032": func() { db.AutoMigrate(&domain.SalesConsignmentDetail{}) },
	"0033": func() { db.AutoMigrate(&domain.SalesDetail{}) },
	"0034": func() { db.AutoMigrate(&domain.SalesReturn{}) },
	"0035": func() { db.AutoMigrate(&domain.SalesReturnDetail{}) },
	"0036": func() { db.AutoMigrate(&domain.Seller{}) },
	"0037": func() { db.AutoMigrate(&domain.Setting{}) },
	"0038": func() { db.AutoMigrate(&domain.Store{}) },
	"0039": func() { db.AutoMigrate(&domain.StoreConsignment{}) },
	"0040": func() { db.AutoMigrate(&domain.Supplier{}) },
	"0041": func() { db.AutoMigrate(&domain.Unit{}) },
	"0042": func() { db.AutoMigrate(&domain.UnitConversion{}) },
	"0043": func() { db.AutoMigrate(&domain.User{}) },
	"0044": func() { db.AutoMigrate(&domain.Recipe{}) },
	"0045": func() { db.AutoMigrate(&domain.RecipeDetail{}) },
	"0046": func() { db.AutoMigrate(&domain.ProductLocation{}) },
	"0047": func() { db.AutoMigrate(&domain.StockOpname{}) },
	"0048": func() { db.AutoMigrate(&domain.StockOpnameDetail{}) },
	"0049": func() { db.AutoMigrate(&domain.ProductPrice{}) },
	"0050": func() { db.AutoMigrate(&domain.Debt{}) },
	"0051": func() { db.AutoMigrate(&domain.Receivable{}) },
	"0052": func() { db.AutoMigrate(&domain.GeneralSetting{}) },
	"0053": func() { db.AutoMigrate(&domain.GeneralSetting{}) },

}