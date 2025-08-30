package controllers

import (
	"net/http"
	"ragamaya-api/api/products/dto"
	"ragamaya-api/api/products/services"
	"ragamaya-api/pkg/exceptions"
	"strconv"

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
	uuid := ctx.Param("uuid")

	var data dto.ProductUpdateReq
	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.Update(ctx, uuid, data)
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

func (h *CompControllersImpl) Search(ctx *gin.Context) {
	var data dto.ProductSearchReq
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
		Size:    len(result),
	})
}

func (h *CompControllersImpl) DeleteThumbnail(ctx *gin.Context) {
	productUUID := ctx.Query("product_uuid")
	id, exc := strconv.ParseUint(ctx.Query("id"), 10, 64)
	if exc != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}
	
	if productUUID == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.DeleteThumbnail(ctx, productUUID, uint(id))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "thumbnail deleted successfully",
	})
}