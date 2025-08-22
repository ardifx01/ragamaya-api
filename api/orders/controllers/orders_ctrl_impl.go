package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ragamaya-api/api/orders/dto"
	"ragamaya-api/api/orders/services"
	"ragamaya-api/pkg/exceptions"
	"time"

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
	var data dto.OrderReq
	data.PaymentType = dto.OrderPaymentType(ctx.Param("payment_type"))

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
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
		Message: "order created",
	})
}

func (h *CompControllersImpl) StreamInfo(ctx *gin.Context) {
	orderUUID := ctx.Query("id")

	if orderUUID == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	ctx.Writer.Header().Set("X-Accel-Buffering", "no")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Flush()

	flusher, ok := ctx.Writer.(http.Flusher)
	if !ok {
		http.Error(ctx.Writer, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	client := dto.StreamClient{Writer: ctx.Writer, Flusher: flusher}

	dto.ClientsMutex.Lock()
	dto.Clients[orderUUID] = append(dto.Clients[orderUUID], client)
	dto.ClientsMutex.Unlock()

	defer h.services.RemoveStreamClient(ctx, orderUUID, client)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Request.Context().Done():
			return
		case t := <-ticker.C:
			payload := dto.OrderStreamRes{
				Type:    "ping",
				Message: t.Format(time.RFC3339),
			}
			jsonData, _ := json.Marshal(payload)
			data := fmt.Sprintf("data:%s\n\n", jsonData)
			_, err := fmt.Fprintf(ctx.Writer, data)
			if err != nil {
				log.Printf("Error writing to client: %v", err)
				return
			}
			flusher.Flush()
		}
	}
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

func (h *CompControllersImpl) FindByUserUUID(ctx *gin.Context) {
	data, err := h.services.FindByUserUUID(ctx)
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
