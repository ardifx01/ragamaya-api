package services

import (
	"ragamaya-api/api/quizzes/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.QuizReq) *exceptions.Exception
	FindAllCategories(ctx *gin.Context) ([]dto.CategoryRes, *exceptions.Exception)
}