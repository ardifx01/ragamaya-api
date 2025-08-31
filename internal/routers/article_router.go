package routers

import (
	"ragamaya-api/api/articles/controllers"

	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	articleGroup := r.Group("/article")
	{
		articleGroup.GET("/categories", compControllers.FindAllCategories)
		articleGroup.GET("/search", compControllers.Search)
		articleGroup.POST("/create", compControllers.Create)
	}
}