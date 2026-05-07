package middleware

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		userIDStr := c.GetHeader("X-User-ID")
         
		if userIDStr == "" || userIDStr == "undefined" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID is missing or undefined"})
            return
        }
	
		uid, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid User ID"})
			return
		}

		userID := int32(uid)

		c.Set("userID", &userID)
		c.Next()
	}
}