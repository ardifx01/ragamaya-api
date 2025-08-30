package services

import (
	"bytes"
	"net/http"
	"ragamaya-api/api/storages/dto"
	"ragamaya-api/api/storages/repositories"
	"ragamaya-api/pkg/config"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/mapper"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	s3client *s3.Client
	validate *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, s3client *s3.Client, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		s3client: s3client,
		validate: validate,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.FilesInput) (*dto.FilesRes, *exceptions.Exception) {
	AWS_BUCKET := config.GetAWSBucket()
	AWS_FOLDER := config.GetStorageFolder()

	input := mapper.MapFilesInputToModel(data)
	input.UUID = uuid.NewString()
	input.PublicURL = "https://cdn.xann.my.id/" + AWS_FOLDER + "/" + input.UUID + "." + data.Extension

	fileKey := AWS_FOLDER + "/" + input.UUID + "." + data.Extension
	fileReader := bytes.NewReader(data.FileBuffer)

	_, exc := s.s3client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(AWS_BUCKET),
		Key:         aws.String(fileKey),
		Body:        fileReader,
		ContentType: aws.String(data.MimeType + "/" + data.MimeSubType),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if exc != nil {
		return nil, exceptions.NewException(http.StatusBadGateway, "Failed to upload file to S3")
	}

	result, err := s.repo.Create(ctx, s.DB, input)
	if err != nil {
		return nil, err
	}

	output := mapper.MapFilesMTO(*result)

	return &output, nil
}

func (s *CompServicesImpl) FindAllImages(ctx *gin.Context) ([]dto.FilesRes, *exceptions.Exception) {
	data, err := s.repo.FindAllImages(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	var output []dto.FilesRes

	for _, item := range *data {
		output = append(output, mapper.MapFilesMTO(item))
	}

	return output, nil
}
