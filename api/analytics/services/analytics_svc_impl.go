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
	orderAnalytics, err := s.getOrderAnalytics(ctx)
	if err != nil {
		return nil, err
	}

	payoutAnalytics, err := s.getPayoutAnalytics(ctx)
	if err != nil {
		return nil, err
	}

	productAnalytics, err := s.getProductAnalytics(ctx)
	if err != nil {
		return nil, err
	}

	userAnalytics, err := s.getUserAnalytics(ctx)
	if err != nil {
		return nil, err
	}

	revenueAnalytics, err := s.getRevenueAnalytics(ctx)
	if err != nil {
		return nil, err
	}

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
	totalOrder, err := s.repo.FindTotalOrders(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	totalOrderSuccess, err := s.repo.FindTotalOrdersByStatus(ctx, s.DB, "settlement")
	if err != nil {
		return nil, err
	}

	totalOrderFailed, err := s.repo.FindTotalOrdersByStatus(ctx, s.DB, "expire")
	if err != nil {
		return nil, err
	}

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
	totalPayout, err := s.repo.FindTotalPayouts(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	totalPayoutSuccess, err := s.repo.FindTotalPayoutsByStatus(ctx, s.DB, string(models.Completed))
	if err != nil {
		return nil, err
	}

	totalPayoutFailed, err := s.repo.FindTotalPayoutsByStatus(ctx, s.DB, string(models.Failed))
	if err != nil {
		return nil, err
	}

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
	totalProducts, err := s.repo.FindTotalProducts(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	totalDigital, err := s.repo.FindTotalProductsByType(ctx, s.DB, "digital")
	if err != nil {
		return nil, err
	}

	totalPhysical, err := s.repo.FindTotalProductsByType(ctx, s.DB, "physical")
	if err != nil {
		return nil, err
	}

	topSellingProducts, err := s.repo.FindTopSellingProducts(ctx, s.DB, 10)
	if err != nil {
		return nil, err
	}

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
	totalUsers, err := s.repo.FindTotalUsers(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	totalSellers, err := s.repo.FindTotalSellers(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	totalVerifiedUsers, err := s.repo.FindTotalVerifiedUsers(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	monthlyNewUsers, err := s.repo.FindMonthlyNewUsers(ctx, s.DB)
	if err != nil {
		return nil, err
	}

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
	totalRevenue, err := s.repo.FindTotalRevenue(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	monthlyRevenue, err := s.repo.FindMonthlyRevenue(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	averageOrderValue, err := s.repo.FindAverageOrderValue(ctx, s.DB)
	if err != nil {
		return nil, err
	}

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
	totalQuizzes, err := s.repo.FindTotalQuizzes(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	totalCertificates, err := s.repo.FindTotalCertificates(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	monthlyQuizTaken, err := s.repo.FindMonthlyQuizTaken(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	monthlyCertificates, err := s.repo.FindMonthlyCertificates(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	quizCompletionRate, err := s.repo.FindQuizCompletionRate(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	return &dto.AnalyticPlatformRes{
		TotalQuizzes:        totalQuizzes,
		TotalCertificates:   totalCertificates,
		QuizCompletionRate:  quizCompletionRate,
		MonthlyQuizTaken:    monthlyQuizTaken,
		MonthlyCertificates: monthlyCertificates,
	}, nil
}
