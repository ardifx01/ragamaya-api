package services

import (
	"ragamaya-api/api/articles/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	FindAllCategories(ctx *gin.Context) ([]dto.CategoryRes, *exceptions.Exception)
}