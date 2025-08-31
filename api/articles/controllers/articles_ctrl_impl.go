package controllers

import (
	"net/http"
	"ragamaya-api/api/articles/dto"
	"ragamaya-api/api/articles/services"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/logger"

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

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	var data dto.ArticleReq
	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		logger.Info(jsonErr.Error())
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	result, err := h.services.Create(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Body:    result,
		Message: "article created",
	})
}
