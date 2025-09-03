package exceptions

import (
	"errors"
	"net/http"
	"ragamaya-api/pkg/logger"
	"strings"

	"gorm.io/gorm"
)

func ParseGormError(tx *gorm.DB, err error) *Exception {
	if tx != nil && tx.Statement != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
	}

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return &Exception{
			Message: "Record not found",
			Status:  http.StatusNotFound,
		}

	case errors.Is(err, gorm.ErrDuplicatedKey):
		return &Exception{
			Message: "Data already exists",
			Status:  http.StatusConflict,
		}

	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return &Exception{
			Message: "Related record not found",
			Status:  http.StatusBadRequest,
		}

	case errors.Is(err, gorm.ErrInvalidData):
		return &Exception{
			Message: "Invalid data",
			Status:  http.StatusBadRequest,
		}

	case strings.Contains(err.Error(), "duplicate key"):
		return &Exception{
			Message: "Data already exists",
			Status:  http.StatusConflict,
		}

	default:
		logger.Error("Database error: %v", err.Error())

		return &Exception{
			Message: "Database error occurred",
			Status:  http.StatusInternalServerError,
		}
	}
}
