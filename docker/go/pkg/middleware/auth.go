package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/pkg/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Authorization token required"})
			c.Abort()
			return
		}

		authVal := strings.Split(authHeader, " ")
		if len(authVal) != 2 || authVal[0] != "Bearer" {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid authorization token"})
			c.Abort()
			return
		}

		claims, err := auth.VerifyToken(authVal[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
