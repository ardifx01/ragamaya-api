package mapper

import (
	"ragamaya-api/api/predicts/dto"
	static "ragamaya-api/static/data"

	"github.com/go-viper/mapstructure/v2"
)

func MapBatikDTO(input static.BatikPattern) dto.PredictRes {
	var output dto.PredictRes

	mapstructure.Decode(input, &output)
	return output
}
