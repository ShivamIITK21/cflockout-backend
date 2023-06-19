package middleware

import (
	"net/http"

	"github.com/ShivamIITK21/cflockout-backend/helpers"
	"github.com/gin-gonic/gin"
)


func Authorize() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error":"No token"})
			c.Abort()
			return
		}

		claims, msg := helpers.ValidateToken(clientToken)
		if msg != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error":msg})
			c.Abort()
			return 
		}

		c.Set("username", claims.Username)
		c.Set("cfid", claims.CFid)
		c.Set("user_type", claims.UserType)
		c.Next()
	}
}