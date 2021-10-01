package services

import (
	"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
	"kalika-be/repository"
	"strings"
)

type (
	UserService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		Update(ctx echo.Context, id int) (res web.Response, err error)
		Delete(ctx echo.Context, id int) (res web.Response, err error)
		FindById(ctx echo.Context, id int) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
		Login(ctx echo.Context) (map[string]interface{}, error)
		Datatable(ctx echo.Context) (res web.Datatable, err error)
	}

	UserServiceImpl struct {
		UserRepository repository.UserRepository
		db *gorm.DB
	}
)

func NewUserService(userRepository repository.UserRepository, db *gorm.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		db: db,
	}
}

func (u *UserServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	userRepo := domain.User{}

	o := new(domain.User)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}

	tx := u.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if o.Id > 0 {
		userRepo, err = u.UserRepository.Update(ctx, tx, o)
	} else {
		userRepo, err = u.UserRepository.Create(ctx, tx, o)
	}

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", userRepo), err
}

func (u *UserServiceImpl) Update(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.User)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Binding", nil), err
	}
	o.Id = id

	tx := u.db.Begin()
	defer helpers.CommitOrRollback(tx)

	userRepo, err := u.UserRepository.Update(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengubah Data", userRepo), err
}

func (u *UserServiceImpl) Delete(ctx echo.Context, id int) (res web.Response, err error) {
	o := new(domain.User)
	if err := ctx.Bind(o); err != nil {
		return helpers.Response(err.Error(), "Error Data Bnding", nil), err
	}
	o.Id = id

	tx := u.db.Begin()
	defer helpers.CommitOrRollback(tx)

	_, err = u.UserRepository.Delete(ctx, tx, o)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Menghapus Data", true), err
}

func (u *UserServiceImpl) FindById(ctx echo.Context, id int) (res web.Response, err error) {
	tx := u.db.Begin()
	defer helpers.CommitOrRollback(tx)

	userRepo, err := u.UserRepository.FindById(ctx, tx, "id", helpers.IntToString(id))

	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("OK", "Sukses Mengambil Data", userRepo), err
}

func (u *UserServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := u.db.Begin()
	defer helpers.CommitOrRollback(tx)

	userRepo, err := u.UserRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", userRepo), err
}

func (u *UserServiceImpl) Login(ctx echo.Context) (map[string]interface{}, error) {
	o := new(domain.User)
	l := new(web.Login)
	if err := ctx.Bind(o); err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}
	tx := u.db.Begin()
	defer helpers.CommitOrRollback(tx)

	userRepo, err := u.UserRepository.Login(ctx, tx, o)
	if err != nil {
		fmt.Println("ERROR LOGIN")
		fmt.Println(err.Error())
		res := map[string]interface{}{}
		res["code"] = 404
		res["message"] = "Username and Password Missmatch"
		res["data"] = nil
		return res, err
	}

	l.Username = userRepo.Username
	l.Name = userRepo.Name
	l.RoleId = userRepo.RoleId
	l.StoreId = userRepo.StoreId

	l.Token = helpers.JwtGenerator(userRepo.Name, userRepo.Username, helpers.IntToString(userRepo.RoleId), helpers.IntToString(userRepo.StoreId), config.Get("JWT_KEY").String())

	toMap, err := helpers.StructToMap(l)
	if err != nil {
		fmt.Println("ERROR LOGIN")
		fmt.Println(err)
	}
	return toMap, nil
}

func (u *UserServiceImpl) Datatable(ctx echo.Context) (res web.Datatable, err error) {
	params,_ := ctx.FormParams()

	tx := u.db.Begin()
	defer helpers.CommitOrRollback(tx)

	draw := strings.TrimSpace(params.Get("draw"))
	limit := strings.TrimSpace(params.Get("length"))
	start := strings.TrimSpace(params.Get("start"))
	search := strings.TrimSpace(params.Get("search[value]"))

	userRepo, totalData, totalFiltered, _ := u.UserRepository.Datatable(ctx, tx, draw, limit, start, search)

	data := make([]interface{}, 0)
	for _, v := range userRepo {
		v.Action = `<div class="flex">`
		v.Action += `<button type="button" class="btn-edit flex mr-3" id="edit-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="check-square" class="w-4 h-4 mr-1"></i> Edit </button>`
		v.Action += `<button type="button" class="btn-delete flex text-theme-6" id="delete-data" data-id=`+helpers.IntToString(v.Id)+`> <i data-feather="trash-2" class="w-4 h-4 mr-1"></i> Delete </button>`
		v.Action += `</div>`

		data = append(data, v)
	}
	res.Data = data
	res.Order = helpers.ParseFormCollection(ctx.Request(), "order")
	res.Draw = helpers.StringToInt(draw)
	res.RecordsFiltered = totalFiltered
	res.RecordsTotal = totalData

	return res, nil
}

