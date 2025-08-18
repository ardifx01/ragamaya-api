package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type SellerRes struct {
	UUID      string `json:"uuid"`
	UserUUID  string `json:"user_uuid"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Address   string `json:"address"`
	AvatarURL string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
}
