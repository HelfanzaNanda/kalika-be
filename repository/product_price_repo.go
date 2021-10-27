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
	ProductPriceRepository interface{
		Create(ctx echo.Context, db *gorm.DB, productPrice *web.ProductPosPost) (web.ProductPosPost, error)
		Update(ctx echo.Context, db *gorm.DB, productPrice *domain.ProductPrice) (domain.ProductPrice, error)
		Delete(ctx echo.Context, db *gorm.DB, productPrice *domain.ProductPrice) (bool, error)
		DeleteByProduct(ctx echo.Context, db *gorm.DB, productId int) (bool, error)
		FindById(ctx echo.Context, db *gorm.DB, key string, value string) (domain.ProductPrice, error)
		FindByProductId(ctx echo.Context, db *gorm.DB, productId int) ([]domain.ProductPrice, error)
		FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) ([]domain.ProductPrice, error)
	}

	ProductPriceRepositoryImpl struct {

	}
)

func NewProductPriceRepository() ProductPriceRepository {
	return &ProductPriceRepositoryImpl{}
}

func (repository *ProductPriceRepositoryImpl) Create(ctx echo.Context, db *gorm.DB, product *web.ProductPosPost) (res web.ProductPosPost, err error) {
	for _, val := range product.ProductPrices {
		val.ProductId = product.Id
		db.Create(&val)
		res.ProductPrices = append(res.ProductPrices, val)
	}
	return res, nil
}

func (repository *ProductPriceRepositoryImpl) Update(ctx echo.Context, db *gorm.DB, productPrice *domain.ProductPrice) (domain.ProductPrice, error) {
	db.Where("id = ?", productPrice.Id).Updates(&productPrice)
	productPriceRes,_ := repository.FindById(ctx, db, "id", helpers.IntToString(productPrice.Id))
	return productPriceRes, nil
}

func (repository *ProductPriceRepositoryImpl) Delete(ctx echo.Context, db *gorm.DB, productPrice *domain.ProductPrice) (bool, error) {
	results := db.Where("id = ?", productPrice.Id).Delete(&productPrice)
	if results.RowsAffected < 1 {
		return false, errors.New("NOT_FOUND|expenseDetail tidak ditemukan")
	}
	return true, nil
}

func (repository *ProductPriceRepositoryImpl) DeleteByProduct(ctx echo.Context, db *gorm.DB, productId int) (bool, error) {
	db.Where("product_id = ?", productId).Delete(domain.ProductPrice{})
	return true, nil
}

func (repository *ProductPriceRepositoryImpl) FindById(ctx echo.Context, db *gorm.DB, key string, value string) (productPriceRes domain.ProductPrice, err error) {
	results := db.Where(key+" = ?", value).First(&productPriceRes)
	if results.RowsAffected < 1 {
		return productPriceRes, errors.New("NOT_FOUND|expenseDetail tidak ditemukan")
	}
	return productPriceRes, nil
}

func (repository *ProductPriceRepositoryImpl) FindByProductId(ctx echo.Context, db *gorm.DB, productId int) (productPriceRes []domain.ProductPrice, err error) {
	results := db.Where("product_id = ?", productId).Find(&productPriceRes)
	if results.RowsAffected < 1 {
		return productPriceRes, errors.New("NOT_FOUND|productPrice tidak ditemukan")
	}
	return productPriceRes, nil
}

func (repository *ProductPriceRepositoryImpl) FindAll(ctx echo.Context, db *gorm.DB, params map[string][]string) (productRes []domain.ProductPrice, err error) {
	results := db.Table("product_prices")

	for k, v := range ctx.QueryParams() {
		if k == "name" {
			results = results.Where(k+" LIKE ?", "%"+v[0]+"%")
		} else if v[0] != "" && k != "id" {
			results = results.Where(k+" = ?", v[0])
		}
	}

	results.Find(&productRes)
	return productRes, nil
}

