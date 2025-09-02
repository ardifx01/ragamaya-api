package routers

import (
	"ragamaya-api/api/quizzes/controllers"
	"ragamaya-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func QuizRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	quizGroup := r.Group("/quiz")
	{
		quizGroup.GET("/categories", compControllers.FindAllCategories)
		quizGroup.GET("/search", compControllers.Search)
		quizGroup.GET("/:slug", compControllers.FindBySlug)
		quizGroup.POST("/analyze/:uuid", middleware.AuthMiddleware(), compControllers.Analyze)
	}
}
