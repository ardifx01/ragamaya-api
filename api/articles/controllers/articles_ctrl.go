package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	FindAllCategories(ctx *gin.Context)
	Create(ctx *gin.Context)
}