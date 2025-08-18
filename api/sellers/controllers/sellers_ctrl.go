package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Register(ctx *gin.Context)
}