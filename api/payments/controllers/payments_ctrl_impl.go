package controllers

import (
	"net/http"
	"ragamaya-api/api/payments/dto"
	"ragamaya-api/api/payments/services"

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

func (h *CompControllersImpl) FindByUserUUID(ctx *gin.Context) {
	data, err := h.services.FindByUserUUID(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Body:    data,
		Message: "payment found",
	})
}

func (h *CompControllersImpl) FindPendingByUserUUID(ctx *gin.Context) {
	data, err := h.services.FindPendingByUserUUID(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Body:    data,
		Message: "payment found",
	})
}

func (h *CompControllersImpl) FindFailedByUserUUID(ctx *gin.Context) {
	data, err := h.services.FindFailedByUserUUID(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Body:    data,
		Message: "payment found",
	})
}

func (h *CompControllersImpl) FindSuccessByUserUUID(ctx *gin.Context) {
	data, err := h.services.FindSuccessByUserUUID(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Body:    data,
		Message: "payment found",
	})
}

func (h *CompControllersImpl) FindRefundByUserUUID(ctx *gin.Context) {
	data, err := h.services.FindRefundByUserUUID(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Body:    data,
		Message: "payment found",
	})
}

func (h *CompControllersImpl) FindSuccessByProductUUID(ctx *gin.Context) {
	data, err := h.services.FindSuccessByProductUUID(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Body:    data,
		Message: "payment found",
	})
}