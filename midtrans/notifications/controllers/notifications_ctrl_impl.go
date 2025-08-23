package controllers

import (
	"crypto/hmac"
	"net/http"
	"os"
	"ragamaya-api/midtrans/notifications/dto"
	"ragamaya-api/midtrans/notifications/services"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"

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

func (h *CompControllersImpl) Payment(ctx *gin.Context) {
	var data dto.PaymentNotificationReq

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	dataString := data.OrderId + data.StatusCode + data.GrossAmount + serverKey
	calculatedSignature := helpers.EncryptToSHA512(dataString)

	if !hmac.Equal([]byte(calculatedSignature), []byte(data.SignatureKey)) {
		ctx.JSON(http.StatusUnauthorized, exceptions.NewException(http.StatusUnauthorized, exceptions.ErrUnauthorized))
		return
	}

	err := h.services.Payment(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "ok",
	})
}
