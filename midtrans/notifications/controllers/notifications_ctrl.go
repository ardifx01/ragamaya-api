package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Payment(ctx *gin.Context)
}