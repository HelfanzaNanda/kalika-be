package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"kalika-be/controllers"
)

func Routes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	api := e.Group("/api")
	userController := controllers.NewUserController()
	divisionController := controllers.NewDivisionController()
	categoryController := controllers.NewCategoryController()
	cakeVariantController := controllers.NewCakeVariantController()
	cakeTypeController := controllers.NewCakeTypeController()
	storeController := controllers.NewStoreController()
	supplierController := controllers.NewSupplierController()
	customerController := controllers.NewCustomerController()
	storeConsignmentController := controllers.NewStoreConsignmentController()
	paymentMethodController := controllers.NewPaymentMethodController()
	sellerController := controllers.NewSellerController()
	expenseCategoryController := controllers.NewExpenseCategoryController()
	rawMaterialController := controllers.NewRawMaterialController()
	customOrderController := controllers.NewCustomOrderController()
	unitController := controllers.NewUnitController()
	unitConversionController := controllers.NewUnitConversionController()
	debtController := controllers.NewDebtController()
	debtDetailController := controllers.NewDebtDetailController()
	receivableController := controllers.NewReceivableController()
	receivableDetailController := controllers.NewReceivableDetailController()
	expenseController := controllers.NewExpenseController()
	expenseDetailController := controllers.NewExpenseDetailController()
	purchaseReturnController := controllers.NewPurchaseReturnController()
	purchaseReturnDetailController := controllers.NewPurchaseReturnDetailController()
	purchaseOrderController := controllers.NewPurchaseOrderController()
	purchaseOrderDetailController := controllers.NewPurchaseOrderDetailController()
	purchaseInvoiceController := controllers.NewPurchaseInvoiceController()
	purchaseInvoiceDetailController := controllers.NewPurchaseInvoiceDetailController()
	purchaseOrderDeliveryController := controllers.NewPurchaseOrderDeliveryController()
	purchaseOrderDeliveryDetailController := controllers.NewPurchaseOrderDeliveryDetailController()
	saleController := controllers.NewSaleController()
	salesDetailController := controllers.NewSalesDetailController()
	salesReturnController := controllers.NewSalesReturnController()
	salesReturnDetailController := controllers.NewSalesReturnDetailController()
	salesConsignmentController := controllers.NewSalesConsignmentController()
	salesConsignmentDetailController := controllers.NewSalesConsignmentDetailController()
	cashRegisterController := controllers.NewCashRegisterController()
	paymentController := controllers.NewPaymentController()


	api.GET("/users", userController.FindAll)
	api.GET("/users/:id", userController.FindById)
	api.POST("/users", userController.Create)
	api.PUT("/users/:id", userController.Update)
	api.DELETE("/users/:id", userController.Delete)

	api.GET("/divisions", divisionController.FindAll)
	api.GET("/divisions/:id", divisionController.FindById)
	api.POST("/divisions", divisionController.Create)
	api.PUT("/divisions/:id", divisionController.Update)
	api.DELETE("/divisions/:id", divisionController.Delete)
	
	api.GET("/categories", categoryController.FindAll)
	api.GET("/categories/:id", categoryController.FindById)
	api.POST("/categories", categoryController.Create)
	api.PUT("/categories/:id", categoryController.Update)
	api.DELETE("/categories/:id", categoryController.Delete)
	
	api.GET("/cake_variants", cakeVariantController.FindAll)
	api.GET("/cake_variants/:id", cakeVariantController.FindById)
	api.POST("/cake_variants", cakeVariantController.Create)
	api.PUT("/cake_variants/:id", cakeVariantController.Update)
	api.DELETE("/cake_variants/:id", cakeVariantController.Delete)
	
	api.GET("/cake_types", cakeTypeController.FindAll)
	api.GET("/cake_types/:id", cakeTypeController.FindById)
	api.POST("/cake_types", cakeTypeController.Create)
	api.PUT("/cake_types/:id", cakeTypeController.Update)
	api.DELETE("/cake_types/:id", cakeTypeController.Delete)
	
	api.GET("/stores", storeController.FindAll)
	api.GET("/stores/:id", storeController.FindById)
	api.POST("/stores", storeController.Create)
	api.PUT("/stores/:id", storeController.Update)
	api.DELETE("/stores/:id", storeController.Delete)
	
	api.GET("/suppliers", supplierController.FindAll)
	api.GET("/suppliers/:id", supplierController.FindById)
	api.POST("/suppliers", supplierController.Create)
	api.PUT("/suppliers/:id", supplierController.Update)
	api.DELETE("/suppliers/:id", supplierController.Delete)
	
	api.GET("/customers", customerController.FindAll)
	api.GET("/customers/:id", customerController.FindById)
	api.POST("/customers", customerController.Create)
	api.PUT("/customers/:id", customerController.Update)
	api.DELETE("/customers/:id", customerController.Delete)

	api.GET("/store_consignments", storeConsignmentController.FindAll)
	api.GET("/store_consignments/:id", storeConsignmentController.FindById)
	api.POST("/store_consignments", storeConsignmentController.Create)
	api.PUT("/store_consignments/:id", storeConsignmentController.Update)
	api.DELETE("/store_consignments/:id", storeConsignmentController.Delete)
	
	api.GET("/payment_methods", paymentMethodController.FindAll)
	api.GET("/payment_methods/:id", paymentMethodController.FindById)
	api.POST("/payment_methods", paymentMethodController.Create)
	api.PUT("/payment_methods/:id", paymentMethodController.Update)
	api.DELETE("/payment_methods/:id", paymentMethodController.Delete)
	
	api.GET("/sellers", sellerController.FindAll)
	api.GET("/sellers/:id", sellerController.FindById)
	api.POST("/sellers", sellerController.Create)
	api.PUT("/sellers/:id", sellerController.Update)
	api.DELETE("/sellers/:id", sellerController.Delete)
	
	api.GET("/expense_categories", expenseCategoryController.FindAll)
	api.GET("/expense_categories/:id", expenseCategoryController.FindById)
	api.POST("/expense_categories", expenseCategoryController.Create)
	api.PUT("/expense_categories/:id", expenseCategoryController.Update)
	api.DELETE("/expense_categories/:id", expenseCategoryController.Delete)
	
	api.GET("/raw_materials", rawMaterialController.FindAll)
	api.GET("/raw_materials/:id", rawMaterialController.FindById)
	api.POST("/raw_materials", rawMaterialController.Create)
	api.PUT("/raw_materials/:id", rawMaterialController.Update)
	api.DELETE("/raw_materials/:id", rawMaterialController.Delete)
	
	api.GET("/units", unitController.FindAll)
	api.GET("/units/:id", unitController.FindById)
	api.POST("/units", unitController.Create)
	api.PUT("/units/:id", unitController.Update)
	api.DELETE("/units/:id", unitController.Delete)
	
	api.GET("/unit_conversions", unitConversionController.FindAll)
	api.GET("/unit_conversions/:id", unitConversionController.FindById)
	api.POST("/unit_conversions", unitConversionController.Create)
	api.PUT("/unit_conversions/:id", unitConversionController.Update)
	api.DELETE("/unit_conversions/:id", unitConversionController.Delete)
	
	api.GET("/custom_orders", customOrderController.FindAll)
	api.GET("/custom_orders/:id", customOrderController.FindById)
	api.POST("/custom_orders", customOrderController.Create)
	api.PUT("/custom_orders/:id", customOrderController.Update)
	api.DELETE("/custom_orders/:id", customOrderController.Delete)
	
	api.GET("/debts", debtController.FindAll)
	api.GET("/debts/:id", debtController.FindById)
	api.POST("/debts", debtController.Create)
	api.PUT("/debts/:id", debtController.Update)
	api.DELETE("/debts/:id", debtController.Delete)
	
	api.GET("/debt_details", debtDetailController.FindAll)
	api.GET("/debt_details/:id", debtDetailController.FindById)
	api.POST("/debt_details", debtDetailController.Create)
	api.PUT("/debt_details/:id", debtDetailController.Update)
	api.DELETE("/debt_details/:id", debtDetailController.Delete)
	
	api.GET("/receivables", receivableController.FindAll)
	api.GET("/receivables/:id", receivableController.FindById)
	api.POST("/receivables", receivableController.Create)
	api.PUT("/receivables/:id", receivableController.Update)
	api.DELETE("/receivables/:id", receivableController.Delete)
	
	api.GET("/receivable_details", receivableDetailController.FindAll)
	api.GET("/receivable_details/:id", receivableDetailController.FindById)
	api.POST("/receivable_details", receivableDetailController.Create)
	api.PUT("/receivable_details/:id", receivableDetailController.Update)
	api.DELETE("/receivable_details/:id", receivableDetailController.Delete)
	
	api.GET("/expenses", expenseController.FindAll)
	api.GET("/expenses/:id", expenseController.FindById)
	api.POST("/expenses", expenseController.Create)
	api.PUT("/expenses/:id", expenseController.Update)
	api.DELETE("/expenses/:id", expenseController.Delete)
	
	api.GET("/expense_details", expenseDetailController.FindAll)
	api.GET("/expense_details/:id", expenseDetailController.FindById)
	api.POST("/expense_details", expenseDetailController.Create)
	api.PUT("/expense_details/:id", expenseDetailController.Update)
	api.DELETE("/expense_details/:id", expenseDetailController.Delete)
	
	api.GET("/purchase_returns", purchaseReturnController.FindAll)
	api.GET("/purchase_returns/:id", purchaseReturnController.FindById)
	api.POST("/purchase_returns", purchaseReturnController.Create)
	api.PUT("/purchase_returns/:id", purchaseReturnController.Update)
	api.DELETE("/purchase_returns/:id", purchaseReturnController.Delete)
	
	api.GET("/purchase_return_details", purchaseReturnDetailController.FindAll)
	api.GET("/purchase_return_details/:id", purchaseReturnDetailController.FindById)
	api.POST("/purchase_return_details", purchaseReturnDetailController.Create)
	api.PUT("/purchase_return_details/:id", purchaseReturnDetailController.Update)
	api.DELETE("/purchase_return_details/:id", purchaseReturnDetailController.Delete)
	
	api.GET("/purchase_orders", purchaseOrderController.FindAll)
	api.GET("/purchase_orders/:id", purchaseOrderController.FindById)
	api.POST("/purchase_orders", purchaseOrderController.Create)
	api.PUT("/purchase_orders/:id", purchaseOrderController.Update)
	api.DELETE("/purchase_orders/:id", purchaseOrderController.Delete)
	
	api.GET("/purchase_order_details", purchaseOrderDetailController.FindAll)
	api.GET("/purchase_order_details/:id", purchaseOrderDetailController.FindById)
	api.POST("/purchase_order_details", purchaseOrderDetailController.Create)
	api.PUT("/purchase_order_details/:id", purchaseOrderDetailController.Update)
	api.DELETE("/purchase_order_details/:id", purchaseOrderDetailController.Delete)
	
	api.GET("/purchase_invoices", purchaseInvoiceController.FindAll)
	api.GET("/purchase_invoices/:id", purchaseInvoiceController.FindById)
	api.POST("/purchase_invoices", purchaseInvoiceController.Create)
	api.PUT("/purchase_invoices/:id", purchaseInvoiceController.Update)
	api.DELETE("/purchase_invoices/:id", purchaseInvoiceController.Delete)
	
	api.GET("/purchase_invoice_details", purchaseInvoiceDetailController.FindAll)
	api.GET("/purchase_invoice_details/:id", purchaseInvoiceDetailController.FindById)
	api.POST("/purchase_invoice_details", purchaseInvoiceDetailController.Create)
	api.PUT("/purchase_invoice_details/:id", purchaseInvoiceDetailController.Update)
	api.DELETE("/purchase_invoice_details/:id", purchaseInvoiceDetailController.Delete)
	
	api.GET("/purchase_order_deliveries", purchaseOrderDeliveryController.FindAll)
	api.GET("/purchase_order_deliveries/:id", purchaseOrderDeliveryController.FindById)
	api.POST("/purchase_order_deliveries", purchaseOrderDeliveryController.Create)
	api.PUT("/purchase_order_deliveries/:id", purchaseOrderDeliveryController.Update)
	api.DELETE("/purchase_order_deliveries/:id", purchaseOrderDeliveryController.Delete)
	
	api.GET("/purchase_order_delivery_details", purchaseOrderDeliveryDetailController.FindAll)
	api.GET("/purchase_order_delivery_details/:id", purchaseOrderDeliveryDetailController.FindById)
	api.POST("/purchase_order_delivery_details", purchaseOrderDeliveryDetailController.Create)
	api.PUT("/purchase_order_delivery_details/:id", purchaseOrderDeliveryDetailController.Update)
	api.DELETE("/purchase_order_delivery_details/:id", purchaseOrderDeliveryDetailController.Delete)
	
	api.GET("/sales", saleController.FindAll)
	api.GET("/sales/:id", saleController.FindById)
	api.POST("/sales", saleController.Create)
	api.PUT("/sales/:id", saleController.Update)
	api.DELETE("/sales/:id", saleController.Delete)
	
	api.GET("/sales_details", salesDetailController.FindAll)
	api.GET("/sales_details/:id", salesDetailController.FindById)
	api.POST("/sales_details", salesDetailController.Create)
	api.PUT("/sales_details/:id", salesDetailController.Update)
	api.DELETE("/sales_details/:id", salesDetailController.Delete)
	
	api.GET("/sales_returns", salesReturnController.FindAll)
	api.GET("/sales_returns/:id", salesReturnController.FindById)
	api.POST("/sales_returns", salesReturnController.Create)
	api.PUT("/sales_returns/:id", salesReturnController.Update)
	api.DELETE("/sales_returns/:id", salesReturnController.Delete)
	
	api.GET("/sales_return_details", salesReturnDetailController.FindAll)
	api.GET("/sales_return_details/:id", salesReturnDetailController.FindById)
	api.POST("/sales_return_details", salesReturnDetailController.Create)
	api.PUT("/sales_return_details/:id", salesReturnDetailController.Update)
	api.DELETE("/sales_return_details/:id", salesReturnDetailController.Delete)
	
	api.GET("/sales_consignment", salesConsignmentController.FindAll)
	api.GET("/sales_consignment/:id", salesConsignmentController.FindById)
	api.POST("/sales_consignment", salesConsignmentController.Create)
	api.PUT("/sales_consignment/:id", salesConsignmentController.Update)
	api.DELETE("/sales_consignment/:id", salesConsignmentController.Delete)
	
	api.GET("/sales_consignment_details", salesConsignmentDetailController.FindAll)
	api.GET("/sales_consignment_details/:id", salesConsignmentDetailController.FindById)
	api.POST("/sales_consignment_details", salesConsignmentDetailController.Create)
	api.PUT("/sales_consignment_details/:id", salesConsignmentDetailController.Update)
	api.DELETE("/sales_consignment_details/:id", salesConsignmentDetailController.Delete)
	
	api.GET("/cash_registers", cashRegisterController.FindAll)
	api.GET("/cash_registers/:id", cashRegisterController.FindById)
	api.POST("/cash_registers", cashRegisterController.Create)
	api.PUT("/cash_registers/:id", cashRegisterController.Update)
	api.DELETE("/cash_registers/:id", cashRegisterController.Delete)
	
	api.GET("/payments", paymentController.FindAll)
	api.GET("/payments/:id", paymentController.FindById)
	api.POST("/payments", paymentController.Create)
	api.PUT("/payments/:id", paymentController.Update)
	api.DELETE("/payments/:id", paymentController.Delete)
	
	


	return e
}