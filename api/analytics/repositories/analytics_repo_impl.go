package repositories

import (
	"ragamaya-api/api/analytics/dto"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

// Product Analytics Implementation
func (r *CompRepositoriesImpl) FindTotalProducts(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.Products{}).Count(&total).Error
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindTotalProductsByType(ctx *gin.Context, tx *gorm.DB, productType string) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.Products{}).Where("product_type = ?", productType).Count(&total).Error
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindTopSellingProducts(ctx *gin.Context, tx *gorm.DB, limit int) ([]dto.TopSellingProductRes, *exceptions.Exception) {
	var results []dto.TopSellingProductRes
	err := tx.Model(&models.Orders{}).
		Select("products.uuid as product_uuid, products.name, COUNT(*) as total_sold, SUM(orders.gross_amt) as revenue").
		Joins("JOIN products ON orders.product_uuid = products.uuid").
		Where("orders.status = ?", "success").
		Group("products.uuid, products.name").
		Order("total_sold DESC").
		Limit(limit).
		Find(&results).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return results, nil
}

func (r *CompRepositoriesImpl) FindMonthlyNewProducts(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyCountRes, *exceptions.Exception) {
	var results []dto.MonthlyCountRes
	query := `
		WITH months AS (
			SELECT TO_CHAR(date_trunc('month', NOW()) - INTERVAL '1 month' * gs, 'YYYY-MM') as month
			FROM generate_series(0, 11) gs
		)
		SELECT m.month as Month,
			COALESCE(COUNT(p.uuid), 0) as Total
		FROM months m
		LEFT JOIN products p 
			ON TO_CHAR(p.created_at, 'YYYY-MM') = m.month
		GROUP BY m.month
		ORDER BY m.month DESC;
	`
	if err := tx.Raw(query).Scan(&results).Error; err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return results, nil
}

// User Analytics Implementation
func (r *CompRepositoriesImpl) FindTotalUsers(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.Users{}).Count(&total).Error
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindTotalSellers(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.Users{}).Where("role = ?", models.Seller).Count(&total).Error
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindTotalVerifiedUsers(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.Users{}).Where("is_email_verified = ?", true).Count(&total).Error
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindMonthlyNewUsers(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyCountRes, *exceptions.Exception) {
	var results []dto.MonthlyCountRes
	query := `
		WITH months AS (
			SELECT TO_CHAR(date_trunc('month', NOW()) - INTERVAL '1 month' * gs, 'YYYY-MM') as month
			FROM generate_series(0, 11) gs
		)
		SELECT m.month as Month,
			COALESCE(COUNT(u.uuid), 0) as Total
		FROM months m
		LEFT JOIN users u 
			ON TO_CHAR(u.created_at, 'YYYY-MM') = m.month
		GROUP BY m.month
		ORDER BY m.month DESC;
	`
	if err := tx.Raw(query).Scan(&results).Error; err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return results, nil
}

func (r *CompRepositoriesImpl) FindMonthlyNewSellers(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyCountRes, *exceptions.Exception) {
	var results []dto.MonthlyCountRes
	query := `
		WITH months AS (
			SELECT TO_CHAR(date_trunc('month', NOW()) - INTERVAL '1 month' * gs, 'YYYY-MM') as month
			FROM generate_series(0, 11) gs
		)
		SELECT m.month as Month,
			COALESCE(COUNT(u.uuid), 0) as Total
		FROM months m
		LEFT JOIN users u 
			ON TO_CHAR(u.created_at, 'YYYY-MM') = m.month
			AND u.role = 'seller'
		GROUP BY m.month
		ORDER BY m.month DESC;
	`
	if err := tx.Raw(query).Scan(&results).Error; err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return results, nil
}

// Revenue Analytics Implementation
func (r *CompRepositoriesImpl) FindTotalRevenue(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.Orders{}).
		Where("status = ?", "settlement").
		Select("COALESCE(SUM(gross_amt), 0)").
		Row().
		Scan(&total)
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindMonthlyRevenue(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyAmountRes, *exceptions.Exception) {
	var results []dto.MonthlyAmountRes
	query := `
		WITH months AS (
			SELECT TO_CHAR(date_trunc('month', NOW()) - INTERVAL '1 month' * gs, 'YYYY-MM') as month
			FROM generate_series(0, 11) gs
		)
		SELECT m.month as Month,
			COALESCE(SUM(o.gross_amt), 0) as Amount
		FROM months m
		LEFT JOIN orders o 
			ON TO_CHAR(o.created_at, 'YYYY-MM') = m.month
			AND o.status = 'settlement'
		GROUP BY m.month
		ORDER BY m.month DESC;
	`
	if err := tx.Raw(query).Scan(&results).Error; err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return results, nil
}

func (r *CompRepositoriesImpl) FindAverageOrderValue(ctx *gin.Context, tx *gorm.DB) (float64, *exceptions.Exception) {
	var avg float64
	err := tx.Model(&models.Orders{}).
		Where("status = ?", "settlement").
		Select("COALESCE(AVG(gross_amt), 0)").
		Row().
		Scan(&avg)
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return avg, nil
}

func (r *CompRepositoriesImpl) FindRevenueByProductType(ctx *gin.Context, tx *gorm.DB) ([]dto.RevenueByProductRes, *exceptions.Exception) {
	var results []dto.RevenueByProductRes
	err := tx.Model(&models.Orders{}).
		Select("products.product_type, COALESCE(SUM(orders.gross_amt), 0) as revenue").
		Joins("JOIN products ON orders.product_uuid = products.uuid").
		Where("orders.status = ?", "settlement").
		Group("products.product_type").
		Find(&results).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return results, nil
}

