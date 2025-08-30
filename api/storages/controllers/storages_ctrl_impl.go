package controllers

import (
	"net/http"
	"ragamaya-api/api/storages/dto"
	"ragamaya-api/api/storages/services"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
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

func (h *CompControllersImpl) Images(ctx *gin.Context) {
	form, exc := ctx.MultipartForm()
	if exc != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, "No files uploaded"))
		return
	}

	var uploadedFiles []interface{}

	for _, file := range files {
		fileResult := make(map[string]interface{})
		fileResult["file_name"] = file.Filename

		if file.Size > (10 * 1024 * 1024) {
			fileResult["status"] = "error"
			fileResult["message"] = "File size exceeds 10MB"
			uploadedFiles = append(uploadedFiles, fileResult)
			continue
		}

		mimeType := file.Header.Get("Content-Type")
		if !strings.HasPrefix(mimeType, "image/") {
			fileResult["status"] = "error"
			fileResult["message"] = "Only image files are allowed"
			uploadedFiles = append(uploadedFiles, fileResult)
			continue
		}

		fileContent, exc := file.Open()
		if exc != nil {
			fileResult["status"] = "error"
			fileResult["message"] = "Error reading file"
			uploadedFiles = append(uploadedFiles, fileResult)
			continue
		}
		defer fileContent.Close()

		buffer := make([]byte, file.Size)
		_, exc = fileContent.Read(buffer)
		if exc != nil {
			fileResult["status"] = "error"
			fileResult["message"] = "Error reading file"
			uploadedFiles = append(uploadedFiles, fileResult)
			continue
		}

		fileName := file.Filename
		fileExtension := fileName[strings.LastIndex(fileName, ".")+1:]
		mimeParts := strings.Split(mimeType, "/")
		mimeMainType, mimeSubType := mimeParts[0], ""
		if len(mimeParts) > 1 {
			mimeSubType = mimeParts[1]
		}

		fileData := dto.FilesInput{
			OriginalFileName: fileName,
			FileBuffer:       buffer,
			Size:             helpers.FormatFileSize(file.Size),
			Extension:        fileExtension,
			MimeType:         mimeMainType,
			MimeSubType:      mimeSubType,
			Meta:             "{}",
		}

		result, err := h.services.Create(ctx, fileData)
		if err != nil {
			fileResult["status"] = "error"
			fileResult["message"] = err.Message
			uploadedFiles = append(uploadedFiles, fileResult)
			continue
		}

		fileResult["status"] = "success"
		fileResult["result"] = result
		uploadedFiles = append(uploadedFiles, fileResult)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "File upload process completed",
		Body:    uploadedFiles,
	})
}

func (h *CompControllersImpl) Image(ctx *gin.Context) {
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

	fileName := files[0].Filename
	fileExtension := fileName[strings.LastIndex(fileName, ".")+1:]
	mimeParts := strings.Split(mimeType, "/")
	mimeMainType, mimeSubType := mimeParts[0], ""
	if len(mimeParts) > 1 {
		mimeSubType = mimeParts[1]
	}

	fileData := dto.FilesInput{
		OriginalFileName: fileName,
		FileBuffer:       buffer,
		Size:             helpers.FormatFileSize(files[0].Size),
		Extension:        fileExtension,
		MimeType:         mimeMainType,
		MimeSubType:      mimeSubType,
		Meta:             "{}",
	}

	result, err := h.services.Create(ctx, fileData)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "File upload process completed",
		Body:    result,
	})
}

func (h *CompControllersImpl) General(ctx *gin.Context) {
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

	fileName := files[0].Filename
	fileExtension := fileName[strings.LastIndex(fileName, ".")+1:]
	mimeType := files[0].Header.Get("Content-Type")
	mimeParts := strings.Split(mimeType, "/")
	mimeMainType, mimeSubType := mimeParts[0], ""
	if len(mimeParts) > 1 {
		mimeSubType = mimeParts[1]
	}

	fileData := dto.FilesInput{
		OriginalFileName: fileName,
		FileBuffer:       buffer,
		Size:             helpers.FormatFileSize(files[0].Size),
		Extension:        fileExtension,
		MimeType:         mimeMainType,
		MimeSubType:      mimeSubType,
		Meta:             "{}",
	}

	result, err := h.services.Create(ctx, fileData)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "File upload process completed",
		Body:    result,
	})
}

func (h *CompControllersImpl) FindAllImages(ctx *gin.Context) {
	data, err := h.services.FindAllImages(ctx)
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
