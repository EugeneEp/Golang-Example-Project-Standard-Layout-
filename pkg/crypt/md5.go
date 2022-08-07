package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5Hash Функция получает MD5 хеш из переданной строки
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
