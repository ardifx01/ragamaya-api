package routers

import (
	"ragamaya-api/injectors"
	internalInjectors "ragamaya-api/internal/injectors"
	"ragamaya-api/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InternalRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	internalController := internalInjectors.InitializeAuthController(validate)
	AuthRoutes(r, internalController)

	r.Use(middleware.InternalMiddleware())
	walletController := injectors.InitializeWalletController(db, validate)
	WalletRouter(r, walletController)
}
