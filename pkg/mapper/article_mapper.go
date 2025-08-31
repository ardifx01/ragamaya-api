package mapper

import (
	"ragamaya-api/api/articles/dto"
	"ragamaya-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapCategoryMTO(input models.ArticleCategory) dto.CategoryRes {
	var output dto.CategoryRes

	mapstructure.Decode(input, &output)
	return output
}
