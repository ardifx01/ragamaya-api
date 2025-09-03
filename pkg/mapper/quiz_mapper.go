package mapper

import (
	"encoding/json"
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

func MapQuizUTM(input dto.QuizUpdateReq) models.Quiz {
	var data models.Quiz
	mapstructure.Decode(input, &data)
	return data
}

func MapQuizMTO(model models.Quiz) dto.QuizRes {
	var output dto.QuizRes
	mapstructure.Decode(model, &output)
	var questions []dto.QuizQuestionRes
	json.Unmarshal([]byte(model.Questions), &questions)
	output.TotalQuestions = len(questions)
	return output
}

func MapQuizMTDO(model models.Quiz) dto.QuizDetailRes {
	var output dto.QuizDetailRes
	mapstructure.Decode(model, &output)
	var questions []dto.QuizQuestionRes
	json.Unmarshal([]byte(model.Questions), &questions)
	output.Questions = questions
	output.TotalQuestions = len(questions)
	return output
}

func MapCertificateMTO(model models.QuizCertificate) dto.CertificateRes {
	var output dto.CertificateRes
	mapstructure.Decode(model, &output)
	output.CreatedAt = model.CreatedAt
	return output
}
