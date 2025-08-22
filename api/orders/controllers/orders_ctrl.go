package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Create(ctx *gin.Context)
	StreamInfo(ctx *gin.Context)
	FindByUUID(ctx *gin.Context)
	FindByUserUUID(ctx *gin.Context)
}
