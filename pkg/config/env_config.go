package config

import (
	"os"

	"github.com/joho/godotenv"
	"ragamaya-api/pkg/logger"
)

type Config struct {
	DBUser            string
	DBPassword        string
	DBHost            string
	DBPort            string
	DBName            string
	Port              string
	JWTSecret         string
	Environment       string
	AdminUsername     string
	AdminPassword     string
	RedisAddr         string
	RedisPass         string
	AWSAccessKey      string
	AWSSecretKey      string
	AWSRegion         string
	AWSBucket         string
	StorageFolder     string
	MidtransServerKey string
	MidtransEnv       string
}

var globalConfig *Config

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		logger.PanicError("No .env file found, using system environment variables")
	} else {
		logger.Info("Loaded .env file")
	}

	config := &Config{
		DBUser:            getEnv("DB_USER"),
		DBPassword:        getEnv("DB_PASSWORD"),
		DBHost:            getEnv("DB_HOST"),
		DBPort:            getEnv("DB_PORT"),
		DBName:            getEnv("DB_NAME"),
		Port:              getEnv("PORT"),
		JWTSecret:         getEnv("JWT_SECRET"),
		Environment:       getEnv("ENVIRONMENT"),
		AdminUsername:     getEnv("ADMIN_USERNAME"),
		AdminPassword:     getEnv("ADMIN_PASSWORD"),
		RedisAddr:         getEnv("REDIS_ADDR"),
		RedisPass:         getEnv("REDIS_PASS"),
		AWSAccessKey:      getEnv("AWS_ACCESS_KEY"),
		AWSSecretKey:      getEnv("AWS_SECRET_KEY"),
		AWSRegion:         getEnv("AWS_REGION"),
		AWSBucket:         getEnv("AWS_BUCKET"),
		StorageFolder:     getEnv("STORAGE_FOLDER"),
		MidtransServerKey: getEnv("MIDTRANS_SERVER_KEY"),
		MidtransEnv:       getEnv("MIDTRANS_ENV"),
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

func GetDBUser() string            { return GetConfig().DBUser }
func GetDBPassword() string        { return GetConfig().DBPassword }
func GetDBHost() string            { return GetConfig().DBHost }
func GetDBPort() string            { return GetConfig().DBPort }
func GetDBName() string            { return GetConfig().DBName }
func GetPort() string              { return GetConfig().Port }
func GetJWTSecret() string         { return GetConfig().JWTSecret }
func GetEnvironment() string       { return GetConfig().Environment }
func GetAdminUsername() string     { return GetConfig().AdminUsername }
func GetAdminPassword() string     { return GetConfig().AdminPassword }
func GetRedisAddr() string         { return GetConfig().RedisAddr }
func GetRedisPass() string         { return GetConfig().RedisPass }
func GetAWSAccessKey() string      { return GetConfig().AWSAccessKey }
func GetAWSSecretKey() string      { return GetConfig().AWSSecretKey }
func GetAWSRegion() string         { return GetConfig().AWSRegion }
func GetAWSBucket() string         { return GetConfig().AWSBucket }
func GetStorageFolder() string     { return GetConfig().StorageFolder }
func GetMidtransServerKey() string { return GetConfig().MidtransServerKey }
func GetMidtransEnv() string       { return GetConfig().MidtransEnv }

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