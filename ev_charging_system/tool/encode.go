package tool

import "encoding/base64"

// EncodeToString 将[]byte编码为string
func EncodeToString(sign []byte) string {
	return base64.StdEncoding.EncodeToString(sign)
}

// DecodeToString 将string解码为[]byte
func DecodeToString(sign string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(sign)
}
