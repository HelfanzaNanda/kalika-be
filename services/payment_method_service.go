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
	PaymentMethodService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	PaymentMethodServiceImpl struct {
		PaymentMethodRepository repository.PaymentMethodRepository
		db *gorm.DB
	}
)

func NewPaymentMethodService(PaymentMethodRepository repository.PaymentMethodRepository, db *gorm.DB) PaymentMethodService {
	return &PaymentMethodServiceImpl{
		PaymentMethodRepository: PaymentMethodRepository,
		db: db,
	}
}

func (service *PaymentMethodServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.PaymentMethod)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentMethodRepo, err := service.PaymentMethodRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", paymentMethodRepo), err
}

func (service PaymentMethodServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PaymentMethod)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentMethodRepo, err := service.PaymentMethodRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", paymentMethodRepo), err
}

func (service PaymentMethodServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.PaymentMethod)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.PaymentMethodRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service PaymentMethodServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentMethodRepo, err := service.PaymentMethodRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", paymentMethodRepo), err
}

func (service PaymentMethodServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	paymentMethodRepo, err := service.PaymentMethodRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", paymentMethodRepo), err
}

