package services

import (
	"ragamaya-api/api/storages/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.FilesInput) (*dto.FilesRes, *exceptions.Exception)
	FindAllImages(ctx *gin.Context) ([]dto.FilesRes, *exceptions.Exception)
}
