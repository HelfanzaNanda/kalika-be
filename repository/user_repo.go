package repository

import (
	"errors"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"
)

type (
	UserRepository interface{
		Create(ctx echo.Context, db *gorm.DB, user *domain.User) (domain.User, error)
		Update(ctx echo.Context, db *gorm.DB, user *domain.User) (domain.User, error)
		Delete(ctx echo.Context, db *gorm.DB, user *domain.User) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.User, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.User, error)
		Login(ctx echo.Context, db *gorm.DB, user *domain.User) (domain.User, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.UserDatatable, int64, int64, error)
	}

	UserRepositoryImpl struct {

	}
)

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, user *domain.User) (domain.User, error) {
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, errors.New("ERROR|User tidak ditemukan")
	}
	//fmt.Println(ctx.Get("userInfo").(map[string]interface{})["name"])
	user.Password = string(hashedPassword)
	db.Create(&user)
	userRes,_ := u.FindById(ctx, db, "id", helpers.IntToString(user.Id))

	return userRes, nil
}

func (u *UserRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, user *domain.User) (domain.User, error) {
	if user.Password != "" {
		password := []byte(user.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			return domain.User{}, errors.New("ERROR|Error Password Encryption")
		}
		user.Password = string(hashedPassword)
	}

	db.Where("id = ?", user.Id).Updates(&user)
	userRes,_ := u.FindById(ctx, db, "id", helpers.IntToString(user.Id))

	return userRes, nil
}

func (u *UserRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, user *domain.User) (bool, error) {
	results := db.Where("id = ?", user.Id).Delete(&user)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|User tidak ditemukan")
	}
	return true, nil
}

func (u *UserRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (userRes domain.User, err error) {
	results := db.Where(key+" = ?", value).First(&userRes)
	if results.RowsAffected < 1 {
		return userRes, errors.New("NOT_FOUND|User tidak ditemukan")
	}
	return userRes, nil
}

func (u *UserRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (res []domain.User, err error) {
	db.Find(&res)
	return res, nil
}

func (u *UserRepositoryImpl) Login(ctx echo.Context, db *gorm.DB, user *domain.User) (userRes domain.User, err error) {
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

func (u *UserRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (userRes []web.UserDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("users").Select(
		"users.id," +
		"users.name," +
		"users.username," +
		"users.role_id," +
		"roles.name as role_name," +
		"users.store_id," +
		"stores.name as store_name," +
		"users.created_at," +
		"users.updated_at").Joins("inner join roles on roles.id = users.role_id").Joins("inner join stores on stores.id = users.store_id")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(users.id = ? OR users.name LIKE ? OR users.username LIKE ?)", search, "%"+search+"%", "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("users.id desc")
	qry.Find(&userRes)
	return userRes, totalData, totalFiltered, nil
}