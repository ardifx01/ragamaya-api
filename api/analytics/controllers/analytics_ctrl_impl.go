package controllers

import (
	"net/http"
	"ragamaya-api/api/analytics/dto"
	"ragamaya-api/api/analytics/services"

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

func (h *CompControllersImpl) GetAnalytics(ctx *gin.Context) {
	data, err := h.services.GetAnalytics(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Body:    data,
		Message: "success",
	})
}
