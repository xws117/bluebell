package controllder

import (
	"github.com/gin-gonic/gin"
)

func JWTAUTH() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Abort()
			return
		}

		c.Next()
	}
}
