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