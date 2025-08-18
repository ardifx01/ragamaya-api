package helpers

import (
	"net/http"
	"ragamaya-api/api/users/dto"
	"ragamaya-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) (dto.UserOutput, *exceptions.Exception) {
	var result dto.UserOutput
	user_data, _ := c.Get("user")

	result, ok := user_data.(dto.UserOutput)
	if !ok {
		return result, exceptions.NewException(http.StatusUnauthorized, exceptions.ErrInvalidTokenStructure)
	}

	return result, nil
}
