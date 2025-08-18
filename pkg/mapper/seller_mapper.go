package mapper

import (
	"ragamaya-api/api/sellers/dto"
	"ragamaya-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapSellerITM(input dto.RegisterReq) models.Sellers {
	var output models.Sellers

	mapstructure.Decode(input, &output)
	return output
}

func MapSellerUTM(input dto.UpdateReq) models.Sellers {
	var output models.Sellers

	mapstructure.Decode(input, &output)
	return output
}

func MapSellerMTO(input models.Sellers) dto.SellerRes {
	var output dto.SellerRes

	mapstructure.Decode(input, &output)
	return output
}
