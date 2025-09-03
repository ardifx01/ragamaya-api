package services

import (
	"ragamaya-api/api/analytics/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	GetAnalytics(ctx *gin.Context) (*dto.AnalyticRes, *exceptions.Exception)
}
