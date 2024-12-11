package tool

import (
	"crypto/sha256"
	"encoding/hex"
)

// CalculateSHA256Hash 计算给定数据的 SHA-256 哈希值，并返回十六进制字符串表示
func CalculateSHA256Hash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
