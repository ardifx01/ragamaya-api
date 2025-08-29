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

type AnalyticsRes struct {
	TotalProducts        int64               `json:"total_products"`
	TotalOrders          int64               `json:"total_orders"`
	TotalRevenue         int64               `json:"total_revenue"`
	TotalRevenueCurrency string              `json:"total_revenue_currency"`
	MonthlyRevenue       []MonthlyRevenueRes `json:"monthly_revenue"`
	MonthlyOrders        []MonthlyOrdersRes  `json:"monthly_orders"`
}

type MonthlyRevenueRes struct {
	Month    string `json:"month"`
	Revenue  int64  `json:"revenue"`
	Currency string `json:"currency"`
}

type MonthlyOrdersRes struct {
	Month       string `json:"month"`
	TotalOrders int64  `json:"total_orders"`
}
