package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Register(ctx *gin.Context)
	FindByUUID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Search(ctx *gin.Context)
	DeleteThumbnail(ctx *gin.Context)
}
