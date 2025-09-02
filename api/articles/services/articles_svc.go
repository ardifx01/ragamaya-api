package services

import (
	"ragamaya-api/api/articles/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	FindAllCategories(ctx *gin.Context) ([]dto.CategoryRes, *exceptions.Exception)
	Create(ctx *gin.Context, data dto.ArticleReq) (*dto.ArticleRes, *exceptions.Exception)
	Search(ctx *gin.Context, data dto.SearchReq) ([]dto.ArticleRes, *exceptions.Exception)
	FindBySlug(ctx *gin.Context, slug string) (*dto.ArticleRes, *exceptions.Exception)
	Update(ctx *gin.Context, data dto.ArticleUpdateReq) (*dto.ArticleRes, *exceptions.Exception)
	Delete(ctx *gin.Context, uuid string) *exceptions.Exception
}
