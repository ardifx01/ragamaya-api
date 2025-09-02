package services

import (
	"ragamaya-api/api/predicts/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Predict(ctx *gin.Context, data dto.PredictReq) (*dto.PredictRes, *exceptions.Exception)
}