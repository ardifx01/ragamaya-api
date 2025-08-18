package dto

type RegisterReq struct {
	Name      string `json:"name" validate:"required"`
	Desc      string `json:"desc" validate:"required"`
	Address   string `json:"address" validate:"required"`
	AvatarURL string `json:"avatar_url"`
}
