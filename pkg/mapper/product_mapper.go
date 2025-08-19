package mapper

import (
	"ragamaya-api/api/products/dto"
	"ragamaya-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapProductITM(input dto.RegisterReq) models.Products {
	var output models.Products

	mapstructure.Decode(input, &output)
	return output
}

func MapProductMTO(input models.Products) dto.ProductRes {
	var output dto.ProductRes

	mapstructure.Decode(input, &output)
	return output
}
