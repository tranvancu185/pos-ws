package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hasBytes := hash.Sum(nil)

	return hex.EncodeToString(hasBytes)
}

func CompareHash(key, hash string) bool {
	password := GetHash(hash)
	return key == password
}
