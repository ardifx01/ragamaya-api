package dto

type ProductType string

const (
	Digital  ProductType = "digital"
	Physical ProductType = "physical"
)

type RegisterReq struct {
	ProductType ProductType `json:"product_type" validate:"required,oneof=digital physical"`
	Name        string      `json:"name" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Price       uint        `json:"price" validate:"required"`
	Stock       int         `json:"stock" validate:"required"`
	Keywords    string      `json:"keywords" validate:"required"`

	Thumbnails   []ProductThumbnails   `json:"thumbnails" validate:"required,dive"`
	DigitalFiles []ProductDigitalFiles `json:"digital_files" validate:"required,dive"`
}

type ProductThumbnails struct {
	ThumbnailURL string `json:"thumbnail_url" validate:"required,url"`
}

type ProductDigitalFiles struct {
	FileURL     string `json:"file_url" validate:"required,url"`
	Description string `json:"description" validate:"required"`
	Extension   string `json:"extension" validate:"required"`
}
