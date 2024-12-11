package middleware

import (
	"asset_management/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware 认证中间件，检查 JWT Token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 请求头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "request not hash authorized"})
			c.Abort()
			return
		}

		// 提取 Bearer token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // 检查是否有 Bearer 前缀
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token format error"})
			c.Abort()
			return
		}

		// 解析并验证 token
		claims, err := tool.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not valid token"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)
		c.Next()
	}
}
