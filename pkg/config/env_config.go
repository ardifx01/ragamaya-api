package config

import (
	"os"

	"ragamaya-api/pkg/logger"
)

type Config struct {
	DB_USER             string
	DB_PASSWORD         string
	DB_HOST             string
	DB_PORT             string
	DB_NAME             string
	PORT                string
	JWT_SECRET          string
	ENVIRONMENT         string
	ADMIN_USERNAME      string
	ADMIN_PASSWORD      string
	REDIS_ADDR          string
	REDIS_PASS          string
	AWS_ACCESS_KEY      string
	AWS_SECRET_KEY      string
	AWS_REGION          string
	AWS_BUCKET          string
	STORAGE_FOLDER      string
	MIDTRANS_SERVER_KEY string
	MIDTRANS_ENV        string
	FRONTEND_BASE_URL   string
	SMTP_EMAIL          string
	SMTP_PASSWORD       string
	SMTP_SERVER         string
	SMTP_PORT           string
}

var globalConfig *Config

func InitConfig() {
	config := &Config{
		DB_USER:             getEnv("DB_USER"),
		DB_PASSWORD:         getEnv("DB_PASSWORD"),
		DB_HOST:             getEnv("DB_HOST"),
		DB_PORT:             getEnv("DB_PORT"),
		DB_NAME:             getEnv("DB_NAME"),
		PORT:                getEnv("PORT"),
		JWT_SECRET:          getEnv("JWT_SECRET"),
		ENVIRONMENT:         getEnv("ENVIRONMENT"),
		ADMIN_USERNAME:      getEnv("ADMIN_USERNAME"),
		ADMIN_PASSWORD:      getEnv("ADMIN_PASSWORD"),
		REDIS_ADDR:          getEnv("REDIS_ADDR"),
		REDIS_PASS:          getEnv("REDIS_PASS"),
		AWS_ACCESS_KEY:      getEnv("AWS_ACCESS_KEY"),
		AWS_SECRET_KEY:      getEnv("AWS_SECRET_KEY"),
		AWS_REGION:          getEnv("AWS_REGION"),
		AWS_BUCKET:          getEnv("AWS_BUCKET"),
		STORAGE_FOLDER:      getEnv("STORAGE_FOLDER"),
		MIDTRANS_SERVER_KEY: getEnv("MIDTRANS_SERVER_KEY"),
		MIDTRANS_ENV:        getEnv("MIDTRANS_ENV"),
		FRONTEND_BASE_URL:   getEnv("FRONTEND_BASE_URL"),
		SMTP_EMAIL:          getEnv("SMTP_EMAIL"),
		SMTP_PASSWORD:       getEnv("SMTP_PASSWORD"),
		SMTP_SERVER:         getEnv("SMTP_SERVER"),
		SMTP_PORT:           getEnv("SMTP_PORT"),
	}

	globalConfig = config
	logger.Info("Configuration initialized successfully")
	return
}

func GetConfig() *Config {
	if globalConfig == nil {
		logger.PanicError("Configuration not initialized. Call InitConfig() first.")
	}
	return globalConfig
}

func GetDBUser() string            { return GetConfig().DB_USER }
func GetDBPassword() string        { return GetConfig().DB_PASSWORD }
func GetDBHost() string            { return GetConfig().DB_HOST }
func GetDBPort() string            { return GetConfig().DB_PORT }
func GetDBName() string            { return GetConfig().DB_NAME }
func GetPort() string              { return GetConfig().PORT }
func GetJWTSecret() string         { return GetConfig().JWT_SECRET }
func GetEnvironment() string       { return GetConfig().ENVIRONMENT }
func GetAdminUsername() string     { return GetConfig().ADMIN_USERNAME }
func GetAdminPassword() string     { return GetConfig().ADMIN_PASSWORD }
func GetRedisAddr() string         { return GetConfig().REDIS_ADDR }
func GetRedisPass() string         { return GetConfig().REDIS_PASS }
func GetAWSAccessKey() string      { return GetConfig().AWS_ACCESS_KEY }
func GetAWSSecretKey() string      { return GetConfig().AWS_SECRET_KEY }
func GetAWSRegion() string         { return GetConfig().AWS_REGION }
func GetAWSBucket() string         { return GetConfig().AWS_BUCKET }
func GetStorageFolder() string     { return GetConfig().STORAGE_FOLDER }
func GetMidtransServerKey() string { return GetConfig().MIDTRANS_SERVER_KEY }
func GetMidtransEnv() string       { return GetConfig().MIDTRANS_ENV }
func GetFrontendBaseURL() string   { return GetConfig().FRONTEND_BASE_URL }
func GetEmail() string             { return GetConfig().SMTP_EMAIL }
func GetEmailPassword() string     { return GetConfig().SMTP_PASSWORD }
func GetEmailServer() string       { return GetConfig().SMTP_SERVER }
func GetEmailPort() string         { return GetConfig().SMTP_PORT }

func IsProduction() bool         { return GetEnvironment() == "production" }
func IsDevelopment() bool        { return GetEnvironment() == "development" }
func IsMidtransProduction() bool { return GetMidtransEnv() == "production" }

func GetAWSConfig() (accessKey, secretKey, region, bucket string) {
	return GetAWSAccessKey(), GetAWSSecretKey(), GetAWSRegion(), GetAWSBucket()
}

func GetStoragePath() string {
	folder := GetStorageFolder()
	return folder
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		logger.PanicError("Environment variable %s is required but not set", key)
	}
	return value
}
