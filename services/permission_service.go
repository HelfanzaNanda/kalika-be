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
	PermissionService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	PermissionServiceImpl struct {
		PermissionRepository repository.PermissionRepository
		db *gorm.DB
	}
)

func NewPermissionService(PermissionRepository repository.PermissionRepository, db *gorm.DB) PermissionService {
	return &PermissionServiceImpl{
		PermissionRepository: PermissionRepository,
		db: db,
	}
}

func (service *PermissionServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.Permission)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	roleHasPermissionRepo, err := service.PermissionRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", roleHasPermissionRepo), err
}

func (service PermissionServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Permission)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	roleHasPermissionRepo, err := service.PermissionRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", roleHasPermissionRepo), err
}

func (service PermissionServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.Permission)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.PermissionRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service PermissionServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	roleHasPermissionRepo, err := service.PermissionRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", roleHasPermissionRepo), err
}

func (service PermissionServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	roleHasPermissionRepo, err := service.PermissionRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", roleHasPermissionRepo), err
}

