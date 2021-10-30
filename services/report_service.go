package services

import (
	"fmt"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	//"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	ReportService interface {
		ProfitLoss(ctx echo.Context) (res web.Response, err error)
		ReceivableLedger(ctx echo.Context) (res web.Response, err error)
		DebtLedger(ctx echo.Context) (res web.Response, err error)
		CashBankLedger(ctx echo.Context) (res web.Response, err error)
		GeneratePdfProfitLoss(ctx echo.Context) (res web.Response, err error)
		GeneratePdfReceivableLedger(ctx echo.Context) (res web.Response, err error)
		GeneratePdfDebtLedger(ctx echo.Context) (res web.Response, err error)
		GeneratePdfCashBankLedger(ctx echo.Context) (res web.Response, err error)
	}

	ReportServiceImpl struct {
		ReportRepository repository.ReportRepository
		db *gorm.DB
	}
)

func NewReportService(ReportRepository repository.ReportRepository, db *gorm.DB) ReportService {
	return &ReportServiceImpl{
		ReportRepository: ReportRepository,
		db: db,
	}
}

func (service *ReportServiceImpl) ProfitLoss(ctx echo.Context) (res web.Response, err error) {
	o := new(web.ProfitLossFilter)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	profitLossRepo, err := service.ReportRepository.ProfitLoss(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", profitLossRepo), err
}

func (service *ReportServiceImpl) ReceivableLedger(ctx echo.Context) (res web.Response, err error) {
	o := new(web.DateRange)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	profitLossRepo, err := service.ReportRepository.ReceivableLedger(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", profitLossRepo), err
}

func (service *ReportServiceImpl) DebtLedger(ctx echo.Context) (res web.Response, err error) {
	o := new(web.DateRange)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	profitLossRepo, err := service.ReportRepository.DebtLedger(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", profitLossRepo), err
}

func (service *ReportServiceImpl) CashBankLedger(ctx echo.Context) (res web.Response, err error) {
	o := new(web.DateRange)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	profitLossRepo, err := service.ReportRepository.CashBankLedger(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", profitLossRepo), err
}

func (service ReportServiceImpl) GeneratePdfProfitLoss(ctx echo.Context) (res web.Response, err error) {
	o := new(web.ProfitLossFilter)

	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	
	reportRepo, err := service.ReportRepository.ProfitLoss(ctx, tx, o)
	totalSales := reportRepo.Sales + reportRepo.TotalCogs + reportRepo.CustomOrder + reportRepo.SalesConsignment + reportRepo.ReceivablePayment
	totalExpense := reportRepo.DebtPayment + reportRepo.SalesReturn + reportRepo.TotalCost
	totalProfitLoss := totalSales - totalExpense

	var datas [][]string
	item := []string{}
	item = append(item, "Laporan Pendapatan")
	datas = append(datas, item)

	items := []string{"Penjualan Tunai", "Total HPP", "Penjualan Pesenan", "Penjualan Pesenan", "Pembayaran Piutang"}
	for index, value := range items {
		item := []string{}
		item = append(item, value)
		if  index == 0{
			item = append(item, helpers.FormatRupiah(reportRepo.Sales))
		}else if index == 1 {
			item = append(item, helpers.FormatRupiah(reportRepo.Sales))
		}else if index == 2{
			item = append(item, helpers.FormatRupiah(reportRepo.CustomOrder))
		}else if index == 3{
			item = append(item, helpers.FormatRupiah(reportRepo.SalesConsignment))
		}else if index == 4{
			item = append(item, helpers.FormatRupiah(reportRepo.ReceivablePayment))
		}
		datas = append(datas, item)
	}

	item = []string{}
	item = append(item, "Laporan Pengeluaran")
	datas = append(datas, item)

	items = []string{"Pembayaran Hutang", "Retur Penjualan"}
	for index, value := range items {
		item := []string{}
		item = append(item, value)
		if  index == 0{
			item = append(item, helpers.FormatRupiah(reportRepo.DebtPayment))
		}else if index == 1 {
			item = append(item, helpers.FormatRupiah(reportRepo.SalesReturn))
		}
		datas = append(datas, item)
	}

	item = []string{}
	item = append(item, fmt.Sprintf("Laporan Biaya (%s)", helpers.FormatRupiah(reportRepo.TotalCost)))
	datas = append(datas, item)

	for _, value := range reportRepo.Cost {
		item := []string{}
		item = append(item, value.Name)
		item = append(item, helpers.FormatRupiah(value.Total))
		datas = append(datas, item)
	}

	item = []string{}
	if totalProfitLoss > 0 {
		item = append(item, fmt.Sprintf("Laba %s", helpers.FormatRupiah(totalProfitLoss)))
	}else {
		item = append(item, fmt.Sprintf("Rugi %s", helpers.FormatRupiah(totalProfitLoss)))
	}
	datas = append(datas, item)

	title := "laporan_laba_rugi"
	headings := []string{" ", " "}
	footer := map[string]float64{}
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer, o.StartDate, o.EndDate)
	
	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}

func (service ReportServiceImpl) GeneratePdfReceivableLedger(ctx echo.Context) (res web.Response, err error) {
	o := new(web.DateRange)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	
	reportRepo, err := service.ReportRepository.ReceivableLedger(ctx, tx, o)
	var datas [][]string
	var debit float64 = 0
	var credit float64 = 0
	var balance float64 = 0
	for _, item := range reportRepo {
		froot := []string{}
		froot = append(froot, item.Date.Format("02 Jan 2006 15:04:05"))
		froot = append(froot, item.Number)
		froot = append(froot, item.Customer)
		froot = append(froot, helpers.FormatRupiah(item.Debit))
		froot = append(froot, helpers.FormatRupiah(item.Credit))
		froot = append(froot, helpers.FormatRupiah(item.Balance))
		datas = append(datas, froot)
		debit += item.Debit
		credit += item.Credit
		balance += item.Balance
	}
	title := "laporan_buku_besar_piutang"
	headings := []string{"Tanggal", "No. Ref", "Kustomer", "Debit", "Kredit", "Saldo"}
	footer := map[string]float64{}
	footer["Total Debit"] = debit
	footer["Total Credit"] = credit
	footer["Total Balance"] = balance
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer, o.StartDate, o.EndDate)
	
	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}
func (service ReportServiceImpl) GeneratePdfDebtLedger(ctx echo.Context) (res web.Response, err error) {
	o := new(web.DateRange)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	
	reportRepo, err := service.ReportRepository.DebtLedger(ctx, tx, o)
	var datas [][]string
	var debit float64 = 0
	var credit float64 = 0
	var balance float64 = 0
	for _, item := range reportRepo {
		froot := []string{}
		froot = append(froot, item.Date.Format("02 Jan 2006 15:04:05"))
		froot = append(froot, item.Number)
		froot = append(froot, item.Supplier)
		froot = append(froot, helpers.FormatRupiah(item.Debit))
		froot = append(froot, helpers.FormatRupiah(item.Credit))
		froot = append(froot, helpers.FormatRupiah(item.Balance))
		datas = append(datas, froot)
		debit += item.Debit
		credit += item.Credit
		balance += item.Balance
	}
	title := "laporan_buku_besar_hutang"
	headings := []string{"Tanggal", "No. Ref", "Supplier", "Debit", "Kredit", "Saldo"}
	footer := map[string]float64{}
	footer["Total Debit"] = debit
	footer["Total Credit"] = credit
	footer["Total Balance"] = balance
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer, o.StartDate, o.EndDate)
	
	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}
func (service ReportServiceImpl) GeneratePdfCashBankLedger(ctx echo.Context) (res web.Response, err error) {
	o := new(web.DateRange)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	
	reportRepo, err := service.ReportRepository.CashBankLedger(ctx, tx, o)
	var datas [][]string
	var debit float64 = 0
	var credit float64 = 0
	for _, item := range reportRepo {
		froot := []string{}
		froot = append(froot, item.Date.Format("02 Jan 2006 15:04:05"))
		//froot = append(froot, item.Number)
		froot = append(froot, item.Type)
		froot = append(froot, item.PaymentMethod)
		froot = append(froot, helpers.FormatRupiah(item.Debit))
		froot = append(froot, helpers.FormatRupiah(item.Credit))
		froot = append(froot, helpers.FormatRupiah(item.Balance))
		datas = append(datas, froot)
		debit += item.Debit
		credit += item.Credit
	}
	title := "laporan_buku_besar_kas_bank"
	headings := []string{"Tanggal", "Tipe", "Metode", "Debit", "Kredit", "Saldo"}
	footer := map[string]float64{}
	footer["Total Balance"] = debit - credit
	footer["Total Debit"] = debit
	footer["Total Credit"] = credit
	resultPdf, err := helpers.GeneratePdf(ctx, title, headings, datas, footer, o.StartDate, o.EndDate)
	
	return helpers.Response("OK", "Sukses Export PDF", resultPdf), err
}
