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

	productControllers "ragamaya-api/api/products/controllers"
	productRepositories "ragamaya-api/api/products/repositories"
	productServices "ragamaya-api/api/products/services"

	orderControllers "ragamaya-api/api/orders/controllers"
	orderRepositories "ragamaya-api/api/orders/repositories"
	orderServices "ragamaya-api/api/orders/services"
	
	paymentControllers "ragamaya-api/api/payments/controllers"
	paymentRepositories "ragamaya-api/api/payments/repositories"
	paymentServices "ragamaya-api/api/payments/services"

	walletControllers "ragamaya-api/api/wallets/controllers"
	walletRepositories "ragamaya-api/api/wallets/repositories"
	walletServices "ragamaya-api/api/wallets/services"
	
	articleControllers "ragamaya-api/api/articles/controllers"
	articleRepositories "ragamaya-api/api/articles/repositories"
	articleServices "ragamaya-api/api/articles/services"
	
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

var userFeatureSet = wire.NewSet(
	userRepositories.NewComponentRepository,
	userServices.NewComponentServices,
	userControllers.NewCompController,
	
	walletRepositories.NewComponentRepository,
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

	userRepositories.NewComponentRepository,
)

var productFeatureSet = wire.NewSet(
	productRepositories.NewComponentRepository,
	productServices.NewComponentServices,
	productControllers.NewCompController,
)

var orderFeatureSet = wire.NewSet(
	orderRepositories.NewComponentRepository,
	orderServices.NewComponentServices,
	orderControllers.NewCompController,
	
	paymentRepositories.NewComponentRepository,
	productRepositories.NewComponentRepository,
)

var paymentFeatureSet = wire.NewSet(
	paymentRepositories.NewComponentRepository,
	paymentServices.NewComponentServices,
	paymentControllers.NewCompController,
)

var walletFeatureSet = wire.NewSet(
	walletRepositories.NewComponentRepository,
	walletServices.NewComponentServices,
	walletControllers.NewCompController,
)

var articleFeatureSet = wire.NewSet(
	articleRepositories.NewComponentRepository,
	articleServices.NewComponentServices,
	articleControllers.NewCompController,
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

func InitializeProductController(db *gorm.DB, validate *validator.Validate) productControllers.CompControllers {
	wire.Build(productFeatureSet)
	return nil
}

func InitializeOrderController(db *gorm.DB, validate *validator.Validate, midtransCore *coreapi.Client) orderControllers.CompControllers {
	wire.Build(orderFeatureSet)
	return nil
}

func InitializePaymentController(db *gorm.DB, validate *validator.Validate) paymentControllers.CompControllers {
	wire.Build(paymentFeatureSet)
	return nil
}

func InitializeWalletController(db *gorm.DB, validate *validator.Validate) walletControllers.CompControllers {
	wire.Build(walletFeatureSet)
	return nil
}

func InitializeArticleController(db *gorm.DB, validate *validator.Validate) articleControllers.CompControllers {
	wire.Build(articleFeatureSet)
	return nil
}