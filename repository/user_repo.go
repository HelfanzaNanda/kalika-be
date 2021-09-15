package repository

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models"
)

type (
	UserRepository interface{
		Create(ctx echo.Context, db *gorm.DB, user *models.User) (models.User, error)
		Update(ctx echo.Context, db *gorm.DB, user *models.User) (models.User, error)
		Delete(ctx echo.Context, db *gorm.DB, user *models.User) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (models.User, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]models.User, error)
		Login(ctx echo.Context, db *gorm.DB, user *models.User) (models.User, error)
	}

	UserRepositoryImpl struct {

	}
)

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u UserRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, user *models.User) (models.User, error) {
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("ERROR HASHING PASSWORD")
		fmt.Println(err)
	}
	//fmt.Println(ctx.Get("userInfo").(map[string]interface{})["name"])
	user.Password = string(hashedPassword)
	db.Create(user)
	userRes,_ := u.FindById(ctx, db, "id", helpers.IntToString(user.Id))

	return userRes, nil
}

func (u UserRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, user *models.User) (models.User, error) {
	return models.User{}, nil
}

func (u UserRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, user *models.User) (bool, error) {
	results := db.Where("id = ?", user.Id).Delete(&user)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND")
	}
	return true, nil
}

func (u UserRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (userRes models.User, err error) {
	results := db.Where(key+" = ?", value).First(&userRes)
	if results.RowsAffected < 1 {
		return userRes, errors.New("NOT_FOUND")
	}
	return userRes, nil
}

func (u UserRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (res []models.User, err error) {
	db.Find(&res)
	return res, nil
}

func (u UserRepositoryImpl) Login(ctx echo.Context, db *gorm.DB, user *models.User) (userRes models.User, err error) {
	db.Where("username = ?", user.Username).First(&userRes)
	if userRes.Username == "" {
		return userRes, errors.New("NOT_FOUND")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRes.Password), []byte(user.Password))
	if err != nil {
		return userRes, errors.New("NOT_FOUND")
	}

	return userRes, nil
}

