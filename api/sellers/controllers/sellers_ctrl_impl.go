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

	err := h.services.Register(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "register success",
	})
}

func (h *CompControllersImpl) FindByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	data, err := h.services.FindByUUID(ctx, uuid)
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

func (h *CompControllersImpl) Update(ctx *gin.Context) {
	var data dto.UpdateReq
	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

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
	err := h.services.Delete(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "delete success",
	})
}

func (h *CompControllersImpl) FindOrders(ctx *gin.Context) {
	var queryParams dto.OrderQueryParams
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewValidationException(err))
		return
	}

	data, err := h.services.FindOrders(ctx, queryParams)
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

func (h *CompControllersImpl) Analytics(ctx *gin.Context) {
	// to be implemented
}