// Platform Analytics Implementation
func (r *CompRepositoriesImpl) FindTotalQuizzes(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.Quiz{}).Count(&total).Error
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindTotalCertificates(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.QuizCertificate{}).Count(&total).Error
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindMonthlyQuizTaken(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyCountRes, *exceptions.Exception) {
	var results []dto.MonthlyCountRes
	query := `
		WITH months AS (
			SELECT TO_CHAR(date_trunc('month', NOW()) - INTERVAL '1 month' * gs, 'YYYY-MM') as month
			FROM generate_series(0, 11) gs
		)
		SELECT m.month as Month,
			COALESCE(COUNT(qa.id), 0) as Total
		FROM months m
		LEFT JOIN quiz_attempts qa 
			ON TO_CHAR(qa.created_at, 'YYYY-MM') = m.month
		GROUP BY m.month
		ORDER BY m.month DESC;
	`
	if err := tx.Raw(query).Scan(&results).Error; err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return results, nil
}

func (r *CompRepositoriesImpl) FindMonthlyCertificates(ctx *gin.Context, tx *gorm.DB) ([]dto.MonthlyCountRes, *exceptions.Exception) {
	var results []dto.MonthlyCountRes
	err := tx.Model(&models.QuizCertificate{}).
		Select("TO_CHAR(created_at, 'YYYY-MM') as Month, COUNT(*) as Total").
		Where("created_at >= ?", time.Now().AddDate(0, -12, 0)).
		Group("TO_CHAR(created_at, 'YYYY-MM')").
		Order("Month DESC").
		Find(&results).Error
	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}
	return results, nil
}

func (r *CompRepositoriesImpl) FindQuizCompletionRate(ctx *gin.Context, tx *gorm.DB) (float64, *exceptions.Exception) {
	// Get total number of quiz attempts
	var totalAttempts float64
	err := tx.Model(&models.QuizAttempt{}).
		Select("COUNT(*)").
		Row().
		Scan(&totalAttempts)
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}

	// Get total number of successful completions
	var successfulCompletions float64
	err = tx.Model(&models.QuizCertificate{}).
		Select("COUNT(*)").
		Row().
		Scan(&successfulCompletions)
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}

	if totalAttempts == 0 {
		return 0, nil
	}

	return (successfulCompletions / totalAttempts) * 100, nil
}

func (r *CompRepositoriesImpl) FindTotalOrders(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.Orders{}).Count(&total).Error
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindTotalOrdersByStatus(ctx *gin.Context, tx *gorm.DB, status string) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.Orders{}).Where("status = ?", status).Count(&total).Error
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindMonthlyOrders(ctx *gin.Context, tx *gorm.DB, status string) ([]dto.MonthlyOrdersRes, *exceptions.Exception) {
	var results []struct {
		Month       string
		TotalOrders int64
	}

	baseQuery := `
					WITH months AS (
						SELECT TO_CHAR(date_trunc('month', NOW()) - INTERVAL '1 month' * gs, 'YYYY-MM') as month
						FROM generate_series(0, 11) gs
					)
					SELECT 
						m.month AS Month,
						COALESCE(COUNT(o.uuid), 0) AS total_orders
					FROM months m
					LEFT JOIN orders o 
						ON TO_CHAR(o.created_at, 'YYYY-MM') = m.month
				`
	var query string
	if status != "" {
		query = baseQuery + " AND o.status = '" + status + "'"
	} else {
		query = baseQuery
	}

	query += `
		GROUP BY m.month
		ORDER BY m.month DESC;
	`

	err := tx.Raw(query).Scan(&results).Error

	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}

	monthlyOrders := make([]dto.MonthlyOrdersRes, len(results))
	for i, result := range results {
		monthlyOrders[i] = dto.MonthlyOrdersRes{
			Month:       result.Month,
			TotalOrders: result.TotalOrders,
		}
	}

	return monthlyOrders, nil
}

func (r *CompRepositoriesImpl) FindTotalPayouts(ctx *gin.Context, tx *gorm.DB) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.WalletPayoutRequest{}).Count(&total).Error
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindTotalPayoutsByStatus(ctx *gin.Context, tx *gorm.DB, status string) (int64, *exceptions.Exception) {
	var total int64
	err := tx.Model(&models.WalletPayoutRequest{}).Where("status = ?", status).Count(&total).Error
	if err != nil {
		return 0, exceptions.ParseGormError(tx, err)
	}
	return total, nil
}

func (r *CompRepositoriesImpl) FindMonthlyPayouts(ctx *gin.Context, tx *gorm.DB, status string) ([]dto.MonthlyPayoutRes, *exceptions.Exception) {
	var results []struct {
		Month       string
		TotalPayout int64
	}

	baseQuery := `
		WITH months AS (
			SELECT TO_CHAR(date_trunc('month', NOW()) - INTERVAL '1 month' * gs, 'YYYY-MM') as month
			FROM generate_series(0, 11) gs
		)
		SELECT m.month as Month,
			COALESCE(COUNT(wp.uuid), 0) as total_payout
		FROM months m
		LEFT JOIN wallet_payout_requests wp 
			ON TO_CHAR(wp.created_at, 'YYYY-MM') = m.month
	`

	var query string
	if status != "" {
		query = baseQuery + " AND wp.status = '" + status + "'"
	} else {
		query = baseQuery
	}

	query += `
		GROUP BY m.month
		ORDER BY m.month DESC;
	`

	err := tx.Raw(query).Scan(&results).Error

	if err != nil {
		return nil, exceptions.ParseGormError(tx, err)
	}

	monthlyPayouts := make([]dto.MonthlyPayoutRes, len(results))
	for i, result := range results {
		monthlyPayouts[i] = dto.MonthlyPayoutRes{
			Month:       result.Month,
			TotalPayout: result.TotalPayout,
		}
	}

	return monthlyPayouts, nil
}
