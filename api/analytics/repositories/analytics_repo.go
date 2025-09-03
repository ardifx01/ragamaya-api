package repositories

import (
	"ragamaya-api/api/analytics/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	// Order Analytics
	FindTotalOrders(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception)
	FindTotalOrdersByStatus(ctx *gin.Context, tx *gorm.DB, status string) (int64, *exceptions.Exception)
	FindMonthlyOrders(ctx *gin.Context, tx *gorm.DB, status string) ([]dto.MonthlyOrdersRes, *exceptions.Exception)

	// Payout Analytics
	FindTotalPayouts(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception)
	FindTotalPayoutsByStatus(ctx *gin.Context, tx *gorm.DB, status string) (int64, *exceptions.Exception)
	FindMonthlyPayouts(ctx *gin.Context, tx *gorm.DB, status string) ([]dto.MonthlyPayoutRes, *exceptions.Exception)

	// Product Analytics
	FindTotalProducts(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception)
	FindTotalProductsByType(ctx *gin.Context, tx *gorm.DB, productType string) (int64, *exceptions.Exception)
	FindTopSellingProducts(ctx *gin.Context, tx *gorm.DB, limit int) ([]dto.TopSellingProductRes, *exceptions.Exception)
	FindMonthlyNewProducts(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyCountRes, *exceptions.Exception)

	// User Analytics
	FindTotalUsers(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception)
	FindTotalSellers(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception)
	FindTotalVerifiedUsers(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception)
	FindMonthlyNewUsers(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyCountRes, *exceptions.Exception)
	FindMonthlyNewSellers(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyCountRes, *exceptions.Exception)

	// Revenue Analytics
	FindTotalRevenue(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception)
	FindMonthlyRevenue(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyAmountRes, *exceptions.Exception)
	FindAverageOrderValue(ctx *gin.Context, tx *gorm.DB) (float64, *exceptions.Exception)
	FindRevenueByProductType(ctx *gin.Context, tx *gorm.DB) ([]dto.RevenueByProductRes, *exceptions.Exception)

	// Platform Analytics
	FindTotalQuizzes(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception)
	FindTotalCertificates(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception)
	FindMonthlyQuizTaken(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyCountRes, *exceptions.Exception)
	FindMonthlyCertificates(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyCountRes, *exceptions.Exception)
	FindQuizCompletionRate(ctx *gin.Context, tx *gorm.DB) (float64, *exceptions.Exception)
}
