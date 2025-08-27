// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	notificationControllers "ragamaya-api/midtrans/notifications/controllers"
	notificationServices "ragamaya-api/midtrans/notifications/services"
	
	orderServices "ragamaya-api/api/orders/services"
	orderRepositories "ragamaya-api/api/orders/repositories"
	paymentRepositories "ragamaya-api/api/payments/repositories"
	productRepositories "ragamaya-api/api/products/repositories"
	walletServices "ragamaya-api/api/wallets/services"
	walletRepositories "ragamaya-api/api/wallets/repositories"
	
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

var notificationFeatureSet = wire.NewSet(
	notificationServices.NewComponentServices,
	notificationControllers.NewCompController,

	orderServices.NewComponentServices,
	orderRepositories.NewComponentRepository,
	paymentRepositories.NewComponentRepository,
	productRepositories.NewComponentRepository,
	walletServices.NewComponentServices,
	walletRepositories.NewComponentRepository,
)

func InitializeNotificationController(db *gorm.DB, validate *validator.Validate, midtransCore *coreapi.Client) notificationControllers.CompControllers {
	wire.Build(notificationFeatureSet)
	return nil
}
