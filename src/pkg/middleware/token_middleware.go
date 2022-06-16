package middleware

import (
	"net/http"
	"os"

	jwt_helper "github.com/aliaydins/todo/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetHeader("access_token") != "" {
			decodedClaims := jwt_helper.VerifyToken(c.GetHeader("access_token"), secretKey, os.Getenv("ENV"))
			if decodedClaims != nil {
				if decodedClaims.Exp < jwt_helper.GetCurrentTime() {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
					c.Abort()
					return
				}

			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
			c.Abort()
		}
		c.Next()
	}
}
