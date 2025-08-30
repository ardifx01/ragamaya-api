package helpers

import (
	"net/http"
	"ragamaya-api/api/users/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) (dto.UserRes, *exceptions.Exception) {
	var result dto.UserRes
	user_data, _ := c.Get("user")

	result, ok := user_data.(dto.UserRes)
	if !ok {
		return result, exceptions.NewException(http.StatusUnauthorized, exceptions.ErrInvalidTokenStructure)
	}

	return result, nil
}
