package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	GetAnalytics(ctx *gin.Context) 
}