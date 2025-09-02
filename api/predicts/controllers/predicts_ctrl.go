package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Predict(ctx *gin.Context)
}
