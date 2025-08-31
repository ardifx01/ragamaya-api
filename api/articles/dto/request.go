package dto

type ArticleReq struct {
	Title     string `json:"title" validate:"required"`
	Thumbnail string `json:"thumbnail" validate:"required,url"`
	Content   string `json:"content" validate:"required"`
	Category  string `json:"category" validate:"required"`
}
