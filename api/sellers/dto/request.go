package dto

type RegisterReq struct {
	Name      string `json:"name" validate:"required"`
	Desc      string `json:"desc" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Whatsapp  string `json:"whatsapp" validate:"required,e164"`
	AvatarURL string `json:"avatar_url"`
}

type UpdateReq struct {
	Name      string `json:"name" validate:"required"`
	Desc      string `json:"desc" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Whatsapp  string `json:"whatsapp" validate:"required,e164"`
	AvatarURL string `json:"avatar_url"`
}

type OrderQueryParams struct {
	Status      string `form:"status" validate:"omitempty,oneof=pending success failed"`
	ProductUUID string `form:"product_uuid" validate:"omitempty,uuid4"`
}

type OrderRes struct {
	UUID        string `json:"uuid"`
	UserUUID    string `json:"user_uuid"`
	ProductUUID string `json:"product_uuid"`
	Quantity    int    `json:"quantity"`
	GrossAmt    int64  `json:"amount"`
	Status      string `json:"status"`

	Product ProductRes `json:"product"`
	User    UserRes    `json:"user"`
}

type ProductRes struct {
	UUID        string `json:"uuid"`
	SellerUUID  string `json:"seller_uuid"`
	ProductType string `json:"product_type"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       uint   `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	Keywords    string `json:"keywords" validate:"required"`

	Thumbnails []ProductThumbnailsRes `json:"thumbnails" validate:"required,dive"`
}

type ProductThumbnailsRes struct {
	ThumbnailURL string `json:"thumbnail_url" validate:"required,url"`
}

type UserRes struct {
	UUID            string `json:"uuid"`
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"is_email_verified"`
	SUB             string `json:"sub"`
	Name            string `json:"name"`
	AvatarURL       string `json:"avatar_url"`
}