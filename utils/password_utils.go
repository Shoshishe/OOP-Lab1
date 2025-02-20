package utils

import (
	"crypto/sha512"
	"fmt"
)

const saltCrypto = "35edtryuiojhgytfe3"

func GenerateHashedPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(saltCrypto)))
}
