// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	userControllers "ragamaya-api/api/users/controllers"
	userRepositories "ragamaya-api/api/users/repositories"
	userServices "ragamaya-api/api/users/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var userFeatureSet = wire.NewSet(
	userRepositories.NewComponentRepository,
	userServices.NewComponentServices,
	userControllers.NewCompController,
)

func InitializeUserController(db *gorm.DB, validate *validator.Validate) userControllers.CompControllers {
	wire.Build(userFeatureSet)
	return nil
}
