package routers

import (
	"ragamaya-api/injectors"
	internalInjectors "ragamaya-api/internal/injectors"
	"ragamaya-api/pkg/middleware"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InternalRouters(r *gin.RouterGroup, db *gorm.DB, s3client *s3.Client, validate *validator.Validate) {
	internalController := internalInjectors.InitializeAuthController(validate)
	AuthRoutes(r, internalController)

	r.Use(middleware.InternalMiddleware())

	walletController := injectors.InitializeWalletController(db, validate)
	articleController := injectors.InitializeArticleController(db, validate)
	storageController := injectors.InitializeStorageController(db, s3client, validate)
	quizController := injectors.InitializeQuizController(db, s3client, validate)
	analyticController := injectors.InitializeAnalyticController(db, validate)

	WalletRouter(r, walletController)
	ArticleRoutes(r, articleController)
	StorageRoutes(r, storageController)
	QuizRoutes(r, quizController)
	AnalyticRoutes(r, analyticController)
}
