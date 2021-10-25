package services

import (

	"github.com/labstack/echo"
	"gorm.io/gorm"

	//"kalika-be/config"
	"kalika-be/helpers"
	"kalika-be/models/web"
	"kalika-be/repository"
)

type (
	GeneralSettingService interface {
		Create(ctx echo.Context) (res web.Response, err error)
		FindAll(ctx echo.Context) (web.Response, error)
	}

	GeneralSettingServiceImpl struct {
		GeneralSettingRepository repository.GeneralSettingRepository
		db                       *gorm.DB
	}
)

func NewGeneralSettingService(GeneralSettingRepository repository.GeneralSettingRepository, db *gorm.DB) GeneralSettingService {
	return &GeneralSettingServiceImpl{
		GeneralSettingRepository: GeneralSettingRepository,
		db:                       db,
	}
}

func (service *GeneralSettingServiceImpl) Create(ctx echo.Context) (res web.Response, err error) {
	params,_ := ctx.FormParams()

	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)
	
	for index, param := range params {
		service.GeneralSettingRepository.CreateOrUpdate(ctx, tx, index, param[0])
	}

	generalSettingRepo, err := service.GeneralSettingRepository.FindAll(ctx, tx)
	if err != nil {
		return helpers.Response(err.Error(), "", nil), err
	}

	return helpers.Response("CREATED", "Sukses Menyimpan Data", generalSettingRepo), err
}

func (service GeneralSettingServiceImpl) FindAll(ctx echo.Context) (res web.Response, err error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	cakeTypeRepo, err := service.GeneralSettingRepository.FindAll(ctx, tx)

	return helpers.Response("OK", "Sukses Mengambil Data", cakeTypeRepo), err
}
