package router

import (
	"kalika-be/controllers"
	"kalika-be/middlewares"
	"kalika-be/repository"
	"kalika-be/services"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) *echo.Echo {
	//validate := validator.New()

	//USER THINGS
	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository, db)
	userController := controllers.NewUserController(userService)

	downloadController := controllers.NewDownloadController()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = middlewares.ErrorHandler

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		//AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	api := e.Group("/api")

	api.POST("/login", userController.Login)
	api.GET("/download", downloadController.DownloadPdf)

	api.Use(middlewares.Auth)

	paymentRepository := repository.NewPaymentRepository()
	paymentService := services.NewPaymentService(paymentRepository, db)
	paymentController := controllers.NewPaymentController(paymentService)
	
	generalSettingRepository := repository.NewGeneralSettingRepository()
	generalSettingService := services.NewGeneralSettingService(generalSettingRepository, db)
	generalSettingController := controllers.NewGeneralSettingController(generalSettingService)

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

	receivableDetailRepository := repository.NewReceivableDetailRepository()
	receivableDetailService := services.NewReceivableDetailService(receivableDetailRepository, db)
	receivableDetailController := controllers.NewReceivableDetailController(receivableDetailService)

	customOrderRepository := repository.NewCustomOrderRepository()
	customOrderService := services.NewCustomOrderService(customOrderRepository, receivableRepository, db)
	customOrderController := controllers.NewCustomOrderController(customOrderService)

	expenseDetailRepository := repository.NewExpenseDetailRepository()
	expenseDetailService := services.NewExpenseDetailService(expenseDetailRepository, db)
	expenseDetailController := controllers.NewExpenseDetailController(expenseDetailService)

	expenseRepository := repository.NewExpenseRepository()
	expenseService := services.NewExpenseService(expenseRepository, expenseDetailRepository, db)
	expenseController := controllers.NewExpenseController(expenseService)

	purchaseReturnRepository := repository.NewPurchaseReturnRepository()
	purchaseReturnDetailRepository := repository.NewPurchaseReturnDetailRepository()
	purchaseReturnService := services.NewPurchaseReturnService(purchaseReturnRepository, purchaseReturnDetailRepository, db)
	purchaseReturnController := controllers.NewPurchaseReturnController(purchaseReturnService)

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

	purchaseOrderDeliveryDetailRepository := repository.NewPurchaseOrderDeliveryDetailRepository()
	purchaseOrderDeliveryDetailService := services.NewPurchaseOrderDeliveryDetailService(purchaseOrderDeliveryDetailRepository, db)
	purchaseOrderDeliveryDetailController := controllers.NewPurchaseOrderDeliveryDetailController(purchaseOrderDeliveryDetailService)

	purchaseOrderDeliveryRepository := repository.NewPurchaseOrderDeliveryRepository()
	purchaseOrderDeliveryService := services.NewPurchaseOrderDeliveryService(purchaseOrderDeliveryRepository, purchaseOrderDeliveryDetailRepository, db)
	purchaseOrderDeliveryController := controllers.NewPurchaseOrderDeliveryController(purchaseOrderDeliveryService)

	salesDetailRepository := repository.NewSalesDetailRepository()
	salesDetailService := services.NewSalesDetailService(salesDetailRepository, db)
	salesDetailController := controllers.NewSalesDetailController(salesDetailService)

	salesReturnRepository := repository.NewSalesReturnRepository()
	salesReturnDetailRepository := repository.NewSalesReturnDetailRepository()
	salesReturnService := services.NewSalesReturnService(salesReturnRepository, salesReturnDetailRepository, db)
	salesReturnController := controllers.NewSalesReturnController(salesReturnService)

	salesReturnDetailService := services.NewSalesReturnDetailService(salesReturnDetailRepository, db)
	salesReturnDetailController := controllers.NewSalesReturnDetailController(salesReturnDetailService)

	salesConsignmentDetailRepository := repository.NewSalesConsignmentDetailRepository()
	salesConsignmentDetailService := services.NewSalesConsignmentDetailService(salesConsignmentDetailRepository, db)
	salesConsignmentDetailController := controllers.NewSalesConsignmentDetailController(salesConsignmentDetailService)

	cashRegisterRepository := repository.NewCashRegisterRepository()
	cashRegisterService := services.NewCashRegisterService(cashRegisterRepository, db)
	cashRegisterController := controllers.NewCashRegisterController(cashRegisterService)

	roleRepository := repository.NewRoleRepository()
	roleService := services.NewRoleService(roleRepository, db)
	roleController := controllers.NewRoleController(roleService)

	rolehasPermissionRepository := repository.NewRoleHasPermissionRepository()
	rolehasPermissionService := services.NewRoleHasPermissionService(rolehasPermissionRepository, db)
	rolehasPermissionController := controllers.NewRoleHasPermissionController(rolehasPermissionService)

	permissionRepository := repository.NewPermissionRepository()
	permissionService := services.NewPermissionService(permissionRepository, db)
	permissionController := controllers.NewPermissionController(permissionService)

	stockOpnameDetailRepository := repository.NewStockOpnameDetailRepository()

	stockOpnameRepository := repository.NewStockOpnameRepository()
	stockOpnameService := services.NewStockOpnameService(stockOpnameRepository, stockOpnameDetailRepository, db)
	stockOpnameController := controllers.NewStockOpnameController(stockOpnameService)

	storeMutationDetailRepository := repository.NewStoreMutationDetailRepository()

	storeMutationRepository := repository.NewStoreMutationRepository()
	storeMutationService := services.NewStoreMutationService(storeMutationRepository, storeMutationDetailRepository, db)
	storeMutationController := controllers.NewStoreMutationController(storeMutationService)

	productionRequestDetailRepository := repository.NewProductionRequestDetailRepository()

	productionRequestRepository := repository.NewProductionRequestRepository()
	productionRequestService := services.NewProductionRequestService(productionRequestRepository, productionRequestDetailRepository, db)
	productionRequestController := controllers.NewProductionRequestController(productionRequestService)

	productLocationRepository := repository.NewProductLocationRepository()
	productLocationService := services.NewProductLocationService(productLocationRepository, db)
	productLocationController := controllers.NewProductLocationController(productLocationService)

	productPriceRepository := repository.NewProductPriceRepository()
	productPriceService := services.NewProductPriceService(productPriceRepository, db)
	productPriceController := controllers.NewProductPriceController(productPriceService)

	checkStockService := services.NewCheckStockService(productLocationRepository, db)
	checkStockController := controllers.NewCheckStockController(checkStockService)

	productRepository := repository.NewProductRepository()
	productService := services.NewProductService(productRepository, productPriceRepository, productLocationRepository, db)
	productController := controllers.NewProductController(productService)

	recipeDetailRepository := repository.NewRecipeDetailRepository()

	recipeRepository := repository.NewRecipeRepository()
	recipeService := services.NewRecipeService(recipeRepository, recipeDetailRepository, db)
	recipeController := controllers.NewRecipeController(recipeService)

	saleRepository := repository.NewSalesRepository()
	saleService := services.NewSalesService(saleRepository, salesDetailRepository, paymentRepository, customerRepository, productLocationRepository, receivableRepository, db)
	saleController := controllers.NewSaleController(saleService)

	rawMaterialRepository := repository.NewRawMaterialRepository()
	rawMaterialService := services.NewRawMaterialService(rawMaterialRepository, productLocationRepository, db)
	rawMaterialController := controllers.NewRawMaterialController(rawMaterialService)

	purchaseOrderRepository := repository.NewPurchaseOrderRepository()
	purchaseOrderService := services.NewPurchaseOrderService(purchaseOrderRepository, purchaseOrderDetailRepository, paymentRepository, productLocationRepository, debtRepository, db)
	purchaseOrderController := controllers.NewPurchaseOrderController(purchaseOrderService)

	salesConsignmentRepository := repository.NewSalesConsignmentRepository()
	salesConsignmentService := services.NewSalesConsignmentService(salesConsignmentRepository, salesConsignmentDetailRepository, paymentRepository, storeConsignmentRepository, productLocationRepository, receivableRepository, db)
	salesConsignmentController := controllers.NewSalesConsignmentController(salesConsignmentService)

	reportRepository := repository.NewReportRepository()
	reportService := services.NewReportService(reportRepository, db)
	reportController := controllers.NewReportController(reportService)

	api.GET("/general_settings", generalSettingController.FindAll)
	api.POST("/general_settings", generalSettingController.Create)

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
	api.POST("/report_custom_order_datatables", customOrderController.ReportDatatable)
	api.POST("/custom_order_pdf", customOrderController.GeneratePdf)
	api.PUT("/custom_orders/:id", customOrderController.Update)
	api.DELETE("/custom_orders/:id", customOrderController.Delete)

	api.POST("/check_stock_datatables", checkStockController.Datatable)
	api.POST("/check_stock_pdf", checkStockController.GeneratePdf)

	api.GET("/debts", debtController.FindAll)
	api.GET("/debts/:id", debtController.FindById)
	api.POST("/debts", debtController.Create)
	api.POST("/debt_datatables", debtController.Datatable)
	api.POST("/report_debt_datatables", debtController.ReportDatatable)
	api.POST("/debt_pdf", debtController.GeneratePdf)
	api.PUT("/debts/:id", debtController.Update)
	api.DELETE("/debts/:id", debtController.Delete)

	api.GET("/debt_details", debtDetailController.FindAll)
	api.GET("/debt_details/:id", debtDetailController.FindById)
	api.POST("/debt_details", debtDetailController.Create)
	api.POST("/debt_detail_datatables", debtDetailController.Datatable)
	api.PUT("/debt_details/:id", debtDetailController.Update)
	api.DELETE("/debt_details/:id", debtDetailController.Delete)

	api.GET("/receivables", receivableController.FindAll)
	api.GET("/receivables/:id", receivableController.FindById)
	api.POST("/receivables", receivableController.Create)
	api.POST("/receivable_datatables", receivableController.Datatable)
	api.POST("/report_receivable_datatables", receivableController.ReportDatatable)
	api.POST("/receivable_pdf", receivableController.GeneratePdf)
	api.PUT("/receivables/:id", receivableController.Update)
	api.DELETE("/receivables/:id", receivableController.Delete)

	api.GET("/receivable_details", receivableDetailController.FindAll)
	api.GET("/receivable_details/:id", receivableDetailController.FindById)
	api.POST("/receivable_details", receivableDetailController.Create)
	api.POST("/receivable_detail_datatables", receivableDetailController.Datatable)
	api.PUT("/receivable_details/:id", receivableDetailController.Update)
	api.DELETE("/receivable_details/:id", receivableDetailController.Delete)

	api.GET("/expenses", expenseController.FindAll)
	api.GET("/expenses/:id", expenseController.FindById)
	api.POST("/expenses", expenseController.Create)
	api.POST("/expense_datatables", expenseController.Datatable)
	api.POST("/report_expense_datatables", expenseController.ReportDatatable)
	api.POST("/expense_pdf", expenseController.GeneratePdf)
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
	api.POST("/purchase_return_datatables", purchaseReturnController.Datatable)
	api.POST("/report_purchase_return_datatables", purchaseReturnController.ReportDatatable)
	api.POST("/purchase_return_pdf", purchaseReturnController.GeneratePdf)
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
	api.POST("/report_purchase_order_datatables", purchaseOrderController.ReportDatatable)
	api.POST("/purchase_order_pdf", purchaseOrderController.GeneratePdf)
	api.PUT("/purchase_orders/:id", purchaseOrderController.Update)
	api.DELETE("/purchase_orders/:id", purchaseOrderController.Delete)

	api.GET("/recipes", recipeController.FindAll)
	api.GET("/recipes/:id", recipeController.FindById)
	api.GET("/recipe_by_product_id/:id", recipeController.FindByProductId)
	api.POST("/recipes", recipeController.Create)
	api.POST("/recipe_datatables", recipeController.Datatable)
	api.PUT("/recipes/:id", recipeController.Update)
	api.DELETE("/recipes/:id", recipeController.Delete)

	api.GET("/purchase_order_details", purchaseOrderDetailController.FindAll)
	api.GET("/purchase_order_details/:id", purchaseOrderDetailController.FindById)
	//api.POST("/purchase_order_details", purchaseOrderDetailController.Create)
	//api.PUT("/purchase_order_details/:id", purchaseOrderDetailController.Update)
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
	api.POST("/report_sale_datatables", saleController.ReportDatatable)
	api.POST("/sale_pdf", saleController.GeneratePdf)
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
	api.POST("/report_sales_return_datatables", salesReturnController.ReportDatatable)
	api.POST("/sales_return_datatables", salesReturnController.Datatable)
	api.POST("/sales_return_pdf", salesReturnController.GeneratePdf)
	api.PUT("/sales_returns/:id", salesReturnController.Update)
	api.DELETE("/sales_returns/:id", salesReturnController.Delete)

	api.GET("/sales_return_details", salesReturnDetailController.FindAll)
	api.GET("/sales_return_details/:id", salesReturnDetailController.FindById)
	api.POST("/sales_return_details", salesReturnDetailController.Create)
	api.PUT("/sales_return_details/:id", salesReturnDetailController.Update)
	api.DELETE("/sales_return_details/:id", salesReturnDetailController.Delete)

	api.GET("/sales_consignments", salesConsignmentController.FindAll)
	api.GET("/sales_consignments/:id", salesConsignmentController.FindById)
	api.POST("/sales_consignments", salesConsignmentController.Create)
	api.POST("/sales_consignment_datatables", salesConsignmentController.Datatable)
	api.POST("/report_sales_consignment_datatables", salesConsignmentController.ReportDatatable)
	api.POST("/sales_consignment_pdf", salesConsignmentController.GeneratePdf)
	api.PUT("/sales_consignments/:id", salesConsignmentController.Update)
	api.DELETE("/sales_consignments/:id", salesConsignmentController.Delete)

	api.GET("/sales_consignment_details", salesConsignmentDetailController.FindAll)
	api.GET("/sales_consignment_details/:id", salesConsignmentDetailController.FindById)
	//api.POST("/sales_consignment_details", salesConsignmentDetailController.Create)
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
	api.POST("/payment_datatables", paymentController.Datatable)
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

	api.GET("/stock_opnames", stockOpnameController.FindAll)
	api.GET("/stock_opnames_pdf/:id", stockOpnameController.GeneratePdf)
	api.GET("/stock_opnames/:id", stockOpnameController.FindById)
	api.POST("/stock_opnames", stockOpnameController.Create)
	api.POST("/stock_opname_datatables", stockOpnameController.Datatable)
	api.PUT("/stock_opnames/:id", stockOpnameController.Update)
	api.DELETE("/stock_opnames/:id", stockOpnameController.Delete)

	api.GET("/store_mutations", storeMutationController.FindAll)
	api.GET("/store_mutations_pdf/:id", storeMutationController.GeneratePdf)
	api.GET("/store_mutations/:id", storeMutationController.FindById)
	api.POST("/store_mutations", storeMutationController.Create)
	api.POST("/store_mutation_datatables", storeMutationController.Datatable)
	api.PUT("/store_mutations/:id", storeMutationController.Update)
	api.DELETE("/store_mutations/:id", storeMutationController.Delete)

	api.GET("/production_requests", productionRequestController.FindAll)
	api.GET("/production_requests_pdf/:id", productionRequestController.GeneratePdf)
	api.GET("/production_requests/:id", productionRequestController.FindById)
	api.POST("/production_requests", productionRequestController.Create)
	api.POST("/production_request_datatables", productionRequestController.Datatable)
	api.PUT("/production_requests/:id", productionRequestController.Update)
	api.DELETE("/production_requests/:id", productionRequestController.Delete)

	api.GET("/product_locations", productLocationController.FindAll)
	api.GET("/product_locations/:id", productLocationController.FindById)
	api.POST("/product_locations", productLocationController.Create)
	api.PUT("/product_locations/:id", productLocationController.Update)
	api.DELETE("/product_locations/:id", productLocationController.Delete)

	api.GET("/product_prices", productPriceController.FindAll)
	api.GET("/product_prices/:id", productPriceController.FindById)

	api.GET("/profit_loss", reportController.ProfitLoss)
	api.GET("/ledger_receivables", reportController.ReceivableLedger)
	api.GET("/ledger_debts", reportController.DebtLedger)
	api.GET("/ledger_cash_banks", reportController.CashBankLedger)

	return e
}
