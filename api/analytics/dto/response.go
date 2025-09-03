package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type AnalyticRes struct {
	AnalyticOrder    AnalyticOrderRes    `json:"analytic_order"`
	AnalyticPayout   AnalyticPayoutRes   `json:"analytic_payout"`
	AnalyticProduct  AnalyticProductRes  `json:"analytic_product"`
	AnalyticUser     AnalyticUserRes     `json:"analytic_user"`
	AnalyticRevenue  AnalyticRevenueRes  `json:"analytic_revenue"`
	AnalyticPlatform AnalyticPlatformRes `json:"analytic_platform"`
}

type AnalyticProductRes struct {
	TotalProducts         int64                  `json:"total_products"`
	TotalDigitalProducts  int64                  `json:"total_digital_products"`
	TotalPhysicalProducts int64                  `json:"total_physical_products"`
	TopSellingProducts    []TopSellingProductRes `json:"top_selling_products"`
	MonthlyNewProducts    []MonthlyCountRes      `json:"monthly_new_products"`
}

type TopSellingProductRes struct {
	ProductUUID string `json:"product_uuid"`
	Name        string `json:"name"`
	TotalSold   int64  `json:"total_sold"`
	Revenue     int64  `json:"revenue"`
}

type AnalyticUserRes struct {
	TotalUsers         int64             `json:"total_users"`
	TotalSellers       int64             `json:"total_sellers"`
	MonthlyNewUsers    []MonthlyCountRes `json:"monthly_new_users"`
	MonthlySellers     []MonthlyCountRes `json:"monthly_new_sellers"`
	TotalVerifiedUsers int64             `json:"total_verified_users"`
}

type AnalyticRevenueRes struct {
	TotalRevenue     int64                 `json:"total_revenue"`
	MonthlyRevenue   []MonthlyAmountRes    `json:"monthly_revenue"`
	AvgOrderValue    float64               `json:"average_order_value"`
	RevenueByProduct []RevenueByProductRes `json:"revenue_by_product"`
}

type RevenueByProductRes struct {
	ProductType string `json:"product_type"`
	Revenue     int64  `json:"revenue"`
}

type AnalyticPlatformRes struct {
	TotalQuizzes        int64             `json:"total_quizzes"`
	TotalCertificates   int64             `json:"total_certificates"`
	MonthlyQuizTaken    []MonthlyCountRes `json:"monthly_quiz_taken"`
	MonthlyCertificates []MonthlyCountRes `json:"monthly_certificates"`
	QuizCompletionRate  float64           `json:"quiz_completion_rate"`
}

type MonthlyCountRes struct {
	Month string `json:"month"`
	Total int64  `json:"total"`
}

type MonthlyAmountRes struct {
	Month  string `json:"month"`
	Amount int64  `json:"amount"`
}

type AnalyticOrderRes struct {
	TotalOrder        int64 `json:"total_order"`
	TotalOrderSuccess int64 `json:"total_order_success"`
	TotalOrderFailed  int64 `json:"total_order_failed"`

	MonthlyOrder        []MonthlyOrdersRes `json:"monthly_orders"`
	MonthlyOrderSuccess []MonthlyOrdersRes `json:"monthly_order_success"`
	MonthlyOrderFailed  []MonthlyOrdersRes `json:"monthly_order_failed"`
}

type MonthlyOrdersRes struct {
	Month       string `json:"month"`
	TotalOrders int64  `json:"total_orders"`
}

type AnalyticPayoutRes struct {
	TotalPayout        int64 `json:"total_payout"`
	TotalPayoutSuccess int64 `json:"total_payout_success"`
	TotalPayoutFailed  int64 `json:"total_payout_failed"`

	MonthlyPayout        []MonthlyPayoutRes `json:"monthly_payout"`
	MonthlyPayoutSuccess []MonthlyPayoutRes `json:"monthly_payout_success"`
	MonthlyPayoutFailed  []MonthlyPayoutRes `json:"monthly_payout_failed"`
}

type MonthlyPayoutRes struct {
	Month       string `json:"month"`
	TotalPayout int64  `json:"total_payouts"`
}
