package middlewares

import "github.com/gin-gonic/gin"

// returns a map with the user names and passwords
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	})
}
