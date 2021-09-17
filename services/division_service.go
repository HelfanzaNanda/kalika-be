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
	DivisionService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	DivisionServiceImpl struct {
		DivisionRepository repository.DivisionRepository
		db *gorm.DB
	}
)

func NewDivisionService(DivisionRepository repository.DivisionRepository, db *gorm.DB) DivisionService {
	return &DivisionServiceImpl{
		DivisionRepository: DivisionRepository,
		db: db,
	}
}

func (service *DivisionServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.Division)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	divisionRepo, err := service.DivisionRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", divisionRepo), err
}

func (service DivisionServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Division)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	divisionRepo, err := service.DivisionRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", divisionRepo), err
}

func (service DivisionServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Division)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.DivisionRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service DivisionServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	divisionRepo, err := service.DivisionRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", divisionRepo), err
}

func (service DivisionServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	divisionRepo, err := service.DivisionRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", divisionRepo), err
}

