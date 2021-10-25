package services

import (
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
	o := new(web.ProfitLossReport)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	profitLossRepo, err := service.ReportRepository.ProfitLoss(ctx, tx)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", profitLossRepo), err
}

func (service *ReportServiceImpl) ReceivableLedger(ctx echo.Context) (res web.Response, err error) {
	o := new(web.ReportLedgerReceivable)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	profitLossRepo, err := service.ReportRepository.ReceivableLedger(ctx, tx)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", profitLossRepo), err
}

func (service *ReportServiceImpl) DebtLedger(ctx echo.Context) (res web.Response, err error) {
	o := new(web.ReportLedgerDebt)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	profitLossRepo, err := service.ReportRepository.DebtLedger(ctx, tx)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", profitLossRepo), err
}

func (service *ReportServiceImpl) CashBankLedger(ctx echo.Context) (res web.Response, err error) {
	o := new(web.ReportLedgerCashBank)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	profitLossRepo, err := service.ReportRepository.CashBankLedger(ctx, tx)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", profitLossRepo), err
}