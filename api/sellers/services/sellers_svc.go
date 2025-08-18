package services

import (
	"ragamaya-api/api/sellers/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Register(ctx *gin.Context, data dto.RegisterReq) *exceptions.Exception
}