package main

import (
	"ragamaya-api/models"
	"ragamaya-api/pkg/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.InitConfig()
	db := config.InitDB()

	err := db.AutoMigrate(
		&models.Users{},
		&models.Sellers{},
		&models.RefreshToken{},
		&models.BlacklistedToken{},
		&models.VerificationToken{},
		&models.Files{},
		&models.Products{},
		&models.ProductThumbnails{},
		&models.ProductDigitalFiles{},
		&models.Orders{},
		&models.Payments{},
		&models.PaymentActions{},
		&models.PaymentVANumbers{},
		&models.ProductDigitalOwned{},
		&models.Wallet{},
		&models.WalletTransactionHistory{},
		&models.WalletPayoutRequest{},
		&models.WalletPayoutTransactionReceipt{},
		&models.ArticleCategory{},
		&models.Article{},
		&models.QuizCategory{},
		&models.Quiz{},
	)
	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}
