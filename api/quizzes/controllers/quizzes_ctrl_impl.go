package controllers

import (
	"net/http"
	"ragamaya-api/api/quizzes/dto"
	"ragamaya-api/api/quizzes/services"
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

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	var data dto.QuizReq
	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.Create(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "quiz created",
	})
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

func (h *CompControllersImpl) Search(ctx *gin.Context) {
	var data dto.SearchReq
	if err := ctx.ShouldBindQuery(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	result, err := h.services.Search(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "data retrieved successfully",
		Body:    result,
	})
}

func (h *CompControllersImpl) FindBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	data, err := h.services.FindBySlug(ctx, slug)
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

func (h *CompControllersImpl) Analyze(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	var data dto.AnalyzeReq
	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	result, err := h.services.Analyze(ctx, uuid, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "data retrieved successfully",
		Body:    result,
	})
}

func (h *CompControllersImpl) Update(ctx *gin.Context) {
	var data dto.QuizUpdateReq
	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	data.UUID = ctx.Param("uuid")
	err := h.services.Update(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "update success",
	})
}

func (h *CompControllersImpl) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.Delete(ctx, uuid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "delete success",
	})
}