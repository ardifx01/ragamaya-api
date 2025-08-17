package mapper

import (
	"ragamaya-api/api/users/dto"
	"ragamaya-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapUserInputToModel(input dto.Users) models.Users {
	var user models.Users

	mapstructure.Decode(input, &user)
	return user
}
