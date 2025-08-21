package routers

import (
	"ragamaya-api/midtrans/injectors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

func MidtransRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate, midtransCore *coreapi.Client) {
	notificationController := injectors.InitializeNotificationController(db, validate, midtransCore)

	NotificationRoutes(r, notificationController)
}
