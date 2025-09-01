package routers

import (
	"ragamaya-api/api/quizzes/controllers"

	"github.com/gin-gonic/gin"
)

func QuizRoutes(r *gin.RouterGroup, compControllers controllers.CompControllers) {
	quizGroup := r.Group("/quiz")
	{
		quizGroup.GET("/categories", compControllers.FindAllCategories)
		quizGroup.GET("/search", compControllers.Search)
	}
}
