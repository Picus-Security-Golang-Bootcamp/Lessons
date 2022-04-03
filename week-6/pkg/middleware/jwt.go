package mw

import (
	"github.com/gin-gonic/gin"
	jwtHelper "github.com/h4yfans/patika-bookstore/pkg/jwt"

	"net/http"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				for _, role := range decodedClaims.Roles {
					if role == "admin" {
						c.Next()
						c.Abort()
						return
					}
				}
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()
		return
	}
}
