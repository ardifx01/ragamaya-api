package services

import (
	"ragamaya-api/api/users/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Login(ctx *gin.Context, data dto.LoginRequest) (*dto.TokenResponse, *exceptions.Exception) 
	RefreshToken(ctx *gin.Context, refreshToken string) (accessToken string, err *exceptions.Exception)
	Logout(ctx *gin.Context, accessToken, refreshToken string) *exceptions.Exception
}
