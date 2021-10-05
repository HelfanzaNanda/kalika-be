package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
	"kalika-be/controllers"
	"kalika-be/middlewares"
	"kalika-be/repository"
	"kalika-be/services"
)

func Routes(db *gorm.DB) *echo.Echo {
	//validate := validator.New()

	//USER THINGS
	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository, db)
	userController := controllers.NewUserController(userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	api := e.Group("/api")

	api.POST("/login", userController.Login)

	api.Use(middlewares.Auth)

	divisionRepository := repository.NewDivisionRepository()
	divisionService := services.NewDivisionService(divisionRepository, db)
	divisionController := controllers.NewDivisionController(divisionService)
	
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db)
	categoryController := controllers.NewCategoryController(categoryService)
	
	cakeVariantRepository := repository.NewCakeVariantRepository()
	cakeVariantService := services.NewCakeVariantService(cakeVariantRepository, db)
	cakeVariantController := controllers.NewCakeVariantController(cakeVariantService)
	
	cakeTypeRepository := repository.NewCakeTypeRepository()
	cakeTypeService := services.NewCakeTypeService(cakeTypeRepository, db)
	cakeTypeController := controllers.NewCakeTypeController(cakeTypeService)
	
	storeRepository := repository.NewStoreRepository()
	storeService := services.NewStoreService(storeRepository, db)
	storeController := controllers.NewStoreController(storeService)
	
	supplierRepository := repository.NewSupplierRepository()
	supplierService := services.NewSupplierService(supplierRepository, db)
	supplierController := controllers.NewSupplierController(supplierService)
	
	customerRepository := repository.NewCustomerRepository()
	customerService := services.NewCustomerService(customerRepository, db)
	customerController := controllers.NewCustomerController(customerService)
	
	storeConsignmentRepository := repository.NewStoreConsignmentRepository()
	storeConsignmentService := services.NewStoreConsignmentService(storeConsignmentRepository, db)
	storeConsignmentController := controllers.NewStoreConsignmentController(storeConsignmentService)
	
	paymentMethodRepository := repository.NewPaymentMethodRepository()
	paymentMethodService := services.NewPaymentMethodService(paymentMethodRepository, db)
	paymentMethodController := controllers.NewPaymentMethodController(paymentMethodService)
	
	sellerRepository := repository.NewSellerRepository()
	sellerService := services.NewSellerService(sellerRepository, db)
	sellerController := controllers.NewSellerController(sellerService)
	
	expenseCategoryRepository := repository.NewExpenseCategoryRepository()
	expenseCategoryService := services.NewExpenseCategoryService(expenseCategoryRepository, db)
	expenseCategoryController := controllers.NewExpenseCategoryController(expenseCategoryService)
	
	rawMaterialRepository := repository.NewRawMaterialRepository()
	rawMaterialService := services.NewRawMaterialService(rawMaterialRepository, db)
	rawMaterialController := controllers.NewRawMaterialController(rawMaterialService)
	
	customOrderRepository := repository.NewCustomOrderRepository()
	customOrderService := services.NewCustomOrderService(customOrderRepository, db)
	customOrderController := controllers.NewCustomOrderController(customOrderService)

	unitRepository := repository.NewUnitRepository()
	unitService := services.NewUnitService(unitRepository, db)
	unitController := controllers.NewUnitController(unitService)
	
	unitConversionRepository := repository.NewUnitConversionRepository()
	unitConversionService := services.NewUnitConversionService(unitConversionRepository, db)
	unitConversionController := controllers.NewUnitConversionController(unitConversionService)
	
	debtRepository := repository.NewDebtRepository()
	debtService := services.NewDebtService(debtRepository, db)
	debtController := controllers.NewDebtController(debtService)
	
	debtDetailRepository := repository.NewDebtDetailRepository()
	debtDetailService := services.NewDebtDetailService(debtDetailRepository, db)
	debtDetailController := controllers.NewDebtDetailController(debtDetailService)
	
	receivableRepository := repository.NewReceivableRepository()
	receivableService := services.NewReceivableService(receivableRepository, db)
	receivableController := controllers.NewReceivableController(receivableService)
	
	receivableDetailRepository := repository.NewReceivableRepository()
	receivableDetailService := services.NewReceivableService(receivableDetailRepository, db)
	receivableDetailController := controllers.NewReceivableController(receivableDetailService)
	
	expenseDetailRepository := repository.NewExpenseDetailRepository()
	expenseDetailService := services.NewExpenseDetailService(expenseDetailRepository, db)
	expenseDetailController := controllers.NewExpenseDetailController(expenseDetailService)
	
	expenseRepository := repository.NewExpenseRepository()
	expenseService := services.NewExpenseService(expenseRepository, expenseDetailRepository, db)
	expenseController := controllers.NewExpenseController(expenseService)
	
	
	purchaseReturnRepository := repository.NewPurchaseReturnRepository()
	purchaseReturnService := services.NewPurchaseReturnService(purchaseReturnRepository, db)
	purchaseReturnController := controllers.NewPurchaseReturnController(purchaseReturnService)
	
	purchaseReturnDetailRepository := repository.NewPurchaseReturnDetailRepository()
	purchaseReturnDetailService := services.NewPurchaseReturnDetailService(purchaseReturnDetailRepository, db)
	purchaseReturnDetailController := controllers.NewPurchaseReturnDetailController(purchaseReturnDetailService)
	
	purchaseOrderDetailRepository := repository.NewPurchaseOrderDetailRepository()
	purchaseOrderDetailService := services.NewPurchaseOrderDetailService(purchaseOrderDetailRepository, db)
	purchaseOrderDetailController := controllers.NewPurchaseOrderDetailController(purchaseOrderDetailService)
	
	purchaseInvoiceRepository := repository.NewPurchaseInvoiceRepository()
	purchaseInvoiceService := services.NewPurchaseInvoiceService(purchaseInvoiceRepository, db)
	purchaseInvoiceController := controllers.NewPurchaseInvoiceController(purchaseInvoiceService)
	
	purchaseInvoiceDetailRepository := repository.NewPurchaseInvoiceDetailRepository()
	purchaseInvoiceDetailService := services.NewPurchaseInvoiceDetailService(purchaseInvoiceDetailRepository, db)
	purchaseInvoiceDetailController := controllers.NewPurchaseInvoiceDetailController(purchaseInvoiceDetailService)
	
	purchaseOrderDeliveryRepository := repository.NewPurchaseOrderDeliveryRepository()
	purchaseOrderDeliveryService := services.NewPurchaseOrderDeliveryService(purchaseOrderDeliveryRepository, db)
	purchaseOrderDeliveryController := controllers.NewPurchaseOrderDeliveryController(purchaseOrderDeliveryService)

	purchaseOrderDeliveryDetailRepository := repository.NewPurchaseOrderDeliveryDetailRepository()
	purchaseOrderDeliveryDetailService := services.NewPurchaseOrderDeliveryDetailService(purchaseOrderDeliveryDetailRepository, db)
	purchaseOrderDeliveryDetailController := controllers.NewPurchaseOrderDeliveryDetailController(purchaseOrderDeliveryDetailService)

	salesDetailRepository := repository.NewSalesDetailRepository()
	salesDetailService := services.NewSalesDetailService(salesDetailRepository, db)
	salesDetailController := controllers.NewSalesDetailController(salesDetailService)

	salesReturnRepository := repository.NewSalesReturnRepository()
	salesReturnService := services.NewSalesReturnService(salesReturnRepository, db)
	salesReturnController := controllers.NewSalesReturnController(salesReturnService)
	
	salesReturnDetailRepository := repository.NewSalesReturnDetailRepository()
	salesReturnDetailService := services.NewSalesReturnDetailService(salesReturnDetailRepository, db)
	salesReturnDetailController := controllers.NewSalesReturnDetailController(salesReturnDetailService)
	
	salesConsignmentRepository := repository.NewSalesConsignmentRepository()
	salesConsignmentService := services.NewSalesConsignmentService(salesConsignmentRepository, db)
	salesConsignmentController := controllers.NewSalesConsignmentController(salesConsignmentService)
	
	salesConsignmentDetailRepository := repository.NewSalesConsignmentDetailRepository()
	salesConsignmentDetailService := services.NewSalesConsignmentDetailService(salesConsignmentDetailRepository, db)
	salesConsignmentDetailController := controllers.NewSalesConsignmentDetailController(salesConsignmentDetailService)
	
	cashRegisterRepository := repository.NewCashRegisterRepository()
	cashRegisterService := services.NewCashRegisterService(cashRegisterRepository, db)
	cashRegisterController := controllers.NewCashRegisterController(cashRegisterService)
	
	paymentRepository := repository.NewPaymentRepository()
	paymentService := services.NewPaymentService(paymentRepository, db)
	paymentController := controllers.NewPaymentController(paymentService)

	saleRepository := repository.NewSalesRepository()
	saleService := services.NewSalesService(saleRepository, salesDetailRepository, paymentRepository, customerRepository, db)
	saleController := controllers.NewSaleController(saleService)

	purchaseOrderRepository := repository.NewPurchaseOrderRepository()
	purchaseOrderService := services.NewPurchaseOrderService(purchaseOrderRepository, purchaseOrderDetailRepository, paymentRepository, db)
	purchaseOrderController := controllers.NewPurchaseOrderController(purchaseOrderService)

	productRepository := repository.NewProductRepository()
	productService := services.NewProductService(productRepository, db)
	productController := controllers.NewProductController(productService)
	
	roleRepository := repository.NewRoleRepository()
	roleService := services.NewRoleService(roleRepository, db)
	roleController := controllers.NewRoleController(roleService)
	
	rolehasPermissionRepository := repository.NewRoleHasPermissionRepository()
	rolehasPermissionService := services.NewRoleHasPermissionService(rolehasPermissionRepository, db)
	rolehasPermissionController := controllers.NewRoleHasPermissionController(rolehasPermissionService)

	permissionRepository := repository.NewPermissionRepository()
	permissionService := services.NewPermissionService(permissionRepository, db)
	permissionController := controllers.NewPermissionController(permissionService)

	api.GET("/users", userController.FindAll)
	api.GET("/users/:id", userController.FindById)
	api.POST("/users", userController.Create)
	api.POST("/user_datatables", userController.Datatable)
	api.PUT("/users/:id", userController.Update)
	api.DELETE("/users/:id", userController.Delete)

	api.GET("/divisions", divisionController.FindAll)
	api.GET("/divisions/:id", divisionController.FindById)
	api.POST("/divisions", divisionController.Create)
	api.POST("/division_datatables", divisionController.Datatable)
	api.PUT("/divisions/:id", divisionController.Update)
	api.DELETE("/divisions/:id", divisionController.Delete)
	
	api.GET("/categories", categoryController.FindAll)
	api.GET("/categories/:id", categoryController.FindById)
	api.POST("/categories", categoryController.Create)
	api.POST("/category_datatables", categoryController.Datatable)
	api.PUT("/categories/:id", categoryController.Update)
	api.DELETE("/categories/:id", categoryController.Delete)
	
	api.GET("/cake_variants", cakeVariantController.FindAll)
	api.GET("/cake_variants/:id", cakeVariantController.FindById)
	api.POST("/cake_variants", cakeVariantController.Create)
	api.POST("/cake_variant_datatables", cakeVariantController.Datatable)
	api.PUT("/cake_variants/:id", cakeVariantController.Update)
	api.DELETE("/cake_variants/:id", cakeVariantController.Delete)
	
	api.GET("/cake_types", cakeTypeController.FindAll)
	api.GET("/cake_types/:id", cakeTypeController.FindById)
	api.POST("/cake_types", cakeTypeController.Create)
	api.POST("/cake_type_datatables", cakeTypeController.Datatable)
	api.PUT("/cake_types/:id", cakeTypeController.Update)
	api.DELETE("/cake_types/:id", cakeTypeController.Delete)
	
	api.GET("/stores", storeController.FindAll)
	api.GET("/stores/:id", storeController.FindById)
	api.POST("/stores", storeController.Create)
	api.POST("/store_datatables", storeController.Datatable)
	api.PUT("/stores/:id", storeController.Update)
	api.DELETE("/stores/:id", storeController.Delete)
	
	api.GET("/suppliers", supplierController.FindAll)
	api.GET("/suppliers/:id", supplierController.FindById)
	api.POST("/suppliers", supplierController.Create)
	api.POST("/supplier_datatables", supplierController.Datatable)
	api.PUT("/suppliers/:id", supplierController.Update)
	api.DELETE("/suppliers/:id", supplierController.Delete)
	
	api.GET("/customers", customerController.FindAll)
	api.GET("/customers/:id", customerController.FindById)
	api.POST("/customers", customerController.Create)
	api.POST("/customer_datatables", customerController.Datatable)
	api.PUT("/customers/:id", customerController.Update)
	api.DELETE("/customers/:id", customerController.Delete)

	api.GET("/store_consignments", storeConsignmentController.FindAll)
	api.GET("/store_consignments/:id", storeConsignmentController.FindById)
	api.POST("/store_consignments", storeConsignmentController.Create)
	api.POST("/store_consignment_datatables", storeConsignmentController.Datatable)
	api.PUT("/store_consignments/:id", storeConsignmentController.Update)
	api.DELETE("/store_consignments/:id", storeConsignmentController.Delete)
	
	api.GET("/payment_methods", paymentMethodController.FindAll)
	api.GET("/payment_methods/:id", paymentMethodController.FindById)
	api.POST("/payment_methods", paymentMethodController.Create)
	api.POST("/payment_method_datatables", paymentMethodController.Datatable)
	api.PUT("/payment_methods/:id", paymentMethodController.Update)
	api.DELETE("/payment_methods/:id", paymentMethodController.Delete)
	
	api.GET("/sellers", sellerController.FindAll)
	api.GET("/sellers/:id", sellerController.FindById)
	api.POST("/sellers", sellerController.Create)
	api.POST("/seller_datatables", sellerController.Datatable)
	api.PUT("/sellers/:id", sellerController.Update)
	api.DELETE("/sellers/:id", sellerController.Delete)
	
	api.GET("/expense_categories", expenseCategoryController.FindAll)
	api.GET("/expense_categories/:id", expenseCategoryController.FindById)
	api.POST("/expense_categories", expenseCategoryController.Create)
	api.POST("/expense_category_datatables", expenseCategoryController.Datatable)
	api.PUT("/expense_categories/:id", expenseCategoryController.Update)
	api.DELETE("/expense_categories/:id", expenseCategoryController.Delete)
	
	api.GET("/raw_materials", rawMaterialController.FindAll)
	api.GET("/raw_materials/:id", rawMaterialController.FindById)
	api.POST("/raw_materials", rawMaterialController.Create)
	api.POST("/raw_material_datatables", rawMaterialController.Datatable)
	api.PUT("/raw_materials/:id", rawMaterialController.Update)
	api.DELETE("/raw_materials/:id", rawMaterialController.Delete)
	
	api.GET("/units", unitController.FindAll)
	api.GET("/units/:id", unitController.FindById)
	api.POST("/units", unitController.Create)
	api.POST("/unit_datatables", unitController.Datatable)
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
	api.POST("/custom_order_datatables", customOrderController.Datatable)
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
	api.POST("/expense_datatables", expenseController.Datatable)
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
	api.POST("/purchase_order_datatables", purchaseOrderController.Datatable)
	api.PUT("/purchase_orders/:id", purchaseOrderController.Update)
	api.DELETE("/purchase_orders/:id", purchaseOrderController.Delete)
	
	api.GET("/purchase_order_details", purchaseOrderDetailController.FindAll)
	api.GET("/purchase_order_details/:id", purchaseOrderDetailController.FindById)
	//api.POST("/purchase_order_details", purchaseOrderDetailController.Create)
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
	api.POST("/sale_datatables", saleController.Datatable)
	api.PUT("/sales/:id", saleController.Update)
	api.DELETE("/sales/:id", saleController.Delete)
	
	api.GET("/sales_details", salesDetailController.FindAll)
	api.GET("/sales_details/:id", salesDetailController.FindById)
	//api.POST("/sales_details", salesDetailController.Create)
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
	api.POST("/cash_register_datatables", cashRegisterController.Datatable)
	api.PUT("/cash_registers/:id", cashRegisterController.Update)
	api.DELETE("/cash_registers/:id", cashRegisterController.Delete)
	
	api.GET("/payments", paymentController.FindAll)
	api.GET("/payments/:id", paymentController.FindById)
	api.POST("/payments", paymentController.Create)
	api.PUT("/payments/:id", paymentController.Update)
	api.DELETE("/payments/:id", paymentController.Delete)
	
	api.GET("/products", productController.FindAll)
	api.GET("/products/:id", productController.FindById)
	api.POST("/products", productController.Create)
	api.POST("/product_datatables", productController.Datatable)
	api.PUT("/products/:id", productController.Update)
	api.DELETE("/products/:id", productController.Delete)
	
	api.GET("/roles", roleController.FindAll)
	api.GET("/roles/:id", roleController.FindById)
	api.POST("/roles", roleController.Create)
	api.POST("/role_datatables", roleController.Datatable)
	api.PUT("/roles/:id", roleController.Update)
	api.DELETE("/roles/:id", roleController.Delete)

	api.GET("/role_has_permissions", rolehasPermissionController.FindAll)
	api.GET("/role_has_permissions/:id", rolehasPermissionController.FindById)
	api.POST("/role_has_permissions", rolehasPermissionController.Create)
	api.PUT("/role_has_permissions/:id", rolehasPermissionController.Update)
	api.DELETE("/role_has_permissions", rolehasPermissionController.Delete)

	api.GET("/permissions", permissionController.FindAll)
	api.GET("/permissions/:id", permissionController.FindById)
	api.POST("/permissions", permissionController.Create)
	api.PUT("/permissions/:id", permissionController.Update)
	api.DELETE("/permissions/:id", permissionController.Delete)

	return e
}