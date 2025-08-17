package services

import (
	"ragamaya-api/internal/auth/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Login(ctx *gin.Context, data dto.Login) (*string, *exceptions.Exception)
}
