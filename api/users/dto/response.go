package dto

type Response struct {
	Status  int         `json:"status" example:"200"`
	Message string      `json:"message" example:"success"`
	Body    interface{} `json:"body,omitempty"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Roles string

const (
	User   Roles = "user"
	Seller Roles = "seller"
)

type UserOutput struct {
	UUID            string `json:"uuid"`
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"is_email_verified"`
	SUB             string `json:"sub"`
	Name            string `json:"name"`
	Role            Roles  `json:"role"`
	AvatarURL       string `json:"avatar_url"`

	SellerProfile SellerRes `json:"seller_profile,omitempty"`
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
