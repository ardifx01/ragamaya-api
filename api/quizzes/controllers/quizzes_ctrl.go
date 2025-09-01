package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Create(ctx *gin.Context)
	FindAllCategories(ctx *gin.Context)
	Search(ctx *gin.Context)
}
