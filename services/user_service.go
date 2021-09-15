package services

import (
	"fmt"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	UserService interface {
		Create(ctx echo.Context) (map[string]interface{}, error)
		Update(ctx echo.Context)
		Delete(ctx echo.Context)
		FindById(ctx echo.Context)
		FindAll(ctx echo.Context)
		Login(ctx echo.Context) (map[string]interface{}, error)
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

func (u *UserServiceImpl) Create(ctx echo.Context) (map[string]interface{}, error) {
	o := new(models.User)
	if err := ctx.Bind(o); err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	tx := u.db.Begin()
	defer helpers.CommitOrRollback(tx)

	userRepo, err := u.UserRepository.Create(ctx, tx, o)
	if err != nil {
		fmt.Println("ERROR CREATING USER")
		fmt.Println(err)
	}

	toMap, err := helpers.StructToMap(userRepo)
	if err != nil {
		fmt.Println("ERROR STRUCT TO MAP")
		fmt.Println(err)
	}

	return toMap, nil
}

func (u UserServiceImpl) Update(ctx echo.Context) {
	panic("implement me")
}

func (u UserServiceImpl) Delete(ctx echo.Context) {
	panic("implement me")
}

func (u UserServiceImpl) FindById(ctx echo.Context) {
	panic("implement me")
}

func (u UserServiceImpl) FindAll(ctx echo.Context) {
	panic("implement me")
}

func (u UserServiceImpl) Login(ctx echo.Context) (map[string]interface{}, error) {
	o := new(models.User)
	l := new(web.Login)
	if err := ctx.Bind(o); err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}
	tx := u.db.Begin()

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

