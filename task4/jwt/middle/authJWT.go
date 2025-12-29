package middle

import (
	"net/http"
	"strings"

	"github.com/Haley01114/init_project/task4/jwt/utils"
	"github.com/gin-gonic/gin"
)

// AuthJWT 自定义中间件: 验证 JWT token
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 提取 Authorization
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供有效的 Authorization"})
			c.Abort()
			return
		}
		// 校验 Authorization：Bearer token
		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization 格式错误"})
			c.Abort()
			return
		}
		// 获取 token
		token := parts[1]
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 Token"})
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
