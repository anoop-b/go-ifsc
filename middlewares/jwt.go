package middlewares

import (
	"go-ifsc/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Token")
		//tokenString := authHeader[len(BEARER_SCHEMA):]
		if authHeader != "" {
			tokenString := authHeader
			jwt := helpers.NewJWTService()
			token, err := jwt.ValidateToken(tokenString)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorised")
				return
			}
			if token.Valid {
				c.Next()
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorised")
		}
		c.Abort()
	}
}
