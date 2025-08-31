package routers

import (
	"net/http"
	"ragamaya-api/injectors"
	"ragamaya-api/pkg/helpers"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

func CompRouters(r *gin.RouterGroup, db *gorm.DB, storage *s3.Client, validate *validator.Validate, midtransCore *coreapi.Client) {
	r.GET("/health", func(ctx *gin.Context) {
		health := helpers.PerformHealthCheck(db)

		statusCode := http.StatusOK
		if health.Status == "unhealthy" {
			statusCode = http.StatusServiceUnavailable
		}

		ctx.JSON(statusCode, health)
	})

	userController := injectors.InitializeUserController(db, validate)
	storageController := injectors.InitializeStorageController(db, storage, validate)
	sellerController := injectors.InitializeSellerController(db, validate)
	productController := injectors.InitializeProductController(db, validate)
	orderController := injectors.InitializeOrderController(db, validate, midtransCore)
	paymentController := injectors.InitializePaymentController(db, validate)
	walletController := injectors.InitializeWalletController(db, validate)
	articleController := injectors.InitializeArticleController(db, validate)

	UserRoutes(r, userController)
	StorageRoutes(r, storageController)
	SellerRoutes(r, sellerController)
	ProductRoutes(r, productController)
	OrderRoutes(r, orderController)
	PaymentRoutes(r, paymentController)
	WalletRoutes(r, walletController)
	ArticleRoutes(r, articleController)
}
