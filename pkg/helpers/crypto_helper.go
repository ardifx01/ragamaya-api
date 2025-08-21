package helpers

import (
	"crypto/sha512"
	"encoding/hex"
)

func EncryptToSHA512(input string) string {
	h := sha512.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
