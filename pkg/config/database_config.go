package config

import (
	"fmt"
	"log"
	"ragamaya-api/pkg/logger"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	GORMLogger "gorm.io/gorm/logger"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func InitDB() *gorm.DB {
	godotenv.Load()
	logger.Info("Initializing database connection...")
	start := time.Now()

	dbConfig := DBConfig{
		User:     GetDBUser(),
		Password: GetDBPassword(),
		Host:     GetDBHost(),
		Port:     GetDBPort(),
		Name:     GetDBName(),
	}

	dbURI := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s TimeZone=Asia/Jakarta", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	connection, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{
		Logger: GORMLogger.Default.LogMode(GORMLogger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	logger.Info("Connected to database in %s", elapsed)

	return connection
}
