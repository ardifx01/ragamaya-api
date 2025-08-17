// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	authControllers "ragamaya-api/internal/auth/controllers"
	authServices "ragamaya-api/internal/auth/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var authFeatureSet = wire.NewSet(
	authServices.NewComponentServices,
	authControllers.NewCompController,
)

func InitializeAuthController(validate *validator.Validate) authControllers.CompControllers {
	wire.Build(authFeatureSet)
	return nil
}
