package routers

import (
	"ragamaya-api/api/predicts/controllers"

	"github.com/gin-gonic/gin"
)

func PredictRoutes(router *gin.RouterGroup, predictController controllers.CompControllers) {
	predictGroup := router.Group("/predict")
	{
		predictGroup.POST("/analyze", predictController.Predict)
	}
}
