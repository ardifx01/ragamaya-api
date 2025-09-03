package routers

import (
	"ragamaya-api/api/analytics/controllers"

	"github.com/gin-gonic/gin"
)

func AnalyticRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	analyticGroup := r.Group("/analytic")
	{
		analyticGroup.GET("/getall", compControllers.GetAnalytics)
	}
}
