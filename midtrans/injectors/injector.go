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
	PaymentRepositories "ragamaya-api/api/payments/repositories"
	ProductRepositories "ragamaya-api/api/products/repositories"
	
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
	PaymentRepositories.NewComponentRepository,
	ProductRepositories.NewComponentRepository,
)

func InitializeNotificationController(db *gorm.DB, validate *validator.Validate, midtransCore *coreapi.Client) notificationControllers.CompControllers {
	wire.Build(notificationFeatureSet)
	return nil
}
