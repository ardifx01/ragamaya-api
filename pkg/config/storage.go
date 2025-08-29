package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func InitStorage() *s3.Client {
	godotenv.Load()

	awsAccessKey := GetAWSAccessKey()
	awsSecretKey := GetAWSSecretKey()
	awsRegion := GetAWSRegion()

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, "")),
	)
	if err != nil {
		log.Fatal("unable to load SDK config, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)

	return client
}
