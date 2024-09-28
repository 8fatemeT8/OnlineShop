package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateRequest(ctx *gin.Context) {
	fmt.Println("hi from validation middleware")
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func HasAdminAccess(c *gin.Context) {
	roleName, err := ExtractTokenCustomClaim(c, "role")
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}

	if roleName != "admin" {
		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}

	c.Next()
}
