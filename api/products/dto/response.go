package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
	Size    int         `json:"size,omitempty"`
}

type ProductRes struct {
	UUID        string      `json:"uuid"`
	SellerUUID  string      `json:"seller_uuid"`
	ProductType ProductType `json:"product_type" validate:"required,oneof=digital physical"`
	Name        string      `json:"name" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Price       uint        `json:"price" validate:"required"`
	Stock       int         `json:"stock" validate:"required"`
	Keywords    string      `json:"keywords" validate:"required"`

	Thumbnails   []ProductThumbnails   `json:"thumbnails" validate:"required,dive"`
	DigitalFiles []ProductDigitalFiles `json:"digital_files" validate:"required,dive"`
}

type ProductThumbnailsRes struct {
	ThumbnailURL string `json:"thumbnail_url" validate:"required,url"`
}

type ProductDigitalFilesRes struct {
	UUID        string `json:"uuid"`
	FileURL     string `json:"file_url" validate:"required,url"`
	Description string `json:"description" validate:"required"`
	Extension   string `json:"extension" validate:"required"`
}
