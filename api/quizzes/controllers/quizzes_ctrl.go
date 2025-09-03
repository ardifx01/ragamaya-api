package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Create(ctx *gin.Context)
	FindAllCategories(ctx *gin.Context)
	Search(ctx *gin.Context)
	FindBySlug(ctx *gin.Context)
	Analyze(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
