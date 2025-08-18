package services

import (
	"ragamaya-api/api/sellers/dto"
	"ragamaya-api/api/sellers/repositories"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
	"ragamaya-api/pkg/mapper"

	userRepositories "ragamaya-api/api/users/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate

	userRepo userRepositories.CompRepositories
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate, userRepo userRepositories.CompRepositories) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
		userRepo: userRepo,
	}
}

func (s *CompServicesImpl) Register(ctx *gin.Context, data dto.RegisterReq) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return err
	}

	input := mapper.MapSellerITM(data)
	input.UUID = uuid.NewString()
	input.UserUUID = userData.UUID

	if input.AvatarURL == "" {
		input.AvatarURL = userData.AvatarURL
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err = s.repo.Create(ctx, tx, input)
	if err != nil {
		return err
	}

	err = s.userRepo.Update(ctx, tx, models.Users{
		UUID: userData.UUID,
		Role: models.Seller,
	})

	return nil
}
