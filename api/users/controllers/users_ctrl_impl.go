package controllers

import (
	"net/http"
	"ragamaya-api/api/users/dto"
	"ragamaya-api/api/users/services"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, "Invalid request body"))
		return
	}

	result, err := h.services.Login(ctx, req)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Body:    result,
		Message: "login success",
	})
}

func (h *CompControllersImpl) Refresh(ctx *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, "Invalid request body"))
		return
	}
	accessToken, err := h.services.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Body:    accessToken,
		Message: "refresh token success",
	})
}

func (h *CompControllersImpl) Logout(ctx *gin.Context) {
	var req dto.LogoutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, "Invalid request body"))
		return
	}
	err := h.services.Logout(ctx, req.AccessToken, req.RefreshToken)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "logout success",
	})
}
