// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	userControllers "ragamaya-api/api/users/controllers"
	userRepositories "ragamaya-api/api/users/repositories"
	userServices "ragamaya-api/api/users/services"
	
	storageControllers "ragamaya-api/api/storages/controllers"
	storageRepositories "ragamaya-api/api/storages/repositories"
	storageServices "ragamaya-api/api/storages/services"
	
	sellerControllers "ragamaya-api/api/sellers/controllers"
	sellerRepositories "ragamaya-api/api/sellers/repositories"
	sellerServices "ragamaya-api/api/sellers/services"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var userFeatureSet = wire.NewSet(
	userRepositories.NewComponentRepository,
	userServices.NewComponentServices,
	userControllers.NewCompController,
)

var storageFeatureSet = wire.NewSet(
	storageRepositories.NewComponentRepository,
	storageServices.NewComponentServices,
	storageControllers.NewCompController,
)

var sellerFeatureSet = wire.NewSet(
	sellerRepositories.NewComponentRepository,
	sellerServices.NewComponentServices,
	sellerControllers.NewCompController,
)

func InitializeUserController(db *gorm.DB, validate *validator.Validate) userControllers.CompControllers {
	wire.Build(userFeatureSet)
	return nil
}

func InitializeStorageController(db *gorm.DB, s3client *s3.Client, validate *validator.Validate) storageControllers.CompControllers {
	wire.Build(storageFeatureSet)
	return nil
}

func InitializeSellerController(db *gorm.DB, validate *validator.Validate) sellerControllers.CompControllers {
	wire.Build(sellerFeatureSet)
	return nil
}
