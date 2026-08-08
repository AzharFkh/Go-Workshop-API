package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequiredRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role")

		for _, role := range roles {
			if userRole == role {
				c.Next()
				return 
			}
		}
		
		c.JSON(http.StatusForbidden, gin.H{
			"error": "forbidden",
		})

		c.Abort()
	}
}