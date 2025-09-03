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

	return &dto.AnalyticRes{
		AnalyticOrder:  *orderAnalytics,
		AnalyticPayout: *payoutAnalytics,
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
