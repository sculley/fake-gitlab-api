package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authentication is a middleware that checks if the request has
// a valid PRIVATE-TOKEN header.
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("PRIVATE-TOKEN") == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "401 Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
