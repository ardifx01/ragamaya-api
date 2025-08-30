package mapper

import (
	"ragamaya-api/api/users/dto"
	"ragamaya-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapUserMTO(input models.Users) dto.UserRes {
	var output dto.UserRes

	mapstructure.Decode(input, &output)
	return output
}
