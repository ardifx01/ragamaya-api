package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	FindByUserUUID(ctx *gin.Context)
	FindTransactionHistoryByUserUUID(ctx *gin.Context)
	RequestPayout(ctx *gin.Context)
	ResponsePayout(ctx *gin.Context) 
}