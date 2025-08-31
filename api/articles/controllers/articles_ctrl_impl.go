package controllers

import (
	"net/http"
	"ragamaya-api/api/articles/dto"
	"ragamaya-api/api/articles/services"

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

func (h *CompControllersImpl) FindAllCategories(ctx *gin.Context) {
	data, err := h.services.FindAllCategories(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "data retrieved successfully",
		Body:    data,
	})
}