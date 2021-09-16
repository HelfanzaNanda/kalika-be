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
	UnitConversionService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	UnitConversionServiceImpl struct {
		UnitConversionRepository repository.UnitConversionRepository
		db *gorm.DB
	}
)

func NewUnitConversionService(UnitConversionRepository repository.UnitConversionRepository, db *gorm.DB) UnitConversionService {
	return &UnitConversionServiceImpl{
		UnitConversionRepository: UnitConversionRepository,
		db: db,
	}
}

func (service *UnitConversionServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.UnitConversion)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	unitConversionRepo, err := service.UnitConversionRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", unitConversionRepo), err
}

func (service UnitConversionServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.UnitConversion)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	unitConversionRepo, err := service.UnitConversionRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", unitConversionRepo), err
}

func (service UnitConversionServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.UnitConversion)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.UnitConversionRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service UnitConversionServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	unitConversionRepo, err := service.UnitConversionRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", unitConversionRepo), err
}

func (service UnitConversionServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	unitConversionRepo, err := service.UnitConversionRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", unitConversionRepo), err
}

