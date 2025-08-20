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

type ProductSearchReq struct {
	Keyword     *string      `form:"keyword" validate:"omitempty"`
	PriceMin    *uint        `form:"price_min" validate:"omitempty,gte=1"`
	PriceMax    *uint        `form:"price_max" validate:"omitempty,gte=1"`
	ProductType *ProductType `form:"product_type" validate:"omitempty,oneof=digital physical"`
	Page        *int         `form:"page" validate:"omitempty,gte=1"`
	PageSize    *int         `form:"page_size" validate:"omitempty,gte=1,lte=100"`
	SellerUUID  *string      `form:"seller_uuid" validate:"omitempty,uuid4"`
}
