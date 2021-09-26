package repository

import (
	"errors"
	"kalika-be/helpers"
	"kalika-be/models/domain"
	"kalika-be/models/web"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	SellerRepository interface{
		Create(ctx echo.Context, db *gorm.DB, seller *domain.Seller) (domain.Seller, error)
		Update(ctx echo.Context, db *gorm.DB, seller *domain.Seller) (domain.Seller, error)
		Delete(ctx echo.Context, db *gorm.DB, seller *domain.Seller) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.Seller, error)
		FindAll(ctx echo.Context, db *gorm.DB) ([]domain.Seller, error)
		Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) ([]web.SellerDatatable, int64, int64, error)
	}

	SellerRepositoryImpl struct {

	}
)

func NewSellerRepository() SellerRepository {
	return &SellerRepositoryImpl{}
}

func (repository SellerRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, seller *domain.Seller) (domain.Seller, error) {
	db.Create(&seller)
	sellerRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(seller.Id))
	return sellerRes, nil
}

func (repository SellerRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, seller *domain.Seller) (domain.Seller, error) {
	db.Where("id = ?", seller.Id).Updates(&seller)
	sellerRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(seller.Id))
	return sellerRes, nil
}

func (repository SellerRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, seller *domain.Seller) (bool, error) {
	results := db.Where("id = ?", seller.Id).Delete(&seller)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|seller tidak ditemukan")
	}
	return true, nil
}

func (repository SellerRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (sellerRes domain.Seller, err error) {
	results := db.Where(key+" = ?", value).First(&sellerRes)
	if results.RowsAffected < 1 {
		return sellerRes, errors.New("NOT_FOUND|seller tidak ditemukan")
	}
	return sellerRes, nil
}

func (repository SellerRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB) (sellerRes []domain.Seller, err error) {
	db.Find(&sellerRes)
	return sellerRes, nil
}

func (repository SellerRepositoryImpl) Datatable(ctx echo.Context, db *gorm.DB, draw string, limit string, start string, search string) (datatableRes []web.SellerDatatable, totalData int64, totalFiltered int64, err error) {
	qry := db.Table("sellers")
	qry.Count(&totalData)
	if search != "" {
		qry.Where("(id = ? OR name LIKE ?)", search, "%"+search+"%")
	}
	qry.Count(&totalFiltered)
	if helpers.StringToInt(limit) > 0 {
		qry.Limit(helpers.StringToInt(limit)).Offset(helpers.StringToInt(start))
	}
	qry.Order("id desc")
	qry.Find(&datatableRes)
	return datatableRes, totalData, totalFiltered, nil
}