package services

import (
	//"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	//"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	PaymentService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	PaymentServiceImpl struct {
		PaymentRepository repository.PaymentRepository
		db *gorm.DB
	}
)

func NewPaymentService(PaymentRepository repository.PaymentRepository, db *gorm.DB) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository: PaymentRepository,
		db: db,
	}
}

func (service *PaymentServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.Payment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentRepo, err := service.PaymentRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", paymentRepo), err
}

func (service PaymentServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Payment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentRepo, err := service.PaymentRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", paymentRepo), err
}

func (service PaymentServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Payment)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.PaymentRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service PaymentServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentRepo, err := service.PaymentRepository.FindById(ctx, tx, "id", helpers.IntToString(id), map[string][]string{})

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", paymentRepo), err
}

func (service PaymentServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentRepo, err := service.PaymentRepository.FindAll(ctx, tx, map[string][]string{})

	return helpers.Response("OK", "Sukses Mengambil Data", paymentRepo), err
}

