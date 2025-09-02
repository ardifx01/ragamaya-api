package services

import (
	"bytes"
	"net/http"
	"ragamaya-api/api/predicts/dto"
	"ragamaya-api/api/predicts/repositories"
	"ragamaya-api/pkg/config"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/logger"
	"ragamaya-api/pkg/mapper"
	static "ragamaya-api/static/data"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
	}
}

func (s *CompServicesImpl) CallMLService(ctx *gin.Context, data dto.PredictReq) (*dto.MLRes, *exceptions.Exception) {
	client := resty.New()
	var response dto.MLRes

	resp, err := client.R().
		SetFileReader("image", "upload.jpg", bytes.NewReader(data.File)).
		SetResult(&response).
		Post(config.GetMLServiceBaseURL() + "/predict")

	if err != nil {
		logger.Error("Request error: %v", err.Error())
		return nil, exceptions.NewException(http.StatusBadGateway, exceptions.ErrInternalServer)
	}

	if resp.IsError() {
		logger.Error("HTTP error: %v", resp.String())
		return nil, exceptions.NewException(http.StatusBadGateway, exceptions.ErrInternalServer)
	}

	logger.Info("Success: %v", resp.String())
	return &response, nil
}

func (s *CompServicesImpl) Predict(ctx *gin.Context, data dto.PredictReq) (*dto.PredictRes, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	predict, err := s.CallMLService(ctx, data)
	if err != nil {
		return nil, err
	}

	detail, ok := static.GetBatikPattern(predict.Detected)
	if !ok {
		return nil, exceptions.NewException(http.StatusBadRequest, "Detail not found")
	}

	result := mapper.MapBatikDTO(*detail)
	result.Score = predict.Alternatives[0].Score
	result.Alternative = predict.Alternatives
	result.Match = predict.Alternatives[0].Match

	return &result, nil
}
