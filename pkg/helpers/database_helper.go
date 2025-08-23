package helpers

import (
	"strings"

	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil || tx.Error != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func IsDuplicateKeyError(err error) bool {
	errMsg := strings.ToLower(err.Error())
	return strings.Contains(errMsg, "duplicate") ||
		strings.Contains(errMsg, "unique") ||
		strings.Contains(errMsg, "constraint")
}