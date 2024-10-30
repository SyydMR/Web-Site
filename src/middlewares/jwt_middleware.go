package middlewares


import (
	"context"
	"net/http"
	"github.com/SyydMR/Web-Site/src/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token not provided"})
			c.Abort()
			return
		}

		userID, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), "id", userID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}