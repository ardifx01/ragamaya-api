package controllers

import (
	"net/http"
	"ragamaya-api/api/sellers/dto"
	"ragamaya-api/api/sellers/services"
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

func (h *CompControllersImpl) Register(ctx *gin.Context) {
	var data dto.RegisterReq
	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "register success",
	})
}