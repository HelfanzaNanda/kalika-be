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
	ReceivableService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	ReceivableServiceImpl struct {
		ReceivableRepository repository.ReceivableRepository
		db *gorm.DB
	}
)

func NewReceivableService(ReceivableRepository repository.ReceivableRepository, db *gorm.DB) ReceivableService {
	return &ReceivableServiceImpl{
		ReceivableRepository: ReceivableRepository,
		db: db,
	}
}

func (service *ReceivableServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.Receivable)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	receivableRepo, err := service.ReceivableRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", receivableRepo), err
}

func (service ReceivableServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Receivable)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	receivableRepo, err := service.ReceivableRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", receivableRepo), err
}

func (service ReceivableServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Receivable)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.ReceivableRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service ReceivableServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	receivableRepo, err := service.ReceivableRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", receivableRepo), err
}

func (service ReceivableServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	receivableRepo, err := service.ReceivableRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", receivableRepo), err
}

