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
	RoleHasPermissionService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	RoleHasPermissionServiceImpl struct {
		RoleHasPermissionRepository repository.RoleHasPermissionRepository
		db *gorm.DB
	}
)

func NewRoleHasPermissionService(RoleHasPermissionRepository repository.RoleHasPermissionRepository, db *gorm.DB) RoleHasPermissionService {
	return &RoleHasPermissionServiceImpl{
		RoleHasPermissionRepository: RoleHasPermissionRepository,
		db: db,
	}
}

func (service *RoleHasPermissionServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	o := new(domain.RoleHasPermission)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	roleHasPermissionRepo, err := service.RoleHasPermissionRepository.Create(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", roleHasPermissionRepo), err
}

func (service RoleHasPermissionServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.RoleHasPermission)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	roleHasPermissionRepo, err := service.RoleHasPermissionRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", roleHasPermissionRepo), err
}

func (service RoleHasPermissionServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.RoleHasPermission)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = service.RoleHasPermissionRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (service RoleHasPermissionServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	roleHasPermissionRepo, err := service.RoleHasPermissionRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", roleHasPermissionRepo), err
}

func (service RoleHasPermissionServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	roleHasPermissionRepo, err := service.RoleHasPermissionRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", roleHasPermissionRepo), err
}

