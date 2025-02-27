package utils

import (
	"crypto/sha256"
	"fmt"
)

const saltCrypto = "35edtryuiojhgytfe3"

func GenerateHashedPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(saltCrypto)))
}
