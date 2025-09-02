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

func MapArticleITM(input dto.ArticleReq) models.Article {
	var output models.Article

	mapstructure.Decode(input, &output)
	return output
}

func MapArticleUTM(input dto.ArticleUpdateReq) models.Article {
	var output models.Article

	mapstructure.Decode(input, &output)
	return output
}

func MapArticleMTO(input models.Article) dto.ArticleRes {
	var output dto.ArticleRes

	mapstructure.Decode(input, &output)
	output.CreatedAt = input.CreatedAt
	return output
}
