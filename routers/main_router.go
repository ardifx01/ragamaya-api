package routers

import (
	"net/http"
	"ragamaya-api/injectors"
	"ragamaya-api/pkg/helpers"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CompRouters(r *gin.RouterGroup, db *gorm.DB, storage *s3.Client, validate *validator.Validate) {
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

	UserRoutes(r, userController)
	StorageRoutes(r, storageController)
}
