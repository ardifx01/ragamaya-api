package mapper

import (
	"ragamaya-api/api/orders/dto"
	"ragamaya-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapOrderITM(input dto.OrderReq) models.Orders {
	var output models.Orders

	mapstructure.Decode(input, &output)
	return output
}

func MapOrderMTO(input models.Orders) dto.OrderRes {
	var output dto.OrderRes

	mapstructure.Decode(input, &output)
	return output
}
