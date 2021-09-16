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
	CategoryService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	CategoryServiceImpl struct {
		CategoryRepository repository.CategoryRepository
		db *gorm.DB
	}
)

func NewCategoryService(CategoryRepository repository.CategoryRepository, db *gorm.DB) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: CategoryRepository,
		db: db,
	}
}

func (service *CategoryServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.Category)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	categoryRepo, err := service.CategoryRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", categoryRepo), err
}

func (service CategoryServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Category)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	categoryRepo, err := service.CategoryRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", categoryRepo), err
}

func (service CategoryServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Category)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.CategoryRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service CategoryServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	categoryRepo, err := service.CategoryRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", categoryRepo), err
}

func (service CategoryServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	categoryRepo, err := service.CategoryRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", categoryRepo), err
}

