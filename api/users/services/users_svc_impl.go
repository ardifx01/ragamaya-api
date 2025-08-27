package services

import (
	"ragamaya-api/api/users/dto"
	"ragamaya-api/api/users/repositories"
	"ragamaya-api/models"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
	"ragamaya-api/pkg/logger"

	walletRepo "ragamaya-api/api/wallets/repositories"

	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate

	walletRepo walletRepo.CompRepositories
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate, walletRepo walletRepo.CompRepositories) CompServices {
	return &CompServicesImpl{
		repo:       compRepositories,
		DB:         db,
		validate:   validate,
		walletRepo: walletRepo,
	}
}

func (s *CompServicesImpl) ClaimGoogleUserData(ctx *gin.Context, accessToken string) (*dto.GoogleUserData, *exceptions.Exception) {
	if accessToken == "" {
		return nil, exceptions.NewException(401, exceptions.ErrInvalidCredentials)
	}

	client := resty.New()
	var userInfo dto.GoogleUserData

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+accessToken).
		SetResult(&userInfo).
		Get("https://www.googleapis.com/oauth2/v3/userinfo")

	if err != nil {
		logger.Error("Request error: " + err.Error())
		return nil, exceptions.NewException(401, exceptions.ErrInvalidCredentials)
	}

	if resp.IsError() {
		logger.Error("HTTP error: " + resp.String())
		return nil, exceptions.NewException(401, exceptions.ErrInvalidCredentials)
	}

	logger.Info("Success: " + resp.String())
	return &userInfo, nil

}

func (s *CompServicesImpl) Login(ctx *gin.Context, data dto.LoginRequest) (*dto.TokenResponse, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	googleData, err := s.ClaimGoogleUserData(ctx, data.AccessToken)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.FindByEmail(ctx, s.DB, googleData.Email)
	if err != nil {
		if err.Status == 404 {
			tx := s.DB.Begin()
			defer helpers.CommitOrRollback(tx)

			user = &models.Users{
				UUID:            uuid.NewString(),
				Email:           googleData.Email,
				IsEmailVerified: googleData.EmailVerified,
				SUB:             googleData.Sub,
				Name:            googleData.Name,
				Role:            "user",
				AvatarURL:       googleData.Picture,
			}
			if err := s.repo.Create(ctx, tx, *user); err != nil {
				return nil, err
			}

			if err := s.walletRepo.Create(ctx, tx, models.Wallet{
				UserUUID: user.UUID,
			}); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.UUID
	claims["email"] = user.Email
	claims["is_email_verified"] = user.IsEmailVerified
	claims["sub"] = user.SUB
	claims["name"] = user.Name
	claims["role"] = user.Role
	claims["avatar_url"] = user.AvatarURL
	claims["seller_profile"] = user.SellerProfile.ToJWTPayload()
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()
	accessTokenStr, signErr := token.SignedString([]byte(jwtSecret))
	if signErr != nil {
		return nil, exceptions.NewException(500, "Failed to generate access token")
	}

	refreshTokenRaw := helpers.GenerateRandomString(64)
	refreshTokenExp := time.Now().Add(time.Hour * 24 * 7)
	refreshTokenModel := models.RefreshToken{
		UserUUID:  user.UUID,
		Token:     refreshTokenRaw,
		ExpiresAt: refreshTokenExp,
		CreatedAt: time.Now(),
	}
	tx := s.DB.Begin()
	if err := s.repo.CreateRefreshToken(ctx, tx, refreshTokenModel); err != nil {
		tx.Rollback()
		return nil, exceptions.NewException(500, "Failed to save refresh token")
	}
	tx.Commit()

	return &dto.TokenResponse{AccessToken: accessTokenStr, RefreshToken: refreshTokenRaw}, nil
}

func (s *CompServicesImpl) RefreshToken(ctx *gin.Context, refreshToken string) (accessToken string, err *exceptions.Exception) {
	tokenModel, errFind := s.repo.FindRefreshToken(ctx, s.DB, refreshToken)
	if errFind != nil {
		return "", errFind
	}
	if tokenModel == nil || tokenModel.ExpiresAt.Before(time.Now()) {
		return "", exceptions.NewException(401, "Refresh token expired or not found")
	}

	user, err := s.repo.FindByUUID(ctx, s.DB, tokenModel.UserUUID)
	if err != nil {
		return "", err
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.UUID
	claims["email"] = user.Email
	claims["is_email_verified"] = user.IsEmailVerified
	claims["sub"] = user.SUB
	claims["name"] = user.Name
	claims["role"] = user.Role
	claims["avatar_url"] = user.AvatarURL
	claims["seller_profile"] = user.SellerProfile.ToJWTPayload()
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()
	accessTokenStr, signErr := token.SignedString([]byte(jwtSecret))
	if signErr != nil {
		return "", exceptions.NewException(500, "Failed to generate access token")
	}

	return accessTokenStr, nil
}

func (s *CompServicesImpl) Logout(ctx *gin.Context, accessToken, refreshToken string) *exceptions.Exception {
	tx := s.DB.Begin()

	claims := jwt.MapClaims{}
	jwtSecret := os.Getenv("JWT_SECRET")
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err == nil {
		exp, _ := claims["exp"].(float64)
		helpers.SetBlacklistedToken(accessToken, time.Unix(int64(exp), 0))
	}

	s.repo.DeleteRefreshToken(ctx, tx, refreshToken)
	tx.Commit()
	return nil
}
