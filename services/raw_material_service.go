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
	RawMaterialService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	RawMaterialServiceImpl struct {
		RawMaterialRepository repository.RawMaterialRepository
		db *gorm.DB
	}
)

func NewRawMaterialService(RawMaterialRepository repository.RawMaterialRepository, db *gorm.DB) RawMaterialService {
	return &RawMaterialServiceImpl{
		RawMaterialRepository: RawMaterialRepository,
		db: db,
	}
}

func (service *RawMaterialServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.RawMaterial)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	rawMaterialRepo, err := service.RawMaterialRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", rawMaterialRepo), err
}

func (service RawMaterialServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.RawMaterial)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	rawMaterialRepo, err := service.RawMaterialRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", rawMaterialRepo), err
}

func (service RawMaterialServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.RawMaterial)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.RawMaterialRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service RawMaterialServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	rawMaterialRepo, err := service.RawMaterialRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", rawMaterialRepo), err
}

func (service RawMaterialServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	rawMaterialRepo, err := service.RawMaterialRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", rawMaterialRepo), err
}

