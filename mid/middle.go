package mid

import (
	"blog-f1sh/config"
	"blog-f1sh/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := parseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		var user models.User
		if err := config.DB.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func parseToken(tokenString string) (*TokenClaims, error) {
	// 实现你的 JWT 解析逻辑
	return nil, nil
}

type TokenClaims struct {
	UserID uint
}
}

func parseToken(tokenString string) (*TokenClaims, error) {
	// 这里实现你的 JWT 解析逻辑
	return nil, nil
}

type TokenClaims struct {
	UserID uint
	// 其他字段...
}
