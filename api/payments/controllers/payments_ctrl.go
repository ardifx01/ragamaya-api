package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	FindByUserUUID(ctx *gin.Context)
	FindPendingByUserUUID(ctx *gin.Context)
	FindFailedByUserUUID(ctx *gin.Context)
	FindSuccessByUserUUID(ctx *gin.Context)
	FindRefundByUserUUID(ctx *gin.Context)
	FindSuccessByProductUUID(ctx *gin.Context)
}
	