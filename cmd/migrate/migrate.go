package main

import (
	"ragamaya-api/models"
	"ragamaya-api/pkg/config"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(
		&models.Users{},
		&models.Sellers{},
		&models.Clients{},
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
	)
	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}
