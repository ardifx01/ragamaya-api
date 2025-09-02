package controllers

import (
	"net/http"
	"ragamaya-api/api/predicts/dto"
	"ragamaya-api/api/predicts/services"
	"ragamaya-api/pkg/exceptions"
	"strings"

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

func (c *CompControllersImpl) Predict(ctx *gin.Context) {
	form, exc := ctx.MultipartForm()
	if exc != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	files := form.File["file"]
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, "No files uploaded"))
		return
	}

	if files[0].Size > (10 * 1024 * 1024) {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, "File size exceeds 10MB"))
		return
	}

	mimeType := files[0].Header.Get("Content-Type")
	if !strings.HasPrefix(mimeType, "image/") {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, "Only image files are allowed"))
		return
	}

	fileContent, exc := files[0].Open()
	if exc != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, "Error reading file"))
		return
	}
	defer fileContent.Close()

	buffer := make([]byte, files[0].Size)
	_, exc = fileContent.Read(buffer)
	if exc != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, "Error reading file"))
		return
	}

	result, err := c.services.Predict(ctx, dto.PredictReq{
		File: buffer,
	})
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
