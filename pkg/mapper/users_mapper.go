package mapper

import (
	"ragamaya-api/api/users/dto"
	"ragamaya-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapUserMTO(input models.Users) dto.UserOutput {
	var user dto.UserOutput

	mapstructure.Decode(input, &user)
	return user
}
