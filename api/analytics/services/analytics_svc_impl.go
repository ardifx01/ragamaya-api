package services

import (
	"ragamaya-api/api/analytics/dto"
	"ragamaya-api/api/analytics/repositories"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
	}
}

func (s *CompServicesImpl) GetAnalytics(ctx *gin.Context) (*dto.AnalyticRes, *exceptions.Exception) {
	// Get order analytics
	orderAnalytics, err := s.getOrderAnalytics(ctx)
	if err != nil {
		return nil, err
	}

	// Get payout analytics
	payoutAnalytics, err := s.getPayoutAnalytics(ctx)
	if err != nil {
		return nil, err
	}

	// Get product analytics
	productAnalytics, err := s.getProductAnalytics(ctx)
	if err != nil {
		return nil, err
	}

	// Get user analytics
	userAnalytics, err := s.getUserAnalytics(ctx)
	if err != nil {
		return nil, err
	}

	// Get revenue analytics
	revenueAnalytics, err := s.getRevenueAnalytics(ctx)
	if err != nil {
		return nil, err
	}

	// Get platform analytics
	platformAnalytics, err := s.getPlatformAnalytics(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.AnalyticRes{
		AnalyticOrder:    *orderAnalytics,
		AnalyticPayout:   *payoutAnalytics,
		AnalyticProduct:  *productAnalytics,
		AnalyticUser:     *userAnalytics,
		AnalyticRevenue:  *revenueAnalytics,
		AnalyticPlatform: *platformAnalytics,
	}, nil
}

func (s *CompServicesImpl) getOrderAnalytics(ctx *gin.Context) (*dto.AnalyticOrderRes, *exceptions.Exception) {
	// Get total orders
	totalOrder, err := s.repo.FindTotalOrders(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get total successful orders
	totalOrderSuccess, err := s.repo.FindTotalOrdersByStatus(ctx, s.DB, "settlement")
	if err != nil {
		return nil, err
	}

	// Get total failed orders
	totalOrderFailed, err := s.repo.FindTotalOrdersByStatus(ctx, s.DB, "expire")
	if err != nil {
		return nil, err
	}

	// Get monthly orders for the last 12 months
	monthlyOrders, err := s.repo.FindMonthlyOrders(ctx, s.DB, "")
	if err != nil {
		return nil, err
	}

	monthlyOrdersSuccess, err := s.repo.FindMonthlyOrders(ctx, s.DB, "settlement")
	if err != nil {
		return nil, err
	}

	monthlyOrdersFailed, err := s.repo.FindMonthlyOrders(ctx, s.DB, "expire")
	if err != nil {
		return nil, err
	}

	return &dto.AnalyticOrderRes{
		TotalOrder:          totalOrder,
		TotalOrderSuccess:   totalOrderSuccess,
		TotalOrderFailed:    totalOrderFailed,
		MonthlyOrder:        monthlyOrders,
		MonthlyOrderSuccess: monthlyOrdersSuccess,
		MonthlyOrderFailed:  monthlyOrdersFailed,
	}, nil
}

func (s *CompServicesImpl) getPayoutAnalytics(ctx *gin.Context) (*dto.AnalyticPayoutRes, *exceptions.Exception) {
	// Get total payouts
	totalPayout, err := s.repo.FindTotalPayouts(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get total successful payouts
	totalPayoutSuccess, err := s.repo.FindTotalPayoutsByStatus(ctx, s.DB, string(models.Completed))
	if err != nil {
		return nil, err
	}

	// Get total failed payouts
	totalPayoutFailed, err := s.repo.FindTotalPayoutsByStatus(ctx, s.DB, string(models.Failed))
	if err != nil {
		return nil, err
	}

	// Get monthly payouts for the last 12 months
	monthlyPayouts, err := s.repo.FindMonthlyPayouts(ctx, s.DB, "")
	if err != nil {
		return nil, err
	}

	monthlyPayoutsSuccess, err := s.repo.FindMonthlyPayouts(ctx, s.DB, string(models.Completed))
	if err != nil {
		return nil, err
	}

	monthlyPayoutsFailed, err := s.repo.FindMonthlyPayouts(ctx, s.DB, string(models.Failed))
	if err != nil {
		return nil, err
	}

	return &dto.AnalyticPayoutRes{
		TotalPayout:          totalPayout,
		TotalPayoutSuccess:   totalPayoutSuccess,
		TotalPayoutFailed:    totalPayoutFailed,
		MonthlyPayout:        monthlyPayouts,
		MonthlyPayoutSuccess: monthlyPayoutsSuccess,
		MonthlyPayoutFailed:  monthlyPayoutsFailed,
	}, nil
}

func (s *CompServicesImpl) getProductAnalytics(ctx *gin.Context) (*dto.AnalyticProductRes, *exceptions.Exception) {
	// Get total products
	totalProducts, err := s.repo.FindTotalProducts(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get total products by type
	totalDigital, err := s.repo.FindTotalProductsByType(ctx, s.DB, "digital")
	if err != nil {
		return nil, err
	}

	totalPhysical, err := s.repo.FindTotalProductsByType(ctx, s.DB, "physical")
	if err != nil {
		return nil, err
	}

	// Get top selling products
	topSellingProducts, err := s.repo.FindTopSellingProducts(ctx, s.DB, 10) // top 10 products
	if err != nil {
		return nil, err
	}

	// Get monthly new products
	monthlyNewProducts, err := s.repo.FindMonthlyNewProducts(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	return &dto.AnalyticProductRes{
		TotalProducts:         totalProducts,
		TotalDigitalProducts:  totalDigital,
		TotalPhysicalProducts: totalPhysical,
		TopSellingProducts:    topSellingProducts,
		MonthlyNewProducts:    monthlyNewProducts,
	}, nil
}

func (s *CompServicesImpl) getUserAnalytics(ctx *gin.Context) (*dto.AnalyticUserRes, *exceptions.Exception) {
	// Get total users
	totalUsers, err := s.repo.FindTotalUsers(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get total sellers
	totalSellers, err := s.repo.FindTotalSellers(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get total verified users
	totalVerifiedUsers, err := s.repo.FindTotalVerifiedUsers(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get monthly new users
	monthlyNewUsers, err := s.repo.FindMonthlyNewUsers(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get monthly new sellers
	monthlyNewSellers, err := s.repo.FindMonthlyNewSellers(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	return &dto.AnalyticUserRes{
		TotalUsers:         totalUsers,
		TotalSellers:       totalSellers,
		TotalVerifiedUsers: totalVerifiedUsers,
		MonthlyNewUsers:    monthlyNewUsers,
		MonthlySellers:     monthlyNewSellers,
	}, nil
}

func (s *CompServicesImpl) getRevenueAnalytics(ctx *gin.Context) (*dto.AnalyticRevenueRes, *exceptions.Exception) {
	// Get total revenue
	totalRevenue, err := s.repo.FindTotalRevenue(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get monthly revenue
	monthlyRevenue, err := s.repo.FindMonthlyRevenue(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get average order value
	averageOrderValue, err := s.repo.FindAverageOrderValue(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get revenue by product type
	revenueByProductType, err := s.repo.FindRevenueByProductType(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	return &dto.AnalyticRevenueRes{
		TotalRevenue:     totalRevenue,
		MonthlyRevenue:   monthlyRevenue,
		AvgOrderValue:    averageOrderValue,
		RevenueByProduct: revenueByProductType,
	}, nil
}

func (s *CompServicesImpl) getPlatformAnalytics(ctx *gin.Context) (*dto.AnalyticPlatformRes, *exceptions.Exception) {
	// Get total quizzes
	totalQuizzes, err := s.repo.FindTotalQuizzes(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get total certificates
	totalCertificates, err := s.repo.FindTotalCertificates(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get monthly quiz taken
	monthlyQuizTaken, err := s.repo.FindMonthlyQuizTaken(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get monthly certificates
	monthlyCertificates, err := s.repo.FindMonthlyCertificates(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	// Get quiz completion rate
	quizCompletionRate, err := s.repo.FindQuizCompletionRate(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	return &dto.AnalyticPlatformRes{
		TotalQuizzes:         totalQuizzes,
		TotalCertificates:    totalCertificates,
		QuizCompletionRate:   quizCompletionRate,
		MonthlyQuizTaken:     monthlyQuizTaken,
		MonthlyCertificates:  monthlyCertificates,
	}, nil
}