package dto

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type CategoryRes struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type ArticleRes struct {
	UUID      string    `json:"uuid"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Thumbnail string    `json:"thumbnail"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`

	Category *CategoryRes `json:"category"`
}
