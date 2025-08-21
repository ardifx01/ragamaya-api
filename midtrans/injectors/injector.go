// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	notificationControllers "ragamaya-api/midtrans/notifications/controllers"
	notificationServices "ragamaya-api/midtrans/notifications/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

var notificationFeatureSet = wire.NewSet(
	notificationServices.NewComponentServices,
	notificationControllers.NewCompController,
)

func InitializeNotificationController(db *gorm.DB, validate *validator.Validate, midtransCore *coreapi.Client) notificationControllers.CompControllers {
	wire.Build(notificationFeatureSet)
	return nil
}
