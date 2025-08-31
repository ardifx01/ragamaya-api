package mapper

import (
	"ragamaya-api/api/quizzes/dto"
	"ragamaya-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapQuizCategoryMTO(input models.QuizCategory) dto.CategoryRes {
	var output dto.CategoryRes

	mapstructure.Decode(input, &output)
	return output
}

func MapQuizITM(input dto.QuizReq) models.Quiz {
	var data models.Quiz
	mapstructure.Decode(input, &data)
	return data
}

func MapQuizMTO(model models.Quiz) dto.QuizRes {
	var output dto.QuizRes
	mapstructure.Decode(model, &output)
	return output
}
