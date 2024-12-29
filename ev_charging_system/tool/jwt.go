package tool

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const jwtSecretKey = "VNERKNVDKJVNFDCcjsklcnmsdklcnsdklc"

type User struct {
	RepairmanId string `json:"repairmanId"`
	UserId      string `json:"userId"`
	UserType    int    `json:"userType"`
}

type JWTClaims struct {
	UserInfo User
	jwt.StandardClaims
}

// GenerateJWT 生成 JWT Token
func GenerateJWT(user User) (string, error) {
	claims := JWTClaims{
		UserInfo: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "go-health",
		},
	}

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名 token
	signedToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", err
	}

	return signedToken, nil
}

// ParseJWT 解析并验证 JWT Token
func ParseJWT(tokenString string) (*JWTClaims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil // 使用硬编码的密钥进行解析
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return nil, err
	}

	// 验证 token 是否有效
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
