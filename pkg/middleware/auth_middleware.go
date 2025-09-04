package middleware

import (
	"net/http"
	"ragamaya-api/api/users/dto"
	"ragamaya-api/pkg/config"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := config.GetJWTSecret()
		var secretKey = []byte(secret)
		var tokenString string

		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			authHeaderParts := strings.Split(authHeader, " ")
			if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
				c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
				return
			} else {
				tokenString = authHeaderParts[1]
			}
		} else if c.Query("authorization") != "" {
			tokenString = c.Query("authorization")
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrForbidden))
			return
		}

		isBlacklisted, _ := helpers.IsTokenBlacklisted(tokenString)
		if isBlacklisted {
			c.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.NewException(http.StatusUnauthorized, "Token is blacklisted"))
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		user := dto.UserRes{
			UUID:            claims["uuid"].(string),
			Email:           claims["email"].(string),
			IsEmailVerified: claims["is_email_verified"].(bool),
			SUB:             claims["sub"].(string),
			Name:            claims["name"].(string),
			Role:            dto.Roles(claims["role"].(string)),
			AvatarURL:       claims["avatar_url"].(string),
		}

		c.Set("user", user)
		c.Next()
	}
}

func SellerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := config.GetJWTSecret()
		var secretKey = []byte(secret)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrForbidden))
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		tokenString := authHeaderParts[1]

		isBlacklisted, _ := helpers.IsTokenBlacklisted(tokenString)
		if isBlacklisted {
			c.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.NewException(http.StatusUnauthorized, "Token is blacklisted"))
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		if dto.Roles(claims["role"].(string)) != dto.Seller {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrNotSeller))
			return
		}

		sellerProfileInterface, ok := claims["seller_profile"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrNotSeller))
			return
		}

		sellerProfile, ok := sellerProfileInterface.(map[string]interface{})
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		user := dto.UserRes{
			UUID:            claims["uuid"].(string),
			Email:           claims["email"].(string),
			IsEmailVerified: claims["is_email_verified"].(bool),
			SUB:             claims["sub"].(string),
			Name:            claims["name"].(string),
			Role:            dto.Roles(claims["role"].(string)),
			AvatarURL:       claims["avatar_url"].(string),
			SellerProfile: dto.SellerRes{
				UUID:      sellerProfile["uuid"].(string),
				Name:      sellerProfile["name"].(string),
				AvatarURL: sellerProfile["avatar_url"].(string),
			},
		}
		c.Set("user", user)
		c.Next()
	}
}

func OptionalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := config.GetJWTSecret()
		var secretKey = []byte(secret)
		var tokenString string

		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			authHeaderParts := strings.Split(authHeader, " ")
			if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
				c.Next()
				return
			} else {
				tokenString = authHeaderParts[1]
			}
		} else if c.Query("authorization") != "" {
			tokenString = c.Query("authorization")
		} else {
			c.Next()
			return
		}

		isBlacklisted, _ := helpers.IsTokenBlacklisted(tokenString)
		if isBlacklisted {
			c.Next()
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.Next()
			return
		}

		if !token.Valid {
			c.Next()
			return
		}

		user := dto.UserRes{
			UUID:            claims["uuid"].(string),
			Email:           claims["email"].(string),
			IsEmailVerified: claims["is_email_verified"].(bool),
			SUB:             claims["sub"].(string),
			Name:            claims["name"].(string),
			Role:            dto.Roles(claims["role"].(string)),
			AvatarURL:       claims["avatar_url"].(string),
		}

		c.Set("user", user)
		c.Next()
	}
}