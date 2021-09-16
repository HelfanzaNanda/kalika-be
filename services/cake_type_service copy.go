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
	CakeTypeService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	CakeTypeServiceImpl struct {
		CakeTypeRepository repository.CakeTypeRepository
		db *gorm.DB
	}
)

func NewCakeTypeService(CakeTypeRepository repository.CakeTypeRepository, db *gorm.DB) CakeTypeService {
	return &CakeTypeServiceImpl{
		CakeTypeRepository: CakeTypeRepository,
		db: db,
	}
}

func (service *CakeTypeServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.CakeType)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cakeTypeRepo, err := service.CakeTypeRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", cakeTypeRepo), err
}

func (service CakeTypeServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CakeType)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cakeTypeRepo, err := service.CakeTypeRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", cakeTypeRepo), err
}

func (service CakeTypeServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CakeType)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.CakeTypeRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service CakeTypeServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cakeTypeRepo, err := service.CakeTypeRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", cakeTypeRepo), err
}

func (service CakeTypeServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cakeTypeRepo, err := service.CakeTypeRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", cakeTypeRepo), err
}

