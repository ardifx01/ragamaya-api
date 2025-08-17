package main

import (
	"ragamaya-api/models"
	"ragamaya-api/pkg/config"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(
		&models.Users{},
		&models.Clients{},
		&models.RefreshToken{},
		&models.BlacklistedToken{},
		&models.VerificationToken{},
	)
	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}
