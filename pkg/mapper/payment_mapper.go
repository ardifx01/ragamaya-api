package mapper

import (
	"ragamaya-api/api/payments/dto"
	"ragamaya-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapPaymentMTO(input models.Payments) dto.PaymentRes {
	var output dto.PaymentRes

	mapstructure.Decode(input, &output)
	output.CreatedAt = input.CreatedAt
	output.UpdatedAt = input.UpdatedAt
	
	return output
}
