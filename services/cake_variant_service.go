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
	CakeVariantService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	CakeVariantServiceImpl struct {
		CakeVariantRepository repository.CakeVariantRepository
		db *gorm.DB
	}
)

func NewCakeVariantService(CakeVariantRepository repository.CakeVariantRepository, db *gorm.DB) CakeVariantService {
	return &CakeVariantServiceImpl{
		CakeVariantRepository: CakeVariantRepository,
		db: db,
	}
}

func (service *CakeVariantServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.CakeVariant)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cakeVariantRepo, err := service.CakeVariantRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", cakeVariantRepo), err
}

func (service CakeVariantServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CakeVariant)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cakeVariantRepo, err := service.CakeVariantRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", cakeVariantRepo), err
}

func (service CakeVariantServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.CakeVariant)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.CakeVariantRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service CakeVariantServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cakeVariantRepo, err := service.CakeVariantRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", cakeVariantRepo), err
}

func (service CakeVariantServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cakeVariantRepo, err := service.CakeVariantRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", cakeVariantRepo), err
}